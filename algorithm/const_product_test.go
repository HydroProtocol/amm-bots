package algorithm

import (
	"fmt"
	"github.com/shopspring/decimal"
	"testing"
)

func TestGenerateConstProductLadders(t *testing.T) {
	base := decimal.New(100, 0)
	quote := decimal.New(10000, 0)
	centerPrice := quote.Div(base)
	ladders, _ := GenerateConstProductLadders(
		base,
		quote,
		decimal.New(90, 0),
		decimal.New(110, 0),
		decimal.New(1, -2),
		decimal.New(1, 0),
	)

	for _, ladder := range ladders {
		if ladder.DownPrice.GreaterThanOrEqual(centerPrice) {
			fmt.Printf("sell %s %s\n", ladder.UpPrice.String(), ladder.Amount.String())
		} else {
			fmt.Printf("buy %s %s\n", ladder.DownPrice.String(), ladder.Amount.String())
		}
	}
}
