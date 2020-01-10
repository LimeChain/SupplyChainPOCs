package poc01_test

import (
	"encoding/json"
	"fmt"
	"github.com/hyperledger/fabric/core/chaincode/shim"
	"github.com/hyperledger/fabric/protos/peer"
	"github.com/shopspring/decimal"
	"testing"
	
	"github.com/LimeChain/SupplyChainPOCs"
	"github.com/LimeChain/SupplyChainPOCs/constants"
	"github.com/LimeChain/SupplyChainPOCs/types"
	"github.com/LimeChain/SupplyChainPOCs/utils"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestSupplyChainPOCs(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "SupplyChainPOCs Suite")
}

var _ = Describe("Tests for SupplyChainChaincode", func() {
	
	stub := shim.NewMockStub("testingStub", new(poc01.SupplyChainChaincode))
	var result peer.Response

	BeforeSuite(func() {
		stub.MockInit("000", [][]byte{
			[]byte(constants.Init),
			[]byte(constants.OrgOne),
			[]byte(constants.OrgTwo),
			[]byte(constants.OrgThree)})
	})

	Describe("init", func() {
		It("Should fail due to invalid arguments length", func() {
			result := stub.MockInit("000", [][]byte {
				[]byte(constants.Init)})

			Expect(result.Status).To(Equal(constants.Status500))
			Expect(result.Message).To(Equal(constants.ErrorArgumentsLength))
		})
	})

	Describe("invoke", func() {

		Describe("invalid function name", func() {
			It("Should unsuccessfully execute invoke due to invalid function name", func() {
				result = stub.MockInvoke("000", [][]byte{
					[]byte(constants.ExampleTest) })

				Expect(result.Status).To(Equal(constants.Status500))

				expectedMessage := fmt.Sprintf(constants.ErrorInvalidFunctionName, constants.ExampleTest)
				Expect(result.Message).To(Equal(expectedMessage))
			})
		})

		Describe("addAssetType", func() {
			var asset types.Asset
			var payload types.Asset

			BeforeEach(func () {
				asset = types.Asset {
					Description: constants.ExampleDescription,
					IsActive: true }

				jsonAsset, _ := json.Marshal(asset)

				result = stub.MockInvoke("000", [][]byte{
					[]byte(constants.AddAssetType),
					jsonAsset })
			})

			It("Should successfully execute addAssetType", func() {
				Expect(result.Status).To(Equal(constants.Status200))

				payload = types.Asset {}
				json.Unmarshal(result.Payload, &payload)

				Expect(payload.Description).To(Equal(asset.Description))
				Expect(payload.IsActive).To(Equal(asset.IsActive))
			})

			It("Should unsuccessfully execute addAssetType due to invalid arguments length", func() {
				result = stub.MockInvoke("000", [][]byte {
					[]byte(constants.AddAssetType)})

				Expect(result.Status).To(Equal(constants.Status500))
				Expect(result.Message).To(Equal(constants.ErrorArgumentsLength))
			})

			It("Should unsuccessfully execute addAssetType due to invalid argument", func() {
				result = stub.MockInvoke("000", [][]byte {
					[]byte(constants.AddAssetType),
					[]byte(constants.ExampleTest)})

				Expect(result.Status).To(Equal(constants.Status500))
			})
		})

		Describe("manufacture functionality", func() {
			var payload types.Record
			var record types.Record

			It("Should successfully execute manufacture", func() {
				asset := types.Asset {
					Description: constants.ExampleDescription,
					IsActive: true }

				assetPayload := utils.CreateAsset(stub, &asset)

				record = types.Record {
					AssetId: assetPayload.Id,
					BatchId: constants.ExampleBatchId,
					Owner: constants.OrgOne,
					Quantity: constants.ExampleQuantity,
					QualityCertificate: "" }

				jsonRecord, _ := json.Marshal(record)
				result = stub.MockInvoke("000", [][]byte{
					[]byte(constants.Manufacture),
					jsonRecord })

				Expect(result.Status).To(Equal(constants.Status200))

				payload = types.Record {}
				json.Unmarshal(result.Payload, &payload)

				Expect(payload.AssetId).To(Equal(record.AssetId))
				Expect(payload.BatchId).To(Equal(record.BatchId))
				Expect(payload.Owner).To(Equal(record.Owner))
				Expect(payload.Quantity).To(Equal(record.Quantity))
				Expect(payload.QualityCertificate).To(Equal(record.QualityCertificate))
			})

			It("Should unsuccessfully execute manufacture due to invalid AssetId", func() {
				record = types.Record {
					AssetId: constants.ExampleAssetId,
					BatchId: constants.ExampleBatchId,
					Owner: constants.OrgOne,
					Quantity: constants.ExampleQuantity,
					QualityCertificate: "" }

				jsonRecord, _ := json.Marshal(record)
				result = stub.MockInvoke("000", [][]byte{
					[]byte(constants.Manufacture),
					jsonRecord })

				Expect(result.Status).To(Equal(constants.Status500))

				expectedMessage := fmt.Sprintf(constants.ErrorAssetIdNotFound, record.AssetId)
				Expect(result.Message).To(Equal(expectedMessage))
			})

			It("Should unsuccessfully execute manufacture due to invalid arguments length", func() {
				result = stub.MockInvoke("000", [][]byte {
					[]byte(constants.Manufacture)})

				Expect(result.Status).To(Equal(constants.Status500))
				Expect(result.Message).To(Equal(constants.ErrorArgumentsLength))
			})

			It("Should unsuccessfully execute manufacture due to invalid argument", func() {
				result = stub.MockInvoke("000", [][]byte {
					[]byte(constants.Manufacture),
					[]byte(constants.ExampleTest)})

				Expect(result.Status).To(Equal(constants.Status500))
			})
		})

		Describe("placeOrder functionality", func() {
			var payload types.Order
			var order types.Order
			var assetPayload types.Asset

			BeforeEach(func() {
				asset := types.Asset {
					Description: constants.ExampleDescription,
					IsActive: true }

				assetPayload = utils.CreateAsset(stub, &asset)
			})

			It("Should successfully execute placeOrder", func() {

				order = types.Order {
					AssetId: assetPayload.Id,
					SellerId: constants.OrgOne,
					BuyerId: constants.OrgTwo,
					Quantity: constants.ExampleQuantity,
					PricePerUnit: decimal.NewFromInt(constants.ExamplePrice) }

				jsonOrder, _ := json.Marshal(order)
				result = stub.MockInvoke("000", [][]byte{
					[]byte(constants.PlaceOrder),
					jsonOrder })

				Expect(result.Status).To(Equal(constants.Status200))

				payload = types.Order {}
				json.Unmarshal(result.Payload, &payload)

				Expect(payload.AssetId).To(Equal(order.AssetId))
				Expect(payload.SellerId).To(Equal(order.SellerId))
				Expect(payload.BuyerId).To(Equal(order.BuyerId))
				Expect(payload.Quantity).To(Equal(order.Quantity))
				Expect(payload.PricePerUnit).To(Equal(order.PricePerUnit))
				Expect(payload.IsCompleted).To(Equal(false))
			})

			It("Should unsuccessfully execute placecOrder due to invalid arguments length", func() {
				result = stub.MockInvoke("000", [][]byte {
					[]byte(constants.Manufacture)})

				Expect(result.Status).To(Equal(constants.Status500))
				Expect(result.Message).To(Equal(constants.ErrorArgumentsLength))
			})

			It("Should unsuccessfully execute placeOrder due to invalid argument", func() {
				result = stub.MockInvoke("000", [][]byte {
					[]byte(constants.Manufacture),
					[]byte(constants.ExampleTest)})

				Expect(result.Status).To(Equal(constants.Status500))
			})

			It("Should unsuccessfully execute placeOrder due to invalid assetId", func() {
				order = types.Order {
					AssetId: constants.ExampleAssetId,
					SellerId: constants.OrgOne,
					BuyerId: constants.OrgTwo,
					Quantity: constants.ExampleQuantity,
					PricePerUnit: decimal.NewFromInt(constants.ExamplePrice) }

				jsonOrder, _ := json.Marshal(order)
				result = stub.MockInvoke("000", [][]byte{
					[]byte(constants.PlaceOrder),
					jsonOrder })

				Expect(result.Status).To(Equal(constants.Status500))

				expectedMessage := fmt.Sprintf(constants.ErrorAssetIdNotFound, order.AssetId)
				Expect(result.Message).To(Equal(expectedMessage))
			})
		})

		Describe("fulfillOrder functionality", func() {
			var order types.Order
			var orderFulfillment types.OrderFulfillment
			var assetPayload types.Asset

			BeforeEach(func() {
				asset := types.Asset {
					Description: constants.ExampleDescription,
					IsActive: true }

				assetPayload = utils.CreateAsset(stub, &asset)
			})

			It("Should unsuccessfully execute fulfillOrder due to invalid order id", func() {
				orderFulfillment = types.OrderFulfillment {
					Id: constants.ExampleTest }

				jsonOrderFulfillment, _ := json.Marshal(orderFulfillment)

				result = stub.MockInvoke("000", [][]byte {
					[]byte(constants.FulfillOrder),
					jsonOrderFulfillment })

				Expect(result.Status).To(Equal(constants.Status500))

				expectedMessage := fmt.Sprintf(constants.ErrorOrderIdNotFound, constants.ExampleTest)
				Expect(result.Message).To(Equal(expectedMessage))
			})

			It("Should unsuccessfully execute fulfillOrder due to invalid arguments length", func() {
				result = stub.MockInvoke("000", [][]byte {
					[]byte(constants.FulfillOrder)})

				Expect(result.Status).To(Equal(constants.Status500))
				Expect(result.Message).To(Equal(constants.ErrorArgumentsLength))
			})

			It("Should unsuccessfully execute fulfillOrder due to invalid argument", func() {
				result = stub.MockInvoke("000", [][]byte {
					[]byte(constants.FulfillOrder),
					[]byte(constants.ExampleTest)})

				Expect(result.Status).To(Equal(constants.Status500))
			})

			It("Should successfully execute fulfillOrder with no records", func() {
				order = utils.CreateOrder(stub, assetPayload.Id)

				orderFulfillment = types.OrderFulfillment {
					Id: order.Id}
				jsonOrderFulfillment, _ := json.Marshal(orderFulfillment)

				result = stub.MockInvoke("000", [][]byte {
					[]byte(constants.FulfillOrder),
					jsonOrderFulfillment })

				Expect(result.Status).To(Equal(constants.Status200))
			})

			It("Should unsuccessfully execute fulfillOrder due to already completed", func() {
				order = utils.CreateOrder(stub, assetPayload.Id)

				orderFulfillment = types.OrderFulfillment {
					Id: order.Id}
				jsonOrderFulfillment, _ := json.Marshal(orderFulfillment)

				result = stub.MockInvoke("000", [][]byte {
					[]byte(constants.FulfillOrder),
					jsonOrderFulfillment })

				Expect(result.Status).To(Equal(constants.Status200))

				result = stub.MockInvoke("000", [][]byte {
					[]byte(constants.FulfillOrder),
					jsonOrderFulfillment })

				Expect(result.Status).To(Equal(constants.Status500))

				expectedMessage := fmt.Sprintf(constants.ErrorOrderIsFulfilled, order.Id)
				Expect(result.Message).To(Equal(expectedMessage))
			})
		})

		Describe("assemble functionality", func() {
			var assembleRequest types.AssembleRequest
			var recordPayload types.Record
			var assetPayload types.Asset

			BeforeEach(func() {
				asset := types.Asset {
					Description: constants.ExampleDescription,
					IsActive: true }

				assetPayload = utils.CreateAsset(stub, &asset)

				record := types.Record {
					AssetId: assetPayload.Id,
					BatchId: constants.ExampleBatchId,
					Owner: constants.OrgOne,
					Quantity: constants.ExampleQuantity,
					QualityCertificate: "" }

				recordPayload = utils.CreateRecord(stub, &record)

				assetPayload = utils.CreateAsset(stub, &asset)
			})

			It("Should unsuccessfully execute assemble due to invalid arguments length", func() {
				result = stub.MockInvoke("000", [][]byte {
					[]byte(constants.Assemble)})

				Expect(result.Status).To(Equal(constants.Status500))
				Expect(result.Message).To(Equal(constants.ErrorArgumentsLength))
			})

			It("Should unsuccessfully execute assemble due to invalid argument", func() {
				result = stub.MockInvoke("000", [][]byte {
					[]byte(constants.Assemble),
					[]byte(constants.ExampleTest)})

				Expect(result.Status).To(Equal(constants.Status500))
			})

			It("Should unsuccessfully execute assemble due to invalid asset id", func() {
				assembleRequest = types.AssembleRequest {
					AssetId: constants.ExampleAssetId,
					BatchId: constants.ExampleBatchId,
					Quantity: constants.ExampleQuantity,
				}

				jsonAssembleRequest, _ := json.Marshal(assembleRequest)

				result = stub.MockInvoke("000", [][]byte {
					[]byte(constants.Assemble),
					jsonAssembleRequest })

				Expect(result.Status).To(Equal(constants.Status500))

				expectedMessage := fmt.Sprintf(constants.ErrorAssetIdNotFound, assembleRequest.AssetId)
				Expect(result.Message).To(Equal(expectedMessage))
			})

			It("Should unsuccessfully execute assemble due to invalid record id", func() {
				assembleRequest = types.AssembleRequest {
					AssetId: assetPayload.Id,
					BatchId: constants.ExampleBatchId,
					Quantity: constants.ExampleQuantity,
					Records: []types.AssetAssembly {
						{
							Id: constants.ExampleRecordId,
							Quantity: constants.ExampleQuantity }}}

				jsonAssembleRequest, _ := json.Marshal(assembleRequest)

				result = stub.MockInvoke("000", [][]byte {
					[]byte(constants.Assemble),
					jsonAssembleRequest })

				Expect(result.Status).To(Equal(constants.Status500))

				expectedMessage := fmt.Sprintf(constants.ErrorRecordIdNotFound, assembleRequest.Records[0].Id)
				Expect(result.Message).To(Equal(expectedMessage))
			})

			It("Should unsuccessfully execute assemble due to invalid record quantity", func() {
				assembleRequest = types.AssembleRequest {
					AssetId: assetPayload.Id,
					BatchId: constants.ExampleBatchId,
					Quantity: constants.ExampleQuantity,
					Records: []types.AssetAssembly {
						{
							Id: recordPayload.Id,
							Quantity: recordPayload.Quantity + 1 }}}

				jsonAssembleRequest, _ := json.Marshal(assembleRequest)

				result = stub.MockInvoke("000", [][]byte {
					[]byte(constants.Assemble),
					jsonAssembleRequest })

				Expect(result.Status).To(Equal(constants.Status500))

				expectedMessage := fmt.Sprintf(constants.ErrorRecordQuantity, assembleRequest.Records[0].Id)
				Expect(result.Message).To(Equal(expectedMessage))
			})

			It("Should successfully execute assemble", func() {
				asset := types.Asset {
					Description: constants.ExampleDescription,
					IsActive: true }

				assetPayloadTwo := utils.CreateAsset(stub, &asset)

				record2 := types.Record {
					AssetId: assetPayloadTwo.Id,
					BatchId: constants.ExampleBatchId,
					Owner: constants.OrgOne,
					Quantity: constants.ExampleQuantity,
					QualityCertificate: "" }

				recordPayloadTwo := utils.CreateRecord(stub, &record2)

				assembleRequest = types.AssembleRequest {
					AssetId: assetPayload.Id,
					BatchId: constants.ExampleBatchId,
					Quantity: constants.ExampleQuantity,
					Records: []types.AssetAssembly {
						{
							Id: recordPayload.Id,
							Quantity: recordPayload.Quantity },
						{
							Id: recordPayloadTwo.Id,
							Quantity: recordPayloadTwo.Quantity }}}

				jsonAssembleRequest, _ := json.Marshal(assembleRequest)

				result = stub.MockInvoke("000", [][]byte {
					[]byte(constants.Assemble),
					jsonAssembleRequest })

				Expect(result.Status).To(Equal(constants.Status200))

				payload := types.Record {}
				json.Unmarshal(result.Payload, &payload)
				Expect(payload.AssembledFrom).To(Equal(assembleRequest.Records))
				Expect(payload.Quantity).To(Equal(assembleRequest.Quantity))
			})
		})

		Describe("sell functionality", func() {
			var sellRequest types.SellRequest
			var recordPayload types.Record

			BeforeEach(func() {
				asset := types.Asset {
					Description: constants.ExampleDescription,
					IsActive: true }

				assetPayload := utils.CreateAsset(stub, &asset)

				record := types.Record {
					AssetId: assetPayload.Id,
					BatchId: constants.ExampleBatchId,
					Owner: constants.OrgOne,
					Quantity: constants.ExampleQuantity,
					QualityCertificate: "" }

				recordPayload = utils.CreateRecord(stub, &record)
			})

			It("Should unsuccessfully execute sell due to invalid arguments length", func() {
				result = stub.MockInvoke("000", [][]byte {
					[]byte(constants.Sell)})

				Expect(result.Status).To(Equal(constants.Status500))
				Expect(result.Message).To(Equal(constants.ErrorArgumentsLength))
			})

			It("Should unsuccessfully execute sell due to invalid argument", func() {
				result = stub.MockInvoke("000", [][]byte {
					[]byte(constants.Sell),
					[]byte(constants.ExampleTest)})

				Expect(result.Status).To(Equal(constants.Status500))
			})

			It("Should unsuccessfully execute sell due to invalid record Id", func() {
				sellRequest = types.SellRequest {
					RecordId: constants.ExampleRecordId,
					Quantity: constants.ExampleQuantity }
				jsonSellRequest, _ := json.Marshal(sellRequest)

				result = stub.MockInvoke("000", [][]byte {
					[]byte(constants.Sell),
					jsonSellRequest })

				Expect(result.Status).To(Equal(constants.Status500))

				expectedMessage := fmt.Sprintf(constants.ErrorRecordIdNotFound, sellRequest.RecordId)
				Expect(result.Message).To(Equal(expectedMessage))
			})

			It("Should unsuccessfully execute sell due to insufficient record quantity", func() {
				sellRequest = types.SellRequest {
					RecordId: recordPayload.Id,
					Quantity: recordPayload.Quantity + 1 }

				jsonSellRequest, _ := json.Marshal(sellRequest)

				result = stub.MockInvoke("000", [][]byte {
					[]byte(constants.Sell),
					jsonSellRequest })

				Expect(result.Status).To(Equal(constants.Status500))

				expectedMessage := fmt.Sprintf(constants.ErrorRecordQuantity, sellRequest.RecordId)
				Expect(result.Message).To(Equal(expectedMessage))
			})

			It("Should successfully execute sell", func() {
				sellRequest = types.SellRequest {
					RecordId: recordPayload.Id,
					Quantity: recordPayload.Quantity }

				jsonSellRequest, _ := json.Marshal(sellRequest)

				result = stub.MockInvoke("000", [][]byte {
					[]byte(constants.Sell),
					jsonSellRequest })

				Expect(result.Status).To(Equal(constants.Status200))

				payload := types.Record {}
				json.Unmarshal(result.Payload, &payload)

				Expect(payload.Quantity).To(Equal(uint64(0)))
			})
		})
	})
})
