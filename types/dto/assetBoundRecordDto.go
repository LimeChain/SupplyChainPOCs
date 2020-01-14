package dto

type AssetBoundRecordDto struct {
	AssetId string `json:"assetId,omitempty"`
	*RecordDto
}
