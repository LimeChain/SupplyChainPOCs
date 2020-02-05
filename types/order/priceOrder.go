package order

import (
	"github.com/shopspring/decimal"
)

type PriceOrder struct {
	*BaseOrder
	PricePerUnit decimal.Decimal `json:"pricePerUnit,omitempty"`
}

func NewPriceOrder(order *BaseOrder, pricePerUnit decimal.Decimal) *PriceOrder {
	return &PriceOrder{
		BaseOrder:    order,
		PricePerUnit: pricePerUnit,
	}
}
