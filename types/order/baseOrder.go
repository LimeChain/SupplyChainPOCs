package order

import (
	"github.com/LimeChain/SupplyChainPOCs/types/dto"
	"time"
)

type BaseOrder struct {
	Id          string    `json:"id,omitempty"`
	SellerId    string    `json:"sellerId,omitempty"`
	BuyerId     string    `json:"buyerId,omitempty"`
	Quantity    uint64    `json:"quantity,omitempty"`
	DateCreated time.Time `json:"dateCreated,omitempty"`
	DateUpdated time.Time `json:"dateFinished,omitempty"`
	IsCompleted bool      `json:"isCompleted,omitempty"`
}

func NewOrder(id string, dto *dto.OrderDto) *BaseOrder {
	return &BaseOrder{
		Id:          id,
		SellerId:    dto.SellerId,
		BuyerId:     dto.BuyerId,
		Quantity:    dto.Quantity,
		DateCreated: time.Now(),
		IsCompleted: false,
	}
}

func (order *BaseOrder) FulfillOrder(status bool) *BaseOrder {
	order.IsCompleted = status
	order.DateUpdated = time.Now()

	return order
}
