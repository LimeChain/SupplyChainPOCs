package dto

import (
	"github.com/shopspring/decimal"
)

type OrderDto struct {
	AssetId      string          `json:"assetId,omitempty"`
	SellerId     string          `json:"sellerId,omitempty"`
	BuyerId      string          `json:"buyerId,omitempty"`
	Quantity     uint64          `json:"quantity,omitempty"`
	PricePerUnit decimal.Decimal `json:"pricePerUnit,omitempty"`
}
