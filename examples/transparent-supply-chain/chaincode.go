package transparent_supply_chain

import (
	"encoding/json"
	"fmt"
	"github.com/LimeChain/SupplyChainPOCs/cc"
	"github.com/LimeChain/SupplyChainPOCs/constants"
	examplesConstants "github.com/LimeChain/SupplyChainPOCs/examples/constants"
	"github.com/LimeChain/SupplyChainPOCs/types/dto"
	"github.com/LimeChain/SupplyChainPOCs/types/order"
	"github.com/LimeChain/SupplyChainPOCs/types/record"
	"github.com/LimeChain/SupplyChainPOCs/utils"
	"github.com/hyperledger/fabric/core/chaincode/shim"
	"github.com/hyperledger/fabric/protos/peer"
	"time"
)

type TSCChaincode struct {
	cc.AssetBoundChaincode
	cc.ComposableChaincode
}

func (tsccc *TSCChaincode) Init(stub shim.ChaincodeStubInterface) peer.Response {
	_, args := stub.GetFunctionAndParameters()

	if len(args) != 3 {
		return shim.Error(constants.ErrorArgumentsLength)
	}

	organizations, err := json.Marshal(args)

	if err != nil {
		return shim.Error(err.Error())
	}

	err = stub.PutState(constants.Organizations, organizations)

	if err != nil {
		return shim.Error(err.Error())
	}

	return shim.Success(organizations)
}

func (tsccc *TSCChaincode) Invoke(stub shim.ChaincodeStubInterface) peer.Response {
	funcName, args := stub.GetFunctionAndParameters()

	switch funcName {
	case constants.AddAsset:
		return tsccc.addAsset(stub, args)
	case examplesConstants.Assemble:
		return tsccc.assemble(stub, args)
	case examplesConstants.Manufacture:
		return tsccc.manufacture(stub, args)
	case constants.PlaceOrder:
		return tsccc.placeOrder(stub, args)
	case constants.FulfillOrder:
		return tsccc.fulfillOrder(stub, args)
	case constants.Sell:
		return tsccc.sell(stub, args)
	case constants.Query:
		return tsccc.query(stub, args)
	}

	return shim.Error(fmt.Sprintf(constants.ErrorInvalidFunctionName, funcName))
}

func (tsccc *TSCChaincode) addAsset(stub shim.ChaincodeStubInterface, args []string) peer.Response {
	if len(args) != 1 {
		return shim.Error(constants.ErrorArgumentsLength)
	}

	assetDto := dto.AssetDto{}
	err := json.Unmarshal([]byte(args[0]), &assetDto)

	if err != nil {
		return shim.Error(err.Error())
	}

	assetId, err := utils.CreateCompositeKey(stub, constants.PrefixAsset)

	if err != nil {
		return shim.Error(err.Error())
	}

	assetStruct := tsccc.AddAsset(assetId, &assetDto)

	jsonAsset, err := json.Marshal(assetStruct)

	if err != nil {
		return shim.Error(err.Error())
	}

	err = stub.PutState(assetStruct.Id, jsonAsset)

	if err != nil {
		return shim.Error(err.Error())
	}

	return shim.Success(jsonAsset)
}

func (tsccc *TSCChaincode) manufacture(stub shim.ChaincodeStubInterface, args []string) peer.Response {
	if len(args) != 1 {
		return shim.Error(constants.ErrorArgumentsLength)
	}

	recordDto := dto.AssetBoundRecordDto{}
	err := json.Unmarshal([]byte(args[0]), &recordDto)

	if err != nil {
		return shim.Error(err.Error())
	}

	assetBytes, _ := stub.GetState(recordDto.AssetId)

	if len(assetBytes) == 0 {
		return shim.Error(fmt.Sprintf(constants.ErrorAssetIdNotFound, recordDto.AssetId))
	}

	recordId, err := utils.CreateCompositeKey(stub, constants.PrefixRecord)

	if err != nil {
		return shim.Error(err.Error())
	}

	recordStruct := tsccc.AssetBoundChaincode.Create(recordId, &recordDto)

	jsonRecord, err := json.Marshal(recordStruct)

	if err != nil {
		return shim.Error(err.Error())
	}

	err = stub.PutState(recordStruct.Id, jsonRecord)

	if err != nil {
		return shim.Error(err.Error())
	}

	return shim.Success(jsonRecord)
}

