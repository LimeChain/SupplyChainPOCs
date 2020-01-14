package record

type AssembableRecord struct {
	*Record
	AssembledFrom AssembleRecord `json:"assembledFrom,omitempty"`
}

func NewAssembableRecord(rec *Record, recordAssemble AssembleRecord) *AssembableRecord {
	return &AssembableRecord{
		Record:        rec,
		AssembledFrom: recordAssemble,
	}
}
