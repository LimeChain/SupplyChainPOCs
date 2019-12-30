package types 

import (
	"time"
	"github.com/shopspring/decimal"
)

type Order struct {
	Id           string          `json:"id,omitempty"`
	AssetId      string          `json:"assetId,omitempty"`
	SellerId     string          `json:"sellerId,omitempty"`
	BuyerId      string          `json:"buyerId,omitempty"`
	Quantity     uint64          `json:"quantity,omitempty"`
	PricePerUnit decimal.Decimal `json:"pricePerUnit,omitempty"`
	DateCreated  time.Time       `json:"dateCreated,omitempty"`
	DateUpdated  time.Time       `json:"dateFinished,omitempty"`
	IsCompleted  bool            `json:"isCompleted,omitempty"`
}
