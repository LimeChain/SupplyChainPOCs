package record

type AssembableRecord struct {
	*BaseRecord
	AssembledFrom RecordParts `json:"assembledFrom,omitempty"`
}

func NewAssembableRecord(rec *BaseRecord, recordParts RecordParts) *AssembableRecord {
	return &AssembableRecord{
		BaseRecord:    rec,
		AssembledFrom: recordParts,
	}
}
