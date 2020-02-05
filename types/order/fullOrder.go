package order

import "github.com/shopspring/decimal"

type FullOrder struct {
	*BaseOrder
	*AssetBoundOrder
	*PriceOrder
}

func NewFullOrder(order *BaseOrder, pricePerUnit decimal.Decimal, assetId string) *FullOrder {
	return &FullOrder{
		BaseOrder: order,
		AssetBoundOrder: &AssetBoundOrder{
			BaseOrder: nil,
			AssetId:   assetId,
		},
		PriceOrder: &PriceOrder{
			BaseOrder:    nil,
			PricePerUnit: pricePerUnit,
		},
	}
}
