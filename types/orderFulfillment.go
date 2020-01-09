package types 

type OrderFulfillment struct {
	Id      string          `json:"id,omitempty"`
	Records []AssetAssembly `json:"records,omitempty"`
}
