package dto

type BaseRecordDto struct {
	BatchId  string `json:"batchId,omitempty"`
	Owner    string `json:"owner,omitempty"`
	Quantity uint64 `json:"quantity,omitempty"`
}
