package poc01_1

import (
	"github.com/LimeChain/SupplyChainPOCs/types/record"
)

type CertifiedRecord struct {
	*record.BaseRecord
	*record.AssembableRecord
	*record.AssetBoundRecord
	QualityCertificates []string `json:"qualityCertificates,omitempty"`
}

func NewCertifiedRecord(recordStruct *record.BaseRecord, assetId string, assembledFrom record.RecordParts, qualityCertificates []string) *CertifiedRecord {
	result := &CertifiedRecord{
		BaseRecord: recordStruct,
		AssembableRecord: &record.AssembableRecord{
			BaseRecord:    nil,
			AssembledFrom: assembledFrom,
		},
		AssetBoundRecord: &record.AssetBoundRecord{
			BaseRecord: nil,
			AssetId:    assetId,
		},
		QualityCertificates: qualityCertificates,
	}

	return result
}
