package order

type AssetBoundOrder struct {
	*Order
	AssetId string `json:"assetId,omitempty"`
}

func NewAssetBoundOrder(order *Order, assetId string) *AssetBoundOrder {
	return &AssetBoundOrder{
		Order:   order,
		AssetId: assetId,
	}
}
