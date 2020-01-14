package dto

type AssembableRecordDto struct {
	*RecordDto
	AssembledFrom RecordQuantityDto `json:"assembledFrom,omitempty"`
}
