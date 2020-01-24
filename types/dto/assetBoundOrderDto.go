package dto

type AssetBoundOrderDto struct {
	AssetId string `json:"assetId,omitempty"`
	*OrderDto
}
