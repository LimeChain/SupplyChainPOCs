package cc

import (
	"encoding/json"
	"fmt"
	"github.com/LimeChain/SupplyChainPOCs/constants"
	"github.com/LimeChain/SupplyChainPOCs/types/dto"
	"github.com/LimeChain/SupplyChainPOCs/types/record"
	"github.com/LimeChain/SupplyChainPOCs/utils"
	"time"

	"github.com/hyperledger/fabric/core/chaincode/shim"
	"github.com/hyperledger/fabric/protos/peer"
)

type AssembableChaincode struct {
	SupplyChainChaincode
}

func (acc *AssembableChaincode) Manufacture(id string, src dto.AssembableRecordDto) *record.AssembableRecord {
	rec := acc.SupplyChainChaincode.Manufacture(id, src.RecordDto)

	records := record.AssembleRecord(src.AssembledFrom)
	return record.NewAssembableRecord(rec, records)
}

func (acc *AssembableChaincode) Assemble(stub shim.ChaincodeStubInterface, assembleRequest *dto.AssembleRequestDto) peer.Response {
	newRecord := record.AssembableRecord{
		Record: &record.Record{
			BatchId:            assembleRequest.BatchId,
			Owner:              utils.GetOrganization(stub, constants.Org2Index),
			Quantity:           assembleRequest.Quantity,
			DateCreated:        time.Now(),
			QualityCertificate: assembleRequest.QualityCertificate,
		},
		AssembledFrom: record.AssembleRecord{}}

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

		recordStruct := record.Record{}
		err := json.Unmarshal(recordBytes, &recordStruct)

		if err != nil {
			return shim.Error(err.Error())
		}

		if recordElem.Quantity > recordStruct.Quantity {
			return shim.Error(fmt.Sprintf(constants.ErrorRecordQuantity, recordElem.Id))
		}

		recordStruct.DecreaseQuantity(recordElem.Quantity)

		jsonRecord, err := json.Marshal(recordStruct)

		if err != nil {
			return shim.Error(err.Error())
		}

		err = stub.PutState(recordStruct.Id, jsonRecord)

		if err != nil {
			return shim.Error(err.Error())
		}

		newRecord.AssembledFrom = append(newRecord.AssembledFrom, recordElem)
	}

	jsonNewRecord, err := json.Marshal(newRecord)

	if err != nil {
		return shim.Error(err.Error())
	}

	err = stub.PutState(newRecord.Id, jsonNewRecord)

	if err != nil {
		return shim.Error(err.Error())
	}

	return shim.Success(jsonNewRecord)
}
