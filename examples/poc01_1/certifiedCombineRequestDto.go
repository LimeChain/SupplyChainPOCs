package poc01_1

import "github.com/LimeChain/SupplyChainPOCs/types/dto"

type CertifiedCombineRequestDto struct {
	*dto.AssetComposeRequestDto
	QualityCertificates []string `json:"qualityCertificates,omitempty"`
}
