package record

type AssetBoundRecord struct {
	*BaseRecord
	AssetId string `json:"assetId,omitempty"`
}

func NewAssetBoundRecord(rec *BaseRecord, assetId string) *AssetBoundRecord {
	return &AssetBoundRecord{
		BaseRecord: rec,
		AssetId:    assetId,
	}
}
