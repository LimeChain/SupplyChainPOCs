package record

type AssetBoundRecord struct {
	*Record
	AssetId string `json:"assetId,omitempty"`
}

func NewAssetBoundRecord(rec *Record, assetId string) *AssetBoundRecord {
	return &AssetBoundRecord{
		Record:  rec,
		AssetId: assetId,
	}
}
