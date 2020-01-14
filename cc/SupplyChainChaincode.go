package cc

import (
	"github.com/LimeChain/SupplyChainPOCs/types/dto"
	"github.com/LimeChain/SupplyChainPOCs/types/order"
	"github.com/LimeChain/SupplyChainPOCs/types/record"
	"github.com/LimeChain/SupplyChainPOCs/utils"
	"github.com/hyperledger/fabric/core/chaincode/shim"
	"github.com/hyperledger/fabric/protos/peer"
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

func (scc *SupplyChainChaincode) Query(stub shim.ChaincodeStubInterface, args [] string) peer.Response {
	queryResults, err := utils.GetQueryResultForQueryString(stub, args[0])

	if err != nil {
		return shim.Error(err.Error())
	}

	return shim.Success(queryResults)
}
