package dto

type AssetDto struct {
	Description string `json:"description,omitempty"`
	IsActive    bool   `json:"isActive,omitempty"`
}
