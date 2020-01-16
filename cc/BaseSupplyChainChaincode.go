package cc

import (
	"github.com/LimeChain/SupplyChainPOCs/types/dto"
	"github.com/LimeChain/SupplyChainPOCs/types/order"
	"github.com/LimeChain/SupplyChainPOCs/types/record"
)

type BaseSupplyChainChaincode struct {
}

func (bsccc *BaseSupplyChainChaincode) Manufacture(id string, dto *dto.RecordDto) *record.Record {
	return record.NewRecord(id, dto)
}

func (bsccc *BaseSupplyChainChaincode) PlaceOrder(id string, orderDto *dto.OrderDto) *order.Order {
	return order.NewOrder(id, orderDto)
}

func (bsccc *BaseSupplyChainChaincode) FulfillOrder(orderStruct *order.Order, status bool) *order.Order {
	return orderStruct.FulfillOrder(status)
}
