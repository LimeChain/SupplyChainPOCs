package cc

import (
	"github.com/LimeChain/SupplyChainPOCs/types/asset"
	"github.com/LimeChain/SupplyChainPOCs/types/dto"
	"github.com/LimeChain/SupplyChainPOCs/types/order"
	"github.com/LimeChain/SupplyChainPOCs/types/record"
)

type AssetBoundChaincode struct {
	BaseSupplyChainChaincode
}

func (abcc *AssetBoundChaincode) AddAsset(id string, assetDto *dto.AssetDto) *asset.Asset {
	return asset.NewAsset(id, assetDto)
}

func (abcc *AssetBoundChaincode) Create(id string, assetBoundRecordDto *dto.AssetBoundRecordDto) *record.AssetBoundRecord {
	rec := abcc.BaseSupplyChainChaincode.Create(id, assetBoundRecordDto.BaseRecordDto)

	return record.NewAssetBoundRecord(rec, assetBoundRecordDto.AssetId)
}

func (abcc *AssetBoundChaincode) PlaceOrder(id string, assetBoundOrderDto *dto.AssetBoundOrderDto) *order.AssetBoundOrder {
	newOrder := abcc.BaseSupplyChainChaincode.PlaceOrder(id, assetBoundOrderDto.OrderDto)

	return order.NewAssetBoundOrder(newOrder, assetBoundOrderDto.AssetId)
}
