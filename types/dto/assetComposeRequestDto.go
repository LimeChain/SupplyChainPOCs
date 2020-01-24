package dto

type AssetComposeRequestDto struct {
	*ComposeRequestDto
	AssetId string `json:"assetId,omitempty"`
}