func (tsccc *TSCChaincode) placeOrder(stub shim.ChaincodeStubInterface, args []string) peer.Response {
	if len(args) != 1 {
		return shim.Error(constants.ErrorArgumentsLength)
	}

	assetBoundOrderDto := dto.AssetBoundOrderDto{}
	err := json.Unmarshal([]byte(args[0]), &assetBoundOrderDto)

	if err != nil {
		return shim.Error(err.Error())
	}

	assetBytes, _ := stub.GetState(assetBoundOrderDto.AssetId)

	if len(assetBytes) == 0 {
		return shim.Error(fmt.Sprintf(constants.ErrorAssetIdNotFound, assetBoundOrderDto.AssetId))
	}

	orderId, err := utils.CreateCompositeKey(stub, constants.PrefixOrder)

	if err != nil {
		return shim.Error(err.Error())
	}

	orderStruct := tsccc.AssetBoundChaincode.PlaceOrder(orderId, &assetBoundOrderDto)

	jsonOrder, err := json.Marshal(orderStruct)

	if err != nil {
		return shim.Error(err.Error())
	}

	err = stub.PutState(orderStruct.Id, jsonOrder)

	if err != nil {
		return shim.Error(err.Error())
	}

	return shim.Success(jsonOrder)
}

func (tsccc *TSCChaincode) fulfillOrder(stub shim.ChaincodeStubInterface, args []string) peer.Response {
	if len(args) != 1 {
		return shim.Error(constants.ErrorArgumentsLength)
	}

	orderFulfillmentDto := dto.OrderFulfillmentDto{}
	err := json.Unmarshal([]byte(args[0]), &orderFulfillmentDto)

	if err != nil {
		return shim.Error(err.Error())
	}

	orderBytes, _ := stub.GetState(orderFulfillmentDto.Id)

	if len(orderBytes) == 0 {
		return shim.Error(fmt.Sprintf(constants.ErrorOrderIdNotFound, orderFulfillmentDto.Id))
	}

	orderStruct := order.AssetBoundOrder{}
	err = json.Unmarshal(orderBytes, &orderStruct)

	if err != nil {
		return shim.Error(err.Error())
	}

	if orderStruct.IsCompleted {
		return shim.Error(fmt.Sprintf(constants.ErrorOrderIsFulfilled, orderFulfillmentDto.Id))
	}

	tsccc.AssetBoundChaincode.FulfillOrder(orderStruct.Order, orderFulfillmentDto.Status)

	if !orderStruct.IsCompleted {
		return shim.Error(fmt.Sprintf(constants.ErrorOrderIsNotFulfilled, orderStruct.Id))
	}

	for _, recordElem := range orderFulfillmentDto.Records {
		recordBytes, _ := stub.GetState(recordElem.Id)

		if len(recordBytes) == 0 {
			return shim.Error(fmt.Sprintf(constants.ErrorRecordIdNotFound, recordElem.Id))
		}
		recordStruct := record.AssetBoundRecord{}
		err := json.Unmarshal(recordBytes, &recordStruct)

		if err != nil {
			return shim.Error(err.Error())
		}

		if recordStruct.AssetId != orderStruct.AssetId {
			return shim.Error(fmt.Sprintf(constants.ErrorRecordDifferentAssetId, recordStruct.Id, recordStruct.AssetId, orderStruct.AssetId))
		}

		if recordElem.Quantity > recordStruct.Quantity {
			return shim.Error(fmt.Sprintf(constants.ErrorRecordQuantity, recordElem.Id))
		}

		recordStruct.DecreaseQuantity(recordElem.Quantity)

		newRecordStruct := record.AssetBoundRecord{
			BaseRecord: &record.BaseRecord{
				BatchId:         recordStruct.BatchId,
				CreationOrderId: orderStruct.Id,
				Owner:           orderStruct.BuyerId,
				Quantity:        recordElem.Quantity,
				DateCreated:     time.Now(),
			},
			AssetId: recordStruct.AssetId,
		}

		newRecordId, err := utils.CreateCompositeKey(stub, constants.PrefixRecord)

		if err != nil {
			return shim.Error(err.Error())
		}

		newRecordStruct.Id = newRecordId

		jsonRecord, err := json.Marshal(recordStruct)

		if err != nil {
			return shim.Error(err.Error())
		}

		err = stub.PutState(recordStruct.Id, jsonRecord)

		if err != nil {
			return shim.Error(err.Error())
		}

		jsonRecord, err = json.Marshal(newRecordStruct)

		if err != nil {
			return shim.Error(err.Error())
		}

		err = stub.PutState(newRecordStruct.Id, jsonRecord)

		if err != nil {
			return shim.Error(err.Error())
		}
	}
	jsonOrder, err := json.Marshal(orderStruct)

	if err != nil {
		return shim.Error(err.Error())
	}

	err = stub.PutState(orderStruct.Id, jsonOrder)

	if err != nil {
		return shim.Error(err.Error())
	}

	return shim.Success(jsonOrder)
}

