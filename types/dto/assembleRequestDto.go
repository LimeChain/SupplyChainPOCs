package dto

type AssembleRequestDto struct {
	AssetId  string         `json:"assetId,omitempty"`
	BatchId  string         `json:"batchId,omitempty"`
	Quantity uint64         `json:"quantity,omitempty"`
	Records  RecordPartsDto `json:"records,omitempty"`
}
