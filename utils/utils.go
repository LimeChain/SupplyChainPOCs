package utils

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/LimeChain/SupplyChainPOCs/constants"
	"github.com/LimeChain/SupplyChainPOCs/types/asset"
	"github.com/LimeChain/SupplyChainPOCs/types/dto"
	"github.com/LimeChain/SupplyChainPOCs/types/order"
	guuid "github.com/google/uuid"
	"github.com/hyperledger/fabric/core/chaincode/shim"
	"github.com/shopspring/decimal"
)

// =========================================================================================
// GetQueryResultForQueryString executes the passed in query string.
// Result set is built and returned as a byte array containing the JSON results.
// =========================================================================================
func GetQueryResultForQueryString(stub shim.ChaincodeStubInterface, queryString string) ([] byte, error) {
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
	return stub.CreateCompositeKey(prefix, []string{guuid.New().String()})
}

func GetOrganization(stub shim.ChaincodeStubInterface, index uint) string {
	organizationsBytes, _ := stub.GetState(constants.Organizations)
	var organizations []string
	json.Unmarshal(organizationsBytes, &organizations)

	return organizations[index]
}

func CreateAsset(stub *shim.MockStub, assetDto *dto.AssetDto) asset.Asset {
	jsonAsset, _ := json.Marshal(assetDto)
	result := stub.MockInvoke("000", [][]byte{
		[]byte(constants.AddAsset),
		jsonAsset})

	payload := asset.Asset{}
	json.Unmarshal(result.Payload, &payload)

	return payload
}

func CreateAssetBoundOrder(stub *shim.MockStub, assetId string) order.BaseOrder {
	orderDto := dto.AssetBoundOrderDto{
		AssetId: assetId,
		OrderDto: &dto.OrderDto{
			SellerId:     constants.OrgOne,
			BuyerId:      constants.OrgTwo,
			Quantity:     constants.ExampleQuantity,
			PricePerUnit: decimal.NewFromInt(constants.ExamplePrice),
		}}

	jsonOrder, _ := json.Marshal(orderDto)
	result := stub.MockInvoke("000", [][]byte{
		[]byte(constants.PlaceOrder),
		jsonOrder})

	payload := order.BaseOrder{}
	json.Unmarshal(result.Payload, &payload)

	return payload
}
