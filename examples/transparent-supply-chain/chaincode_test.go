package transparent_supply_chain_test

import (
	"encoding/json"
	"fmt"
	"github.com/LimeChain/SupplyChainPOCs/constants"
	examplesConstants "github.com/LimeChain/SupplyChainPOCs/examples/constants"
	"github.com/LimeChain/SupplyChainPOCs/examples/transparent-supply-chain"
	"github.com/LimeChain/SupplyChainPOCs/types/asset"
	"github.com/LimeChain/SupplyChainPOCs/types/dto"
	"github.com/LimeChain/SupplyChainPOCs/types/order"
	"github.com/LimeChain/SupplyChainPOCs/types/record"
	"github.com/LimeChain/SupplyChainPOCs/utils"
	"github.com/hyperledger/fabric/core/chaincode/shim"
	"github.com/hyperledger/fabric/protos/peer"
	"github.com/shopspring/decimal"
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestSupplyChainPOCs(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "TransparentSupplyChainChaincode Suite")
}

var _ = Describe("Tests for TransparentSupplyChainChaincode", func() {

	stub := shim.NewMockStub("testingStub", new(transparent_supply_chain.TSCChaincode))
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
			result := stub.MockInit("000", [][]byte{
				[]byte(constants.Init)})

			Expect(result.Status).To(Equal(constants.Status500))
			Expect(result.Message).To(Equal(constants.ErrorArgumentsLength))
		})
	})

	Describe("invoke", func() {

		Describe("invalid function name", func() {
			It("Should unsuccessfully execute invoke due to invalid function name", func() {
				result = stub.MockInvoke("000", [][]byte{
					[]byte(constants.ExampleTest)})

				Expect(result.Status).To(Equal(constants.Status500))

				expectedMessage := fmt.Sprintf(constants.ErrorInvalidFunctionName, constants.ExampleTest)
				Expect(result.Message).To(Equal(expectedMessage))
			})
		})

		Describe("addAsset", func() {
			var assetDto dto.AssetDto
			var payload asset.Asset

			BeforeEach(func() {
				assetDto = dto.AssetDto{
					Description: constants.ExampleDescription,
					IsActive:    true}

				jsonAsset, _ := json.Marshal(assetDto)

				result = stub.MockInvoke("000", [][]byte{
					[]byte(constants.AddAsset),
					jsonAsset})
			})

			It("Should successfully execute addAsset", func() {
				Expect(result.Status).To(Equal(constants.Status200))

				payload = asset.Asset{}
				json.Unmarshal(result.Payload, &payload)

				Expect(payload.Description).To(Equal(assetDto.Description))
				Expect(payload.IsActive).To(Equal(assetDto.IsActive))
			})

			It("Should unsuccessfully execute addAsset due to invalid arguments length", func() {
				result = stub.MockInvoke("000", [][]byte{
					[]byte(constants.AddAsset)})

				Expect(result.Status).To(Equal(constants.Status500))
				Expect(result.Message).To(Equal(constants.ErrorArgumentsLength))
			})

			It("Should unsuccessfully execute addAsset due to invalid argument", func() {
				result = stub.MockInvoke("000", [][]byte{
					[]byte(constants.AddAsset),
					[]byte(constants.ExampleTest)})

				Expect(result.Status).To(Equal(constants.Status500))
			})
		})

		Describe("manufacture functionality", func() {
			var payload record.AssetBoundRecord
			var recordDto dto.AssetBoundRecordDto

			It("Should successfully execute manufacture", func() {
				assetDto := dto.AssetDto{
					Description: constants.ExampleDescription,
					IsActive:    true}

				assetPayload := utils.CreateAsset(stub, &assetDto)

				recordDto = dto.AssetBoundRecordDto{
					BaseRecordDto: &dto.BaseRecordDto{
						BatchId:  constants.ExampleBatchId,
						Owner:    constants.OrgOne,
						Quantity: constants.ExampleQuantity,
					},
					AssetId: assetPayload.Id,
				}

				jsonRecord, _ := json.Marshal(recordDto)
				result = stub.MockInvoke("000", [][]byte{
					[]byte(examplesConstants.Manufacture),
					jsonRecord})

				Expect(result.Status).To(Equal(constants.Status200))

				payload = record.AssetBoundRecord{}
				json.Unmarshal(result.Payload, &payload)

				Expect(payload.AssetId).To(Equal(recordDto.AssetId))
				Expect(payload.BatchId).To(Equal(recordDto.BatchId))
				Expect(payload.Owner).To(Equal(recordDto.Owner))
				Expect(payload.Quantity).To(Equal(recordDto.Quantity))
			})

			It("Should unsuccessfully execute manufacture due to invalid AssetId", func() {
				recordDto = dto.AssetBoundRecordDto{
					BaseRecordDto: &dto.BaseRecordDto{
						BatchId:  constants.ExampleBatchId,
						Owner:    constants.OrgOne,
						Quantity: constants.ExampleQuantity,
					},
					AssetId: constants.ExampleAssetId,
				}

				jsonRecord, _ := json.Marshal(recordDto)
				result = stub.MockInvoke("000", [][]byte{
					[]byte(examplesConstants.Manufacture),
					jsonRecord})

				Expect(result.Status).To(Equal(constants.Status500))

				expectedMessage := fmt.Sprintf(constants.ErrorAssetIdNotFound, recordDto.AssetId)
				Expect(result.Message).To(Equal(expectedMessage))
			})

			It("Should unsuccessfully execute manufacture due to invalid arguments length", func() {
				result = stub.MockInvoke("000", [][]byte{
					[]byte(examplesConstants.Manufacture)})

				Expect(result.Status).To(Equal(constants.Status500))
				Expect(result.Message).To(Equal(constants.ErrorArgumentsLength))
			})

			It("Should unsuccessfully execute manufacture due to invalid argument", func() {
				result = stub.MockInvoke("000", [][]byte{
					[]byte(examplesConstants.Manufacture),
					[]byte(constants.ExampleTest)})

				Expect(result.Status).To(Equal(constants.Status500))
			})
		})

		Describe("placeOrder functionality", func() {
			var payload order.FullOrder
			var assetBoundOrderDto dto.AssetBoundOrderDto
			var assetPayload asset.Asset

			BeforeEach(func() {
				assetDto := dto.AssetDto{
					Description: constants.ExampleDescription,
					IsActive:    true}

				assetPayload = utils.CreateAsset(stub, &assetDto)
			})

			It("Should successfully execute placeOrder", func() {

				assetBoundOrderDto = dto.AssetBoundOrderDto{
					AssetId: assetPayload.Id,
					OrderDto: &dto.OrderDto{
						SellerId:     constants.OrgOne,
						BuyerId:      constants.OrgTwo,
						Quantity:     constants.ExampleQuantity,
						PricePerUnit: decimal.NewFromInt(constants.ExamplePrice),
					}}

				jsonOrder, _ := json.Marshal(assetBoundOrderDto)
				result = stub.MockInvoke("000", [][]byte{
					[]byte(constants.PlaceOrder),
					jsonOrder})

				Expect(result.Status).To(Equal(constants.Status200))

				payload = order.FullOrder{}
				json.Unmarshal(result.Payload, &payload)

				Expect(payload.AssetId).To(Equal(assetBoundOrderDto.AssetId))
				Expect(payload.SellerId).To(Equal(assetBoundOrderDto.SellerId))
				Expect(payload.BuyerId).To(Equal(assetBoundOrderDto.BuyerId))
				Expect(payload.Quantity).To(Equal(assetBoundOrderDto.Quantity))
				Expect(payload.PricePerUnit).To(Equal(assetBoundOrderDto.PricePerUnit))
				Expect(payload.IsCompleted).To(Equal(false))
			})

			It("Should unsuccessfully execute placecOrder due to invalid arguments length", func() {
				result = stub.MockInvoke("000", [][]byte{
					[]byte(constants.PlaceOrder)})

				Expect(result.Status).To(Equal(constants.Status500))
				Expect(result.Message).To(Equal(constants.ErrorArgumentsLength))
			})

			It("Should unsuccessfully execute placeOrder due to invalid argument", func() {
				result = stub.MockInvoke("000", [][]byte{
					[]byte(constants.PlaceOrder),
					[]byte(constants.ExampleTest)})

				Expect(result.Status).To(Equal(constants.Status500))
			})

			It("Should unsuccessfully execute placeOrder due to invalid assetId", func() {
				assetBoundOrderDto = dto.AssetBoundOrderDto{
					AssetId: constants.ExampleAssetId,
					OrderDto: &dto.OrderDto{
						SellerId:     constants.OrgOne,
						BuyerId:      constants.OrgTwo,
						Quantity:     constants.ExampleQuantity,
						PricePerUnit: decimal.NewFromInt(constants.ExamplePrice)}}

				jsonOrder, _ := json.Marshal(assetBoundOrderDto)
				result = stub.MockInvoke("000", [][]byte{
					[]byte(constants.PlaceOrder),
					jsonOrder})

				Expect(result.Status).To(Equal(constants.Status500))

				expectedMessage := fmt.Sprintf(constants.ErrorAssetIdNotFound, assetBoundOrderDto.AssetId)
				Expect(result.Message).To(Equal(expectedMessage))
			})
		})

		Describe("fulfillOrder functionality", func() {
			var orderStruct order.BaseOrder
			var orderFulfillment dto.OrderFulfillmentDto
			var assetPayload asset.Asset
			var jsonOrderFulfillment []byte

			BeforeEach(func() {
				assetDto := dto.AssetDto{
					Description: constants.ExampleDescription,
					IsActive:    true}

				assetPayload = utils.CreateAsset(stub, &assetDto)
				orderStruct = utils.CreateAssetBoundOrder(stub, assetPayload.Id)

				orderFulfillment = dto.OrderFulfillmentDto{
					Id:     orderStruct.Id,
					Status: true}
				jsonOrderFulfillment, _ = json.Marshal(orderFulfillment)
			})

			It("Should unsuccessfully execute fulfillOrder due to invalid order id", func() {
				orderFulfillment = dto.OrderFulfillmentDto{
					Id: constants.ExampleTest}

				jsonOrderFulfillment, _ := json.Marshal(orderFulfillment)

				result = stub.MockInvoke("000", [][]byte{
					[]byte(constants.FulfillOrder),
					jsonOrderFulfillment})

				Expect(result.Status).To(Equal(constants.Status500))

				expectedMessage := fmt.Sprintf(constants.ErrorOrderIdNotFound, constants.ExampleTest)
				Expect(result.Message).To(Equal(expectedMessage))
			})

			It("Should unsuccessfully execute fulfillOrder due to invalid arguments length", func() {
				result = stub.MockInvoke("000", [][]byte{
					[]byte(constants.FulfillOrder)})

				Expect(result.Status).To(Equal(constants.Status500))
				Expect(result.Message).To(Equal(constants.ErrorArgumentsLength))
			})

			It("Should unsuccessfully execute fulfillOrder due to invalid argument", func() {
				result = stub.MockInvoke("000", [][]byte{
					[]byte(constants.FulfillOrder),
					[]byte(constants.ExampleTest)})

				Expect(result.Status).To(Equal(constants.Status500))
			})

			It("Should successfully execute fulfillOrder with no record", func() {
				result = stub.MockInvoke("000", [][]byte{
					[]byte(constants.FulfillOrder),
					jsonOrderFulfillment})

				Expect(result.Status).To(Equal(constants.Status200))
			})

			It("Should unsuccessfully execute fulfillOrder due to already completed", func() {
				result = stub.MockInvoke("000", [][]byte{
					[]byte(constants.FulfillOrder),
					jsonOrderFulfillment})

				Expect(result.Status).To(Equal(constants.Status200))

				result = stub.MockInvoke("000", [][]byte{
					[]byte(constants.FulfillOrder),
					jsonOrderFulfillment})

				Expect(result.Status).To(Equal(constants.Status500))

				expectedMessage := fmt.Sprintf(constants.ErrorOrderIsFulfilled, orderStruct.Id)
				Expect(result.Message).To(Equal(expectedMessage))
			})

			It("Should unsuccessfully execute fulfillOrder due to already completed", func() {
				orderFulfillment = dto.OrderFulfillmentDto{

					Id:     orderStruct.Id,
					Status: false,
				}
				jsonOrderFulfillment, _ := json.Marshal(orderFulfillment)

				result = stub.MockInvoke("000", [][]byte{
					[]byte(constants.FulfillOrder),
					jsonOrderFulfillment})

				Expect(result.Status).To(Equal(constants.Status500))

				expectedMessage := fmt.Sprintf(constants.ErrorOrderIsNotFulfilled, orderStruct.Id)
				Expect(result.Message).To(Equal(expectedMessage))
			})
		})

		Describe("assemble functionality", func() {
			var assetComposeRequest dto.AssetComposeRequestDto
			var recordPayload record.BaseRecord
			var assetPayload asset.Asset

			BeforeEach(func() {
				asset := dto.AssetDto{
					Description: constants.ExampleDescription,
					IsActive:    true}

				assetPayload = utils.CreateAsset(stub, &asset)

				recordDto := dto.AssetBoundRecordDto{
					AssetId: assetPayload.Id,
					BaseRecordDto: &dto.BaseRecordDto{
						BatchId:  constants.ExampleBatchId,
						Owner:    constants.OrgOne,
						Quantity: constants.ExampleQuantity,
					}}

				jsonRecordDto, _ := json.Marshal(recordDto)

				result := stub.MockInvoke("000", [][]byte{
					[]byte(examplesConstants.Manufacture),
					jsonRecordDto,
				})

				json.Unmarshal(result.Payload, &recordPayload)

				assetPayload = utils.CreateAsset(stub, &asset)
			})

			It("Should unsuccessfully execute assemble due to invalid arguments length", func() {
				result = stub.MockInvoke("000", [][]byte{
					[]byte(examplesConstants.Assemble)})

				Expect(result.Status).To(Equal(constants.Status500))
				Expect(result.Message).To(Equal(constants.ErrorArgumentsLength))
			})

			It("Should unsuccessfully execute assemble due to invalid argument", func() {
				result = stub.MockInvoke("000", [][]byte{
					[]byte(examplesConstants.Assemble),
					[]byte(constants.ExampleTest)})

				Expect(result.Status).To(Equal(constants.Status500))
			})

			It("Should unsuccessfully execute assemble due to invalid asset id", func() {
				assetComposeRequest = dto.AssetComposeRequestDto{
					AssetId: constants.ExampleAssetId,
					ComposeRequestDto: &dto.ComposeRequestDto{
						BatchId:  constants.ExampleBatchId,
						Quantity: constants.ExampleQuantity,
						Records:  nil,
					},
				}

				jsonComposeRequest, _ := json.Marshal(assetComposeRequest)

				result = stub.MockInvoke("000", [][]byte{
					[]byte(examplesConstants.Assemble),
					jsonComposeRequest})

				Expect(result.Status).To(Equal(constants.Status500))

				expectedMessage := fmt.Sprintf(constants.ErrorAssetIdNotFound, assetComposeRequest.AssetId)
				Expect(result.Message).To(Equal(expectedMessage))
			})

			It("Should unsuccessfully execute compose due to invalid record id", func() {
				assetComposeRequest = dto.AssetComposeRequestDto{
					AssetId: assetPayload.Id,
					ComposeRequestDto: &dto.ComposeRequestDto{
						BatchId:  constants.ExampleBatchId,
						Quantity: constants.ExampleQuantity,
						Records: dto.RecordPartsDto{
							{
								Id:       constants.ExampleRecordId,
								Quantity: constants.ExampleQuantity}},
					},
				}

				jsonComposeRequest, _ := json.Marshal(assetComposeRequest)

				result = stub.MockInvoke("000", [][]byte{
					[]byte(examplesConstants.Assemble),
					jsonComposeRequest})

				Expect(result.Status).To(Equal(constants.Status500))

				expectedMessage := fmt.Sprintf(constants.ErrorRecordIdNotFound, assetComposeRequest.Records[0].Id)
				Expect(result.Message).To(Equal(expectedMessage))
			})

			It("Should unsuccessfully execute assemble due to invalid record quantity", func() {
				assetComposeRequest = dto.AssetComposeRequestDto{
					AssetId: assetPayload.Id,
					ComposeRequestDto: &dto.ComposeRequestDto{
						BatchId:  constants.ExampleBatchId,
						Quantity: constants.ExampleQuantity,
						Records: dto.RecordPartsDto{
							{
								Id:       recordPayload.Id,
								Quantity: recordPayload.Quantity + 1}},
					}}

				jsonComposeRequest, _ := json.Marshal(assetComposeRequest)

				result = stub.MockInvoke("000", [][]byte{
					[]byte(examplesConstants.Assemble),
					jsonComposeRequest})

				Expect(result.Status).To(Equal(constants.Status500))

				expectedMessage := fmt.Sprintf(constants.ErrorRecordQuantity, assetComposeRequest.Records[0].Id)
				Expect(result.Message).To(Equal(expectedMessage))
			})

			It("Should successfully execute assemble", func() {
				asset := dto.AssetDto{
					Description: constants.ExampleDescription,
					IsActive:    true}

				assetPayloadTwo := utils.CreateAsset(stub, &asset)

				record2Dto := dto.AssetBoundRecordDto{
					AssetId: assetPayloadTwo.Id,
					BaseRecordDto: &dto.BaseRecordDto{
						BatchId:  constants.ExampleBatchId,
						Owner:    constants.OrgOne,
						Quantity: constants.ExampleQuantity,
					}}

				jsonRecordDto, _ := json.Marshal(record2Dto)

				result := stub.MockInvoke("000", [][]byte{
					[]byte(examplesConstants.Manufacture),
					jsonRecordDto,
				})
				var recordPayloadTwo record.BaseRecord
				json.Unmarshal(result.Payload, &recordPayloadTwo)

				assetComposeRequest = dto.AssetComposeRequestDto{
					AssetId: assetPayload.Id,
					ComposeRequestDto: &dto.ComposeRequestDto{
						BatchId:  constants.ExampleBatchId,
						Quantity: constants.ExampleQuantity,
						Records: dto.RecordPartsDto{
							{
								Id:       recordPayload.Id,
								Quantity: recordPayload.Quantity},
							{
								Id:       recordPayloadTwo.Id,
								Quantity: recordPayloadTwo.Quantity}},
					}}

				jsonComposeRequest, _ := json.Marshal(assetComposeRequest)

				result = stub.MockInvoke("000", [][]byte{
					[]byte(examplesConstants.Assemble),
					jsonComposeRequest})

				Expect(result.Status).To(Equal(constants.Status200))

				payload := record.ComposableRecord{}
				json.Unmarshal(result.Payload, &payload)
				Expect(payload.Quantity).To(Equal(assetComposeRequest.Quantity))
				for index, _ := range payload.ComposedFrom {
					Expect(payload.ComposedFrom[index]).To(Equal(assetComposeRequest.Records[index]))
				}
			})
		})

		Describe("sell functionality", func() {
			var sellRequest dto.SellDto
			var recordPayload record.BaseRecord

			BeforeEach(func() {
				assetDto := dto.AssetDto{
					Description: constants.ExampleDescription,
					IsActive:    true}

				assetPayload := utils.CreateAsset(stub, &assetDto)

				recordDto := dto.AssetBoundRecordDto{
					AssetId: assetPayload.Id,
					BaseRecordDto: &dto.BaseRecordDto{
						BatchId:  constants.ExampleBatchId,
						Owner:    constants.OrgOne,
						Quantity: constants.ExampleQuantity,
					}}

				jsonRecordDto, _ := json.Marshal(recordDto)

				result := stub.MockInvoke("000", [][]byte{
					[]byte(examplesConstants.Manufacture),
					jsonRecordDto,
				})

				json.Unmarshal(result.Payload, &recordPayload)
			})

			It("Should unsuccessfully execute sell due to invalid arguments length", func() {
				result = stub.MockInvoke("000", [][]byte{
					[]byte(constants.Sell)})

				Expect(result.Status).To(Equal(constants.Status500))
				Expect(result.Message).To(Equal(constants.ErrorArgumentsLength))
			})

			It("Should unsuccessfully execute sell due to invalid argument", func() {
				result = stub.MockInvoke("000", [][]byte{
					[]byte(constants.Sell),
					[]byte(constants.ExampleTest)})

				Expect(result.Status).To(Equal(constants.Status500))
			})

			It("Should unsuccessfully execute sell due to invalid record Id", func() {
				sellRequest = dto.SellDto{
					RecordId: constants.ExampleRecordId,
					Quantity: constants.ExampleQuantity}
				jsonSellRequest, _ := json.Marshal(sellRequest)

				result = stub.MockInvoke("000", [][]byte{
					[]byte(constants.Sell),
					jsonSellRequest})

				Expect(result.Status).To(Equal(constants.Status500))

				expectedMessage := fmt.Sprintf(constants.ErrorRecordIdNotFound, sellRequest.RecordId)
				Expect(result.Message).To(Equal(expectedMessage))
			})

			It("Should unsuccessfully execute sell due to insufficient record quantity", func() {
				sellRequest = dto.SellDto{
					RecordId: recordPayload.Id,
					Quantity: recordPayload.Quantity + 1}

				jsonSellRequest, _ := json.Marshal(sellRequest)

				result = stub.MockInvoke("000", [][]byte{
					[]byte(constants.Sell),
					jsonSellRequest})

				Expect(result.Status).To(Equal(constants.Status500))

				expectedMessage := fmt.Sprintf(constants.ErrorRecordQuantity, sellRequest.RecordId)
				Expect(result.Message).To(Equal(expectedMessage))
			})

			It("Should successfully execute sell", func() {
				sellRequest = dto.SellDto{
					RecordId: recordPayload.Id,
					Quantity: recordPayload.Quantity}

				jsonSellRequest, _ := json.Marshal(sellRequest)

				result = stub.MockInvoke("000", [][]byte{
					[]byte(constants.Sell),
					jsonSellRequest})

				Expect(result.Status).To(Equal(constants.Status200))

				payload := record.BaseRecord{}
				json.Unmarshal(result.Payload, &payload)

				Expect(payload.Quantity).To(Equal(uint64(0)))
			})
		})
	})
})
