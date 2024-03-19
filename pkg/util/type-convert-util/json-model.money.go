package typeconvertutil

import (
	"github.com/shopspring/decimal"
	jsonmodel "github.com/synzofficial/nsd-synz-sharelib-api/pkg/model/json-model"
)

func NewJsonModelMoney(n1 int64, n2 int32) jsonmodel.Money {
	return jsonmodel.Money(decimal.New(n1, n2))
}

func JsonMoneyToFloat64(j jsonmodel.Money) float64 {
	return decimal.Decimal(j).InexactFloat64()
}
