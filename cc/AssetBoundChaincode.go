package cc

import (
	"github.com/LimeChain/SupplyChainPOCs/types/asset"
	"github.com/LimeChain/SupplyChainPOCs/types/dto"
	"github.com/LimeChain/SupplyChainPOCs/types/record"
)

type AssetBoundChaincode struct {
	SupplyChainChaincode
}

func (abcc *AssetBoundChaincode) Manufacture(id string, src dto.AssetBoundRecordDto) *record.AssetBoundRecord {
	rec := abcc.SupplyChainChaincode.Manufacture(id, src.RecordDto)

	return record.NewAssetBoundRecord(rec, src.AssetId)
}

func (abcc *AssetBoundChaincode) AddAssetType(id string, assetDto *dto.AssetDto) *asset.Asset {
	return asset.NewAsset(id, assetDto)
}
