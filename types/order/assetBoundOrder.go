package order

type AssetBoundOrder struct {
	*BaseOrder
	AssetId string `json:"assetId,omitempty"`
}

func NewAssetBoundOrder(order *BaseOrder, assetId string) *AssetBoundOrder {
	return &AssetBoundOrder{
		BaseOrder: order,
		AssetId:   assetId,
	}
}
