package order

import (
	"github.com/LimeChain/SupplyChainPOCs/types/dto"
	"github.com/shopspring/decimal"
	"time"
)

type Order struct {
	Id           string          `json:"id,omitempty"`
	SellerId     string          `json:"sellerId,omitempty"`
	BuyerId      string          `json:"buyerId,omitempty"`
	Quantity     uint64          `json:"quantity,omitempty"`
	PricePerUnit decimal.Decimal `json:"pricePerUnit,omitempty"`
	DateCreated  time.Time       `json:"dateCreated,omitempty"`
	DateUpdated  time.Time       `json:"dateFinished,omitempty"`
	IsCompleted  bool            `json:"isCompleted,omitempty"`
}

func NewOrder(id string, dto *dto.OrderDto) *Order {
	return &Order{
		Id:           id,
		SellerId:     dto.SellerId,
		BuyerId:      dto.BuyerId,
		Quantity:     dto.Quantity,
		PricePerUnit: dto.PricePerUnit,
		DateCreated:  time.Now(),
		IsCompleted:  false,
	}
}

func (order *Order) FulfillOrder(status bool) *Order {
	order.IsCompleted = status
	order.DateUpdated = time.Now()

	return order
}
