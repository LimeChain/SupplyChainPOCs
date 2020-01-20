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

type ComposableChaincode struct {
	BaseSupplyChainChaincode
}

func (ccc *ComposableChaincode) Create(id string, composableRecordDto *dto.ComposableRecordDto) *record.ComposableRecord {
	rec := ccc.BaseSupplyChainChaincode.Create(id, composableRecordDto.BaseRecordDto)

	records := record.RecordParts(composableRecordDto.ComposedFrom)
	return record.NewComposableRecord(rec, records)
}

func (ccc *ComposableChaincode) Compose(stub shim.ChaincodeStubInterface, id string, composeRequest *dto.ComposeRequestDto) (*record.ComposableRecord, record.RecordParts, error) {
	newRecord := record.ComposableRecord{
		BaseRecord: &record.BaseRecord{
			Id:          id,
			BatchId:     composeRequest.BatchId,
			Owner:       utils.GetOrganization(stub, constants.Org2Index),
			Quantity:    composeRequest.Quantity,
			DateCreated: time.Now(),
		},
		ComposedFrom: record.RecordParts{}}

	updatedRecords := record.RecordParts{}

	for _, recordElem := range composeRequest.Records {
		recordBytes, _ := stub.GetState(recordElem.Id)

		if len(recordBytes) == 0 {
			return nil, nil, errors.New(fmt.Sprintf(constants.ErrorRecordIdNotFound, recordElem.Id))
		}

		recordStruct := record.BaseRecord{}
		err := json.Unmarshal(recordBytes, &recordStruct)

		if err != nil {
			return nil, nil, err
		}

		if recordElem.Quantity > recordStruct.Quantity {
			return nil, nil, errors.New(fmt.Sprintf(constants.ErrorRecordQuantity, recordElem.Id))
		}

		updatedRecords = append(updatedRecords, record.RecordParts{
			{Id: recordStruct.Id, Quantity: recordStruct.GetNewQuantity(recordElem.Quantity)}}...)

		newRecord.ComposedFrom = append(newRecord.ComposedFrom, recordElem)
	}

	return &newRecord, updatedRecords, nil
}
