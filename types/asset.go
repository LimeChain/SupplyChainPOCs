package types

import (
	"time"
)

type Asset struct {
	Id            string          `json:"id,omitempty"`
	Description   string          `json:"description,omitempty"`
	DateCreated   time.Time       `json:"dateCreated,omitempty"`
	LastUpdated   time.Time       `json:"lastUpdated,omitempty"`
	IsActive      bool            `json:"isActive,omitempty"`
	AssembledFrom []AssetAssembly `json:"assembledFrom,omitempty"`
}
