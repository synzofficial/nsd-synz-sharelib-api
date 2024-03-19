package typeconvertutil

import "github.com/shopspring/decimal"

func DecimalToFloat64(d decimal.Decimal) float64 {
	return d.InexactFloat64()
}
