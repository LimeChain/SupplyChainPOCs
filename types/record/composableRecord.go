package record

type ComposableRecord struct {
	*BaseRecord
	ComposedFrom RecordParts `json:"composedFrom,omitempty"`
}

func NewComposableRecord(rec *BaseRecord, recordParts RecordParts) *ComposableRecord {
	return &ComposableRecord{
		BaseRecord:   rec,
		ComposedFrom: recordParts,
	}
}
