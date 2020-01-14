package asset

type ComposableAsset struct {
	*Asset
	AssembledFrom AssetAssemble `json:"assembledFrom,omitempty"`
}
