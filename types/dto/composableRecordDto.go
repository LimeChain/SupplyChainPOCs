package dto

type ComposableRecordDto struct {
	*BaseRecordDto
	ComposedFrom RecordPartsDto `json:"composedFrom,omitempty"`
}
