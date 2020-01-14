package dto

type RecordDto struct {
	BatchId            string `json:"batchId,omitempty"`
	Owner              string `json:"owner,omitempty"`
	Quantity           uint64 `json:"quantity,omitempty"`
	QualityCertificate string `json:"qualityCertificate,omitempty"`
}
