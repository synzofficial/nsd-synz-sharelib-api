package jsonmodel

import (
	"fmt"
	"strings"

	"github.com/shopspring/decimal"
)

type Money decimal.Decimal

func (m Money) MarshalJSON() ([]byte, error) {
	decimalString := decimal.Decimal(m).StringFixed(2)
	return []byte(decimalString), nil
}

func (m *Money) UnmarshalJSON(b []byte) error {
	decimalString := strings.ReplaceAll(string(b), "\"", "")
	d, err := decimal.NewFromString(decimalString)
	if err != nil {
		return fmt.Errorf("unable to convert %s to decimal", d)
	}
	*m = Money(d)
	return nil
}
