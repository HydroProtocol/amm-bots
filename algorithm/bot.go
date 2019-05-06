package algorithm

import (
	"github.com/hydroprotocol/amm-bots/client"
	"github.com/hydroprotocol/amm-bots/utils"
	"github.com/shopspring/decimal"
	"github.com/sirupsen/logrus"
	"sync"
	"time"
)

func NewConstProductBot(
	client *client.HydroClient,
	minPrice decimal.Decimal,
	maxPrice decimal.Decimal,
	priceGap decimal.Decimal,
	expandInventory decimal.Decimal,
	web3Url string) *ConstProductBot {
	baseToken, quoteToken, err := client.GetTradingErc20()
	if err != nil {
		panic(err)
	}
	var lock sync.Mutex
	bot := ConstProductBot{
		client,
		baseToken,
		quoteToken,
		map[string]ConstProductLadder{},
		minPrice,
		maxPrice,
		priceGap,
		expandInventory,
		web3Url,
		&lock,
	}
	return &bot
}

type ConstProductBot struct {
	client          *client.HydroClient
	baseToken       *utils.ERC20
	quoteToken      *utils.ERC20
	ladderMap       map[string]ConstProductLadder
	minPrice        decimal.Decimal
	maxPrice        decimal.Decimal
	priceGap        decimal.Decimal
	expandInventory decimal.Decimal
	web3Url         string
	updateLock      *sync.Mutex
}

func (b *ConstProductBot) Run() {
	b.init()
	for true {
		for key := range b.ladderMap {
			b.updateLock.Lock()
			b.maintainOrder(key)
			b.updateLock.Unlock()
		}
		time.Sleep(15 * time.Second)
	}
}

func (b *ConstProductBot) init() {
	_, _ = b.client.CancelAllPendingOrders()
	baseTokenAmount, _, err := b.baseToken.GetBalance(b.web3Url, b.client.Address)
	if err != nil {
		panic(err)
	}
	quoteTokenAmount, _, err := b.quoteToken.GetBalance(b.web3Url, b.client.Address)
	if err != nil {
		panic(err)
	}
	ladders, err := GenerateConstProductLadders(
		*baseTokenAmount,
		*quoteTokenAmount,
		b.minPrice,
		b.maxPrice,
		b.priceGap,
		b.expandInventory,
	)
	centerPrice := quoteTokenAmount.Div(*baseTokenAmount)
	for _, ladder := range ladders {
		if ladder.UpPrice.LessThanOrEqual(centerPrice) {
			b.createOrder(ladder, utils.BUY)
		} else {
			b.createOrder(ladder, utils.SELL)
		}
	}
}

func (b *ConstProductBot) createOrder(ladder ConstProductLadder, side string) {
	var price decimal.Decimal
	if side == utils.SELL {
		price = ladder.UpPrice
	} else {
		price = ladder.DownPrice
	}
	orderId, err := b.client.CreateOrder(
		price,
		ladder.Amount,
		side,
		utils.LIMIT,
		0,
	)
	if err != nil {
		logrus.Warn("create order failed ", err)
	} else {
		b.ladderMap[orderId] = ladder
	}
}

func (b *ConstProductBot) maintainOrder(orderId string) {
	orderInfo, err := b.client.GetOrder(orderId)
	if err != nil {
		logrus.Warn("get order info failed ", err)
	} else {
		if orderInfo.Status == utils.ORDER_CLOSE && orderInfo.FilledAmount.GreaterThan(decimal.Zero) {
			b.createOrder(b.ladderMap[orderId], utils.ToggleSide(orderInfo.Side))
			delete(b.ladderMap, orderId)
		}
	}
}

func (b *ConstProductBot) ElegantExit() {
	b.updateLock.Lock()
	_, _ = b.client.CancelAllPendingOrders()
}
