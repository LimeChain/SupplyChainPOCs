package types

type AssembleRequest struct {
	AssetId  string          `json:"assetId,omitempty"`
	Quantity uint64          `json:"quantity,omitempty"`
	Assets   []AssetAssembly `json:"assets,omitempty"`
}