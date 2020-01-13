package types

import (
	"time"
)

type Record struct {
	Id                 string          `json:"id,omitempty"`
	AssetId            string          `json:"assetId,omitempty"`
	BatchId            string          `json:"batchId,omitempty"`
	CreationOrderId    string          `json:"creationOrderId,omitempty"`
	Owner              string          `json:"owner,omitempty"`
	Quantity           uint64          `json:"quantity,omitempty"`
	DateCreated        time.Time       `json:"dateCreated,omitempty"`
	LastUpdated        time.Time       `json:"lastUpdated,omitempty"`
	AssembledFrom      []AssetAssembly `json:"assembledFrom,omitempty"`
	QualityCertificate string          `json:"qualityCertificate,omitempty"`
}
