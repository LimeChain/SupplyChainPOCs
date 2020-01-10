package poc01

import (
	"encoding/json"
	"fmt"
	"github.com/LimeChain/SupplyChainPOCs/constants"
	"github.com/LimeChain/SupplyChainPOCs/types"
	"github.com/LimeChain/SupplyChainPOCs/utils"
	"github.com/hyperledger/fabric/core/chaincode/shim"
	"github.com/hyperledger/fabric/protos/peer"
	"time"
)

type SupplyChainChaincode struct {
}

func (scc *SupplyChainChaincode) Init(stub shim.ChaincodeStubInterface) peer.Response {
	_, args := stub.GetFunctionAndParameters()

	if len(args) != 3 {
		return shim.Error(constants.ErrorArgumentsLength)
	}

	organizations, _ := json.Marshal(args)
	stub.PutState(constants.Organizations, organizations)
	
	return shim.Success(organizations)
}

func (scc *SupplyChainChaincode) Invoke(stub shim.ChaincodeStubInterface) peer.Response {
	funcName, args := stub.GetFunctionAndParameters()

	switch funcName {
		case constants.AddAssetType:
			return scc.addAssetType(stub, args)
		case constants.Manufacture:
			return scc.manufacture(stub, args)
		case constants.PlaceOrder:
			return scc.placeOrder(stub, args)
		case constants.FulfillOrder:
			return scc.fulfillOrder(stub, args)
		case constants.Assemble:
			return scc.assemble(stub, args)
		case constants.Sell:
			return scc.sell(stub, args)
		case constants.Query:
			return scc.query(stub, args)
	}

	return shim.Error(fmt.Sprintf(constants.ErrorInvalidFunctionName, funcName))
}

func (scc *SupplyChainChaincode) addAssetType(stub shim.ChaincodeStubInterface, args []string) peer.Response {
	if len(args) != 1 {
		return shim.Error(constants.ErrorArgumentsLength)
	}

	asset := types.Asset {}
	err := json.Unmarshal([]byte(args[0]), &asset)

	if err != nil {
		return shim.Error(err.Error())
	}

	assetId, err := utils.CreateCompositeKey(stub, constants.PrefixAsset)

	if err != nil {
		return shim.Error(err.Error())
	}

	asset.Id = assetId
	asset.DateCreated = time.Now()

	jsonAsset, _ := json.Marshal(asset)

	err = stub.PutState(asset.Id, jsonAsset)
	if err != nil {
		return shim.Error(err.Error())
	}

	return shim.Success(jsonAsset)
}

func (scc *SupplyChainChaincode) manufacture(stub shim.ChaincodeStubInterface, args []string) peer.Response {
	if len(args) != 1 {
		return shim.Error(constants.ErrorArgumentsLength)
	}

	record := types.Record {}
	err := json.Unmarshal([]byte(args[0]), &record)

	if err != nil {
		return shim.Error(err.Error())
	}

	assetBytes, _ := stub.GetState(record.AssetId)

	if len(assetBytes) == 0 {
		return shim.Error(fmt.Sprintf(constants.ErrorAssetIdNotFound, record.AssetId))
	}

	recordId, err := utils.CreateCompositeKey(stub, constants.PrefixRecord)

	if err != nil {
		return shim.Error(err.Error())
	}

	record.Id = recordId
	record.DateCreated = time.Now()
 
	jsonRecord, _ := json.Marshal(record)

	err = stub.PutState(record.Id, jsonRecord)
	if err != nil {
		return shim.Error(err.Error())
	}

	return shim.Success(jsonRecord)
}

