package poc01_1

import (
	"github.com/LimeChain/SupplyChainPOCs/types/dto"
)

type CertifiedRecordDto struct {
	*dto.BaseRecordDto
	*dto.ComposableRecordDto
	*dto.AssetBoundRecordDto
	QualityCertificates []string `json:"qualityCertificates,omitempty"`
}
