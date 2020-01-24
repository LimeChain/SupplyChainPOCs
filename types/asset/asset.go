package asset

import (
	"github.com/LimeChain/SupplyChainPOCs/types/dto"
	"time"
)

type Asset struct {
	Id            string          `json:"id,omitempty"`
	Description   string          `json:"description,omitempty"`
	DateCreated   time.Time       `json:"dateCreated,omitempty"`
	LastUpdated   time.Time       `json:"lastUpdated,omitempty"`
	IsActive      bool            `json:"isActive,omitempty"`
}

func NewAsset(id string, dto *dto.AssetDto) *Asset {
	return &Asset{
		Id:          id,
		Description: dto.Description,
		DateCreated: time.Now(),
		IsActive:    dto.IsActive,
	}
}