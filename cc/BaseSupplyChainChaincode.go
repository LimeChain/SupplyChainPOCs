package cc

import (
	"github.com/LimeChain/SupplyChainPOCs/types/dto"
	"github.com/LimeChain/SupplyChainPOCs/types/order"
	"github.com/LimeChain/SupplyChainPOCs/types/record"
)

type BaseSupplyChainChaincode struct {
}

func (bsccc *BaseSupplyChainChaincode) Create(id string, dto *dto.BaseRecordDto) *record.BaseRecord {
	return record.NewRecord(id, dto)
}

func (bsccc *BaseSupplyChainChaincode) PlaceOrder(id string, orderDto *dto.OrderDto) *order.BaseOrder {
	return order.NewOrder(id, orderDto)
}

func (bsccc *BaseSupplyChainChaincode) FulfillOrder(orderStruct *order.BaseOrder, status bool) *order.BaseOrder {
	return orderStruct.FulfillOrder(status)
}
