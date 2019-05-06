package main

import (
	"github.com/hydroprotocol/amm-bots/algorithm"
	"github.com/hydroprotocol/amm-bots/client"
	"github.com/shopspring/decimal"
	"os"
)

func main() {

	botType := os.Getenv("BOT_TYPE")

	switch botType {
	case "CONST_PRODUCT":
		startConstProductBot()
	}

}

/*
Env checklist:
 - BOT_PRIVATE_KEY
 - BOT_QUOTE_TOKEN
 - BOT_BASE_URL
 - BOT_MIN_PRICE
 - BOT_MAX_PRICE
 - BOT_PRICE_GAP
 - BOT_EXPAND_INVENTORY
 - BOT_WEB3_URL
*/
func startConstProductBot() {
	privateKey := os.Getenv("BOT_PRIVATE_KEY")

	makerClient := client.NewHydroClient(
		privateKey,
		os.Getenv("BOT_BASE_TOKEN"),
		os.Getenv("BOT_QUOTE_TOKEN"),
		os.Getenv("BOT_BASE_URL"),
	)

	minPrice, _ := decimal.NewFromString(os.Getenv("BOT_MIN_PRICE"))
	maxPrice, _ := decimal.NewFromString(os.Getenv("BOT_MAX_PRICE"))
	priceGap, _ := decimal.NewFromString(os.Getenv("BOT_PRICE_GAP"))
	expandInventory, _ := decimal.NewFromString(os.Getenv("BOT_EXPAND_INVENTORY"))
	web3Url := os.Getenv("BOT_WEB3_URL")

	bot := algorithm.NewConstProductBot(
		makerClient,
		minPrice,
		maxPrice,
		priceGap,
		expandInventory,
		web3Url,
	)

	bot.Run()

}
