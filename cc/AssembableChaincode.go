package cc

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/LimeChain/SupplyChainPOCs/constants"
	"github.com/LimeChain/SupplyChainPOCs/types/dto"
	"github.com/LimeChain/SupplyChainPOCs/types/record"
	"github.com/LimeChain/SupplyChainPOCs/utils"
	"github.com/hyperledger/fabric/core/chaincode/shim"
	"time"
)

type AssembableChaincode struct {
	SupplyChainChaincode
}

func (acc *AssembableChaincode) Manufacture(id string, src dto.AssembableRecordDto) *record.AssembableRecord {
	rec := acc.SupplyChainChaincode.Manufacture(id, src.RecordDto)

	records := record.RecordParts(src.AssembledFrom)
	return record.NewAssembableRecord(rec, records)
}

func (acc *AssembableChaincode) Assemble(stub shim.ChaincodeStubInterface, id string, assembleRequest *dto.AssembleRequestDto, ) (*record.AssembableRecord, record.RecordParts, error) {
	newRecord := record.AssembableRecord{
		Record: &record.Record{
			Id:                 id,
			BatchId:            assembleRequest.BatchId,
			Owner:              utils.GetOrganization(stub, constants.Org2Index),
			Quantity:           assembleRequest.Quantity,
			DateCreated:        time.Now(),
			QualityCertificate: assembleRequest.QualityCertificate,
		},
		AssembledFrom: record.RecordParts{}}

	updatedRecords := record.RecordParts{}

	for _, recordElem := range assembleRequest.Records {
		recordBytes, _ := stub.GetState(recordElem.Id)

		if len(recordBytes) == 0 {
			return nil, nil, errors.New(fmt.Sprintf(constants.ErrorRecordIdNotFound, recordElem.Id))
		}

		recordStruct := record.Record{}
		err := json.Unmarshal(recordBytes, &recordStruct)

		if err != nil {
			return nil, nil, err
		}

		if recordElem.Quantity > recordStruct.Quantity {
			return nil, nil, errors.New(fmt.Sprintf(constants.ErrorRecordQuantity, recordElem.Id))
		}

		updatedRecords = append(updatedRecords, record.RecordParts{
			{Id: recordStruct.Id, Quantity: recordStruct.GetNewQuantity(recordElem.Quantity)}}...)

		newRecord.AssembledFrom = append(newRecord.AssembledFrom, recordElem)
	}

	return &newRecord, updatedRecords, nil
}
