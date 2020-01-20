package transparent_supply_chain_2

import "github.com/LimeChain/SupplyChainPOCs/types/dto"

type CertifiedCombineRequestDto struct {
	*dto.AssetComposeRequestDto
	QualityCertificates []string `json:"qualityCertificates,omitempty"`
}
