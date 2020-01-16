package cc

import (
	"github.com/LimeChain/SupplyChainPOCs/types/dto"
	"github.com/LimeChain/SupplyChainPOCs/types/order"
	"github.com/LimeChain/SupplyChainPOCs/types/record"
)

type SupplyChainChaincode struct {
}

func (scc *SupplyChainChaincode) Manufacture(id string, dto *dto.RecordDto) *record.Record {
	return record.NewRecord(id, dto)
}

func (scc *SupplyChainChaincode) PlaceOrder(id string, orderDto *dto.OrderDto) *order.Order {
	return order.NewOrder(id, orderDto)
}

func (scc *SupplyChainChaincode) FulfillOrder(orderStruct *order.Order, status bool) *order.Order {
	return orderStruct.FulfillOrder(status)
}
