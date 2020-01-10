package utils

import (
	"encoding/json"
	"bytes"
	"fmt"
	"github.com/hyperledger/fabric/core/chaincode/shim"
	"github.com/LimeChain/SupplyChainPOCs/constants"
	"github.com/LimeChain/SupplyChainPOCs/types"
	"github.com/shopspring/decimal"
	guuid "github.com/google/uuid"
)

// =========================================================================================
// GetQueryResultForQueryString executes the passed in query string.
// Result set is built and returned as a byte array containing the JSON results.
// =========================================================================================
func GetQueryResultForQueryString(stub shim.ChaincodeStubInterface, queryString string)([] byte, error) {
	fmt.Printf("- getQueryResultForQueryString queryString:\n%s\n", queryString)
	resultsIterator, err := stub.GetQueryResult(queryString)
	fmt.Println(resultsIterator)
	defer resultsIterator.Close()
	if err != nil {
			return nil, err
	}
	// buffer is a JSON array containing QueryRecords
	var buffer bytes.Buffer
	buffer.WriteString("[")
	bArrayMemberAlreadyWritten := false
	for resultsIterator.HasNext() {
			queryResponse,
			err := resultsIterator.Next()
			if err != nil {
					return nil, err
			}
			// Add a comma before array members, suppress it for the first array member
			if bArrayMemberAlreadyWritten == true {
					buffer.WriteString(",")
			}
			buffer.WriteString("{\"Key\":")
			buffer.WriteString("\"")
			buffer.WriteString(queryResponse.Key)
			buffer.WriteString("\"")
			buffer.WriteString(", \"Record\":")
			// Record is a JSON object, so we write as-is
			buffer.WriteString(string(queryResponse.Value))
			buffer.WriteString("}")
			bArrayMemberAlreadyWritten = true
	}
	buffer.WriteString("]")
	fmt.Printf("- getQueryResultForQueryString queryResult:\n%s\n", buffer.String())
	return buffer.Bytes(), nil
}

func CreateCompositeKey(stub shim.ChaincodeStubInterface, prefix string) (string, error) {
	return (stub.CreateCompositeKey(prefix, []string{guuid.New().String()}))
}

func GetOrganization(stub shim.ChaincodeStubInterface, index uint) string {
	organizationsBytes, _ := stub.GetState(constants.Organizations)
	var organizations []string
	json.Unmarshal(organizationsBytes, &organizations)

	return organizations[index]
}

func CreateAsset(stub *shim.MockStub, asset *types.Asset) types.Asset {
	jsonAsset, _ := json.Marshal(asset)
	result := stub.MockInvoke("000", [][]byte{
		[]byte(constants.AddAssetType),
		jsonAsset})
	
	payload := types.Asset {}
	json.Unmarshal(result.Payload, &payload)

	return payload
}

func CreateRecord(stub *shim.MockStub, record *types.Record) types.Record {
	jsonRecord, _ := json.Marshal(record)
	result := stub.MockInvoke("000", [][]byte{
		[]byte(constants.Manufacture),
		jsonRecord})
	
	payload := types.Record {}
	json.Unmarshal(result.Payload, &payload)

	return payload
}

func CreateOrder(stub *shim.MockStub) types.Order {
	order := types.Order {
		AssetId: constants.ExampleAssetId,
		SellerId: constants.ORG_ONE,
		BuyerId: constants.ORG_TWO,
		Quantity: constants.ExampleQuantity,
		PricePerUnit: decimal.NewFromInt(constants.ExamplePrice) }

	jsonOrder, _ := json.Marshal(order)
	result := stub.MockInvoke("000", [][]byte{
		[]byte(constants.PlaceOrder),
		jsonOrder})

	payload := types.Order {}
	json.Unmarshal(result.Payload, &payload)

	return payload
}
