package record

type AssembableRecord struct {
	*Record
	AssembledFrom RecordParts `json:"assembledFrom,omitempty"`
}

func NewAssembableRecord(rec *Record, recordParts RecordParts) *AssembableRecord {
	return &AssembableRecord{
		Record:        rec,
		AssembledFrom: recordParts,
	}
}
