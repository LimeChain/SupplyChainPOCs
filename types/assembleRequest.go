package types

type AssembleRequest struct {
	AssetId            string          `json:"assetId,omitempty"`
	BatchId            string          `json:"batchId,omitempty"`
	Quantity           uint64          `json:"quantity,omitempty"`
	Records            []AssetAssembly `json:"records,omitempty"`
	QualityCertificate string          `json:"qualityCertificate,omitempty"`
}
