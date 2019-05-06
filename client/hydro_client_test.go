package client

import (
	"fmt"
	"github.com/hydroprotocol/amm-bots/utils"
	"github.com/shopspring/decimal"
	"testing"
)

func getTestClient(t *testing.T) *HydroClient {
	client := NewHydroClient(
		"0xa6553a3cbade744d6c6f63e557345402abd93e25cd1f1dba8bb0d374de2fcf4f",
		"HOT",
		"DAI",
		"http://localhost:3001",
	)
	return client
}

func TestHydroClient_CreateOrder(t *testing.T) {
	client := getTestClient(t)
	res, err := client.CreateOrder(
		decimal.New(1, -1),
		decimal.New(10, 0),
		utils.BUY,
		utils.LIMIT,
		0,
	)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%+v", res)
}

func TestHydroClient_CancelOrder(t *testing.T) {
	client := getTestClient(t)
	err := client.CancelOrder("0xf4a79c89f39b9d933e50b71df866f1fa2edc841d6b88d1400e54ea8505512f03")
	if err != nil {
		panic(err)
	}
}

func TestHydroClient_GetOrder(t *testing.T) {
	client := getTestClient(t)
	info, err := client.GetOrder("0x2778e141888aa17036ee6c691a158095d39941b8a39ec6f6a25dfa7d803c7c9c")
	if err != nil {
		panic(err)
	}
	fmt.Printf("%+v", info)
}

func TestHydroClient_GetAllPendingOrders(t *testing.T) {
	client := getTestClient(t)
	info, err := client.GetAllPendingOrders()
	if err != nil {
		panic(err)
	}
	fmt.Printf("%+v", info)
}

func TestHydroClient_CancelAllPendingOrders(t *testing.T) {
	client := getTestClient(t)
	success, err := client.CancelAllPendingOrders()
	if err != nil {
		panic(err)
	}
	fmt.Printf("%+v", success)
}