func (scc *SupplyChainChaincode) placeOrder(stub shim.ChaincodeStubInterface, args []string) peer.Response {
	if len(args) != 1 {
		return shim.Error(constants.ErrorArgumentsLength)
	}

	order := types.Order {}
	err := json.Unmarshal([]byte(args[0]), &order)

	if err != nil {
		return shim.Error(err.Error())
	}

	assetBytes, _ := stub.GetState(order.AssetId)

	if len(assetBytes) == 0 {
		return shim.Error(fmt.Sprintf(constants.ErrorAssetIdNotFound, order.AssetId))
	}

	orderId, err := utils.CreateCompositeKey(stub, constants.PrefixOrder)

	if err != nil {
		return shim.Error(err.Error())
	}

	order.Id = orderId
	order.DateCreated = time.Now()
	order.IsCompleted = false

	jsonOrder, _ := json.Marshal(order)

	err = stub.PutState(order.Id, jsonOrder)

	if err != nil {
		return shim.Error(err.Error())
	}

	return shim.Success(jsonOrder)
}

func (scc *SupplyChainChaincode) fulfillOrder(stub shim.ChaincodeStubInterface, args []string) peer.Response {
	if len(args) != 1 {
		return shim.Error(constants.ErrorArgumentsLength)
	}

	orderFulfillment := types.OrderFulfillment {}
	err := json.Unmarshal([]byte(args[0]), &orderFulfillment)

	if err != nil {
		return shim.Error(err.Error())
	}

	orderBytes, _ := stub.GetState(orderFulfillment.Id)

	if len(orderBytes) == 0 {
		return shim.Error(fmt.Sprintf(constants.ErrorOrderIdNotFound, orderFulfillment.Id))
	}

	order := types.Order {}
	err = json.Unmarshal(orderBytes, &order)

	if err != nil {
		return shim.Error(err.Error())
	}

	if order.IsCompleted {
		return shim.Error(fmt.Sprintf(constants.ErrorOrderIsFulfilled,orderFulfillment.Id))
	}

	for _, recordElem := range orderFulfillment.Records {
		recordBytes, _ := stub.GetState(recordElem.Id)

		if len(recordBytes) == 0 {
			return shim.Error(fmt.Sprintf(constants.ErrorRecordIdNotFound, recordElem.Id))
		}
		record := types.Record {}
		err := json.Unmarshal(recordBytes, &record)

		if err != nil {
			return shim.Error(err.Error())
		}

		if record.AssetId != order.AssetId {
			return shim.Error(fmt.Sprintf(constants.ErrorRecordDifferentAssetId, record.Id, record.AssetId, order.AssetId))
		}

		if recordElem.Quantity > record.Quantity {
			return shim.Error(fmt.Sprintf(constants.ErrorRecordQuantity, recordElem.Id))
		}
		
		record.Quantity -= recordElem.Quantity
		record.LastUpdated = time.Now()

		newRecord := types.Record {
			AssetId: record.AssetId,
			BatchId: record.BatchId,
			Owner: order.BuyerId,
			Quantity: recordElem.Quantity,
			DateCreated: time.Now(),
			AssembledFrom: record.AssembledFrom,
			QualityCertificate: record.QualityCertificate }

		newRecordId, err := utils.CreateCompositeKey(stub, constants.PrefixRecord)
		
		if err != nil {
			return shim.Error(err.Error())
		}

		newRecord.Id = newRecordId
		
		jsonRecord, _ := json.Marshal(record)
		err = stub.PutState(record.Id, jsonRecord)

		if err != nil {
			return shim.Error(err.Error())
		}

		jsonRecord, _ = json.Marshal(newRecord)
		err = stub.PutState(newRecord.Id, jsonRecord)

		if err != nil {
			return shim.Error(err.Error())
		}
	}

	order.IsCompleted = true
	jsonOrder, _ := json.Marshal(order)
	err = stub.PutState(order.Id, jsonOrder)

	if err != nil {
		return shim.Error(err.Error())
	}

	return shim.Success(jsonOrder)
}

