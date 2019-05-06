package algorithm

import (
	"github.com/shopspring/decimal"
	"math"
)

type ConstProductLadder struct {
	UpPrice   decimal.Decimal
	DownPrice decimal.Decimal
	Amount    decimal.Decimal
}

func GenerateConstProductLadders(
	baseTokenAmount decimal.Decimal,
	quoteTokenAmount decimal.Decimal,
	minPrice decimal.Decimal,
	maxPrice decimal.Decimal,
	priceGap decimal.Decimal,
	expandInventory decimal.Decimal,
) ([]ConstProductLadder, error) {
	ladders := []ConstProductLadder{}

	// product
	product := baseTokenAmount.Mul(quoteTokenAmount)

	// P-center
	centerPrice := quoteTokenAmount.Div(baseTokenAmount)

	var upPrice decimal.Decimal
	var downPrice decimal.Decimal
	var lastBaseAmount decimal.Decimal

	// ask Orders
	downPrice = centerPrice
	lastBaseAmount = baseTokenAmount
	for true {
		upPrice = downPrice.Mul(decimal.New(1, 0).Add(priceGap))
		if upPrice.GreaterThan(maxPrice) {
			break
		}
		f, _ := product.Div(upPrice).Float64()
		newBaseAmount := decimal.NewFromFloat(math.Sqrt(f))
		ladders = append(ladders, ConstProductLadder{
			upPrice,
			downPrice,
			newBaseAmount.Sub(lastBaseAmount).Abs().Mul(expandInventory),
		})
		downPrice = upPrice
		lastBaseAmount = newBaseAmount
	}

	// bid Orders
	upPrice = centerPrice
	lastBaseAmount = baseTokenAmount
	for true {
		downPrice = upPrice.Div(decimal.New(1, 0).Add(priceGap))
		if downPrice.LessThan(minPrice) {
			break
		}
		f, _ := product.Div(downPrice).Float64()
		newBaseAmount := decimal.NewFromFloat(math.Sqrt(f))
		ladders = append(ladders, ConstProductLadder{
			upPrice,
			downPrice,
			newBaseAmount.Sub(lastBaseAmount).Abs().Mul(expandInventory),
		})
		upPrice = downPrice
		lastBaseAmount = newBaseAmount
	}

	return ladders, nil
}