func (tsccc *TSCChaincode) assemble(stub shim.ChaincodeStubInterface, args []string) peer.Response {
	if len(args) != 1 {
		return shim.Error(constants.ErrorArgumentsLength)
	}

	assetComposeRequest := dto.AssetComposeRequestDto{}
	err := json.Unmarshal([]byte(args[0]), &assetComposeRequest)

	if err != nil {
		return shim.Error(err.Error())
	}

	assetBytes, _ := stub.GetState(assetComposeRequest.AssetId)

	if len(assetBytes) == 0 {
		return shim.Error(fmt.Sprintf(constants.ErrorAssetIdNotFound, assetComposeRequest.AssetId))
	}

	newRecordId, err := utils.CreateCompositeKey(stub, constants.PrefixRecord)
	newRecord, updatedRecords, err := tsccc.ComposableChaincode.Compose(stub, newRecordId, assetComposeRequest.ComposeRequestDto)

	if err != nil {
		return shim.Error(err.Error())
	}

	for _, updatedRecord := range updatedRecords {
		recordBytes, _ := stub.GetState(updatedRecord.Id)

		recordStruct := record.ComposableRecord{}

		err := json.Unmarshal(recordBytes, &recordStruct)

		if err != nil {
			return shim.Error(err.Error())
		}

		recordStruct.SetQuantity(updatedRecord.Quantity)

		jsonRecord, err := json.Marshal(recordStruct)

		if err != nil {
			return shim.Error(err.Error())
		}

		err = stub.PutState(recordStruct.Id, jsonRecord)

		if err != nil {
			return shim.Error(err.Error())
		}
	}

	jsonNewRecord, err := json.Marshal(newRecord)

	if err != nil {
		return shim.Error(err.Error())
	}

	err = stub.PutState(newRecord.Id, jsonNewRecord)

	return shim.Success(jsonNewRecord)
}

func (tsccc *TSCChaincode) sell(stub shim.ChaincodeStubInterface, args []string) peer.Response {
	if len(args) != 1 {
		return shim.Error(constants.ErrorArgumentsLength)
	}

	sellRequest := dto.SellDto{}
	err := json.Unmarshal([]byte(args[0]), &sellRequest)

	if err != nil {
		return shim.Error(err.Error())
	}

	recordBytes, _ := stub.GetState(sellRequest.RecordId)

	if len(recordBytes) == 0 {
		return shim.Error(fmt.Sprintf(constants.ErrorRecordIdNotFound, sellRequest.RecordId))
	}

	recordStruct := record.BaseRecord{}
	err = json.Unmarshal(recordBytes, &recordStruct)

	if err != nil {
		return shim.Error(err.Error())
	}

	if sellRequest.Quantity > recordStruct.Quantity {
		return shim.Error(fmt.Sprintf(constants.ErrorRecordQuantity, recordStruct.Id))
	}

	recordStruct.DecreaseQuantity(sellRequest.Quantity)

	jsonRecord, err := json.Marshal(recordStruct)

	if err != nil {
		return shim.Error(err.Error())
	}

	err = stub.PutState(sellRequest.RecordId, jsonRecord)

	if err != nil {
		return shim.Error(err.Error())
	}

	return shim.Success(jsonRecord)
}

func (tsccc *TSCChaincode) query(stub shim.ChaincodeStubInterface, args []string) peer.Response {
	queryResults, err := utils.GetQueryResultForQueryString(stub, args[0])

	if err != nil {
		return shim.Error(err.Error())
	}

	return shim.Success(queryResults)
}