func (scc *SupplyChainChaincode) assemble(stub shim.ChaincodeStubInterface, args []string) peer.Response {
	if len(args) != 1 {
		return shim.Error(constants.ErrorArgumentsLength)
	}

	assembleRequest := types.AssembleRequest {}
	err := json.Unmarshal([]byte(args[0]), &assembleRequest)

	if err != nil {
		return shim.Error(err.Error())
	}

	assetBytes, _ := stub.GetState(assembleRequest.AssetId)
	
	if len(assetBytes) == 0 {
		return shim.Error(fmt.Sprintf(constants.ErrorAssetIdNotFound, assembleRequest.AssetId))
	}

	newRecord := types.Record {
		AssetId: assembleRequest.AssetId,
		BatchId: assembleRequest.BatchId,
		Owner: utils.GetOrganization(stub, constants.Org2Index),
		Quantity: assembleRequest.Quantity,
		DateCreated: time.Now(),
		AssembledFrom: []types.AssetAssembly{},
		QualityCertificate: assembleRequest.QualityCertificate }

	newRecordId, err := utils.CreateCompositeKey(stub, constants.PrefixRecord)
	
	if err != nil {
		return shim.Error(err.Error())
	}
	
	newRecord.Id = newRecordId

	for _, recordElem := range assembleRequest.Records {
		recordBytes, _ := stub.GetState(recordElem.Id)

		if len(recordBytes) == 0 {
			return shim.Error(fmt.Sprintf(constants.ErrorRecordIdNotFound, recordElem.Id))
		}
		
		record := types.Record {}
		err := json.Unmarshal(recordBytes, &record)

		if err != nil {
			return shim.Error(err.Error())
		}

		if recordElem.Quantity > record.Quantity {
			return shim.Error(fmt.Sprintf(constants.ErrorRecordQuantity, recordElem.Id))
		}
		
		record.Quantity -= recordElem.Quantity
		record.LastUpdated = time.Now()

		jsonRecord, _ := json.Marshal(record)
		err = stub.PutState(record.Id, jsonRecord)

		if err != nil {
			return shim.Error(err.Error())
		}
		
		newRecord.AssembledFrom = append(newRecord.AssembledFrom, recordElem)
	}

	jsonNewRecord, _ := json.Marshal(newRecord)
	err = stub.PutState(newRecord.Id, jsonNewRecord)

	if err != nil {
		return shim.Error(err.Error())
	}

	return shim.Success(jsonNewRecord)
}

func (scc *SupplyChainChaincode) sell(stub shim.ChaincodeStubInterface, args []string) peer.Response {
	if len(args) != 1 {
		return shim.Error(constants.ErrorArgumentsLength)
	}

	sellRequest := types.SellRequest {}
	err := json.Unmarshal([]byte(args[0]), &sellRequest)

	if err != nil {
		return shim.Error(err.Error())
	}

	recordBytes, _ := stub.GetState(sellRequest.RecordId)

	if len(recordBytes) == 0 {
		return shim.Error(fmt.Sprintf(constants.ErrorRecordIdNotFound, sellRequest.RecordId))
	}

	record := types.Record {}
	err = json.Unmarshal(recordBytes, &record)

	if err != nil {
		return shim.Error(err.Error())
	}

	if sellRequest.Quantity > record.Quantity {
		return shim.Error(fmt.Sprintf(constants.ErrorRecordQuantity, record.Id))
	}

	record.Quantity -= sellRequest.Quantity
	record.LastUpdated = time.Now()

	jsonRecord, _ := json.Marshal(record)

	err = stub.PutState(sellRequest.RecordId, jsonRecord)

	if err != nil {
		return shim.Error(err.Error())
	}

	return shim.Success(jsonRecord)
}

func (scc *SupplyChainChaincode) query(stub shim.ChaincodeStubInterface, args []string) peer.Response {
	queryResults, err := utils.GetQueryResultForQueryString(stub, args[0])

	if err != nil {
		return shim.Error(err.Error())
	}

	return shim.Success(queryResults)
}
