package client

import (
	"github.com/hydroprotocol/amm-bots/utils"
	"github.com/shopspring/decimal"
)

type ITicker struct {
	Status int    `json:"status"`
	Desc   string `json:"desc"`
	Data   struct {
		Ticker struct {
			MarketID  string `json:"marketId"`
			Price     string `json:"price"`
			Volume    string `json:"volume"`
			Bid       string `json:"bid"`
			Ask       string `json:"ask"`
			Low       string `json:"low"`
			High      string `json:"high"`
			UpdatedAt int64  `json:"updatedAt"`
		} `json:"ticker"`
	} `json:"data"`
}

type IHydroMarkets struct {
	Status int    `json:"status"`
	Desc   string `json:"desc"`
	Data   struct {
		Markets []struct {
			ID                     string   `json:"id"`
			BaseToken              string   `json:"baseToken"`
			BaseTokenProjectURL    string   `json:"baseTokenProjectUrl"`
			BaseTokenName          string   `json:"baseTokenName"`
			BaseTokenDecimals      int      `json:"baseTokenDecimals"`
			BaseTokenAddress       string   `json:"baseTokenAddress"`
			QuoteToken             string   `json:"quoteToken"`
			QuoteTokenDecimals     int      `json:"quoteTokenDecimals"`
			QuoteTokenAddress      string   `json:"quoteTokenAddress"`
			MinOrderSize           string   `json:"minOrderSize"`
			PricePrecision         int      `json:"pricePrecision"`
			PriceDecimals          int      `json:"priceDecimals"`
			AmountDecimals         int      `json:"amountDecimals"`
			AsMakerFeeRate         string   `json:"asMakerFeeRate"`
			AsTakerFeeRate         string   `json:"asTakerFeeRate"`
			GasFeeAmount           string   `json:"gasFeeAmount"`
			SupportedOrderTypes    []string `json:"supportedOrderTypes"`
			MarketOrderMaxSlippage string   `json:"marketOrderMaxSlippage"`
			LastPriceIncrease      string   `json:"lastPriceIncrease"`
			LastPrice              string   `json:"lastPrice"`
			Price24H               string   `json:"price24h"`
			Amount24H              string   `json:"amount24h"`
			QuoteTokenVolume24H    string   `json:"quoteTokenVolume24h"`
		} `json:"markets"`
	} `json:"data"`
}

type IHydroOrderResp struct {
	ID              string `json:"id"`
	Type            string `json:"type"`
	Version         string `json:"version"`
	Status          string `json:"status"`
	Amount          string `json:"amount"`
	AvailableAmount string `json:"availableAmount"`
	PendingAmount   string `json:"pendingAmount"`
	CanceledAmount  string `json:"canceledAmount"`
	ConfirmedAmount string `json:"confirmedAmount"`
	Price           string `json:"price"`
	AveragePrice    string `json:"averagePrice"`
	Side            string `json:"side"`
	MakerFeeRate    string `json:"makerFeeRate"`
	TakerFeeRate    string `json:"takerFeeRate"`
	MakerRebateRate string `json:"makerRebateRate"`
	GasFeeAmount    string `json:"gasFeeAmount"`
	Account         string `json:"account"`
	CreatedAt       int64  `json:"createdAt"`
	MarketID        string `json:"marketId"`
}

type ITokenBalance struct {
	Status int    `json:"status"`
	Desc   string `json:"desc"`
	Data   struct {
		Balance       string `json:"balance"`
		LockedBalance string `json:"lockedBalance"`
	} `json:"data"`
}

type IOrder struct {
	Status int    `json:"status"`
	Desc   string `json:"desc"`
	Data   struct {
		Order IHydroOrderResp `json:"order"`
	} `json:"data"`
}

type IAllPendingOrders struct {
	Status int    `json:"status"`
	Desc   string `json:"desc"`
	Data   struct {
		Count  int               `json:"count"`
		Orders []IHydroOrderResp `json:"orders"`
	} `json:"data"`
}

type IBuildOrder struct {
	Status int    `json:"status"`
	Desc   string `json:"desc"`
	Data   struct {
		Order struct {
			ID string `json:"id"`
		} `json:"order"`
	} `json:"data"`
}

type IPlaceOrder struct {
	Status int    `json:"status"`
	Desc   string `json:"desc"`
	Data   struct {
		Order struct {
			ID string `json:"id"`
		} `json:"order"`
	} `json:"data"`
}

type ICancelOrder struct {
	Status int    `json:"status"`
	Desc   string `json:"desc"`
}

type BasicOrder struct {
	Amount decimal.Decimal
	Price  decimal.Decimal
	Side   string
}

type StdOrder struct {
	Id              string
	Status          string
	Amount          decimal.Decimal
	Price           decimal.Decimal
	AvailableAmount decimal.Decimal
	FilledAmount    decimal.Decimal
	Side            string
}

var EmptyStdOrder = StdOrder{
	"",
	utils.ORDER_CLOSE,
	decimal.New(0, 0),
	decimal.New(0, 0),
	decimal.New(0, 0),
	decimal.New(0, 0),
	utils.SELL,
}
