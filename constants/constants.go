package constants

// Error Messages
const (
	ErrorArgumentsLength        = "Invalid arguments length"
	ErrorAssetIdNotFound        = "Asset with Id: %s not found"
	ErrorInvalidFunctionName    = "Invalid invoke function name: %s"
	ErrorOrderIdNotFound        = "Order with Id: %s not found"
	ErrorOrderIsFulfilled       = "Order with Id: %s is already fulfilled"
	ErrorOrderIsNotFulfilled    = "Order with Id: %s has not been fulfilled"
	ErrorRecordIdNotFound       = "Record with Id: %s not found"
	ErrorRecordQuantity         = "Insufficient quantity for Record with Id: %s"
	ErrorRecordDifferentAssetId = "Record with Id: %s has AssetId: %s, and not AssetId: %s"
	ErrorStartChaincode         = "Error starting chaincode: %s"
)

// Response Statuses
const (
	Status500 = int32(500)
	Status200 = int32(200)
)

// Organizations
const (
	OrgOne        = "Org1MSP"
	OrgTwo        = "Org2MSP"
	OrgThree      = "Org3MSP"
	Org1Index     = 0
	Org2Index     = 1
	Org3Index     = 2
	Organizations = "organizations"
)

// Chaincode Functions
const (
	AddAssetType = "addAssetType"
	Compose      = "compose"
	Create       = "create"
	Init         = "init"
	FulfillOrder = "fulfillOrder"
	PlaceOrder   = "placeOrder"
	Query        = "query"
	Sell         = "sell"
)

// Prefixes
const (
	PrefixRecord = "record"
	PrefixOrder  = "order"
	PrefixAsset  = "asset"
)

// Test data
const (
	ExampleAssetId     = "assetId"
	ExampleBatchId     = "batchId"
	ExampleDescription = "description"
	ExampleRecordId    = "recordId"
	ExampleTest        = "test"
	ExamplePrice       = 10
	ExampleQuantity    = 2
)
