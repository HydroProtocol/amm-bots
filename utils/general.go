package utils

import (
	"github.com/shopspring/decimal"
)

func SetDecimal(d decimal.Decimal, decimalNum int) decimal.Decimal {
	return d.Truncate(int32(decimalNum))
}

func SetPrecision(d decimal.Decimal, precision int) decimal.Decimal {
	if precision <= 0 {
		panic("precision must greater than 0")
	}
	numString := d.String()
	precisionCount := 0
	endPosition := 0
	for _, c := range numString {
		if c != '.' {
			precisionCount += 1
		}
		if precisionCount > precision {
			break
		}
		endPosition += 1
	}
	validDecimal, err := decimal.NewFromString(numString[:endPosition])
	if err != nil {
		panic("set precision failed")
	}
	return validDecimal
}

func ToggleSide(side string) string {
	if side == SELL {
		return BUY
	} else {
		return SELL
	}
}
