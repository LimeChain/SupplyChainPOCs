package dto

type ComposeRequestDto struct {
	BatchId  string         `json:"batchId,omitempty"`
	Quantity uint64         `json:"quantity,omitempty"`
	Records  RecordPartsDto `json:"records,omitempty"`
}
