package transparent_supply_chain_2

import (
	"github.com/LimeChain/SupplyChainPOCs/types/record"
)

type CertifiedRecord struct {
	*record.BaseRecord
	*record.ComposableRecord
	*record.AssetBoundRecord
	QualityCertificates []string `json:"qualityCertificates,omitempty"`
}

func NewCertifiedRecord(recordStruct *record.BaseRecord, assetId string, composedFrom record.RecordParts, qualityCertificates []string) *CertifiedRecord {
	return &CertifiedRecord{
		BaseRecord: recordStruct,
		ComposableRecord: &record.ComposableRecord{
			BaseRecord:   nil,
			ComposedFrom: composedFrom,
		},
		AssetBoundRecord: &record.AssetBoundRecord{
			BaseRecord: nil,
			AssetId:    assetId,
		},
		QualityCertificates: qualityCertificates,
	}
}
