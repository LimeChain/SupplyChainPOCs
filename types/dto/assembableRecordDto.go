package dto

type AssembableRecordDto struct {
	*RecordDto
	AssembledFrom RecordPartsDto `json:"assembledFrom,omitempty"`
}
