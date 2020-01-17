package dto

type AssembableRecordDto struct {
	*BaseRecordDto
	AssembledFrom RecordPartsDto `json:"assembledFrom,omitempty"`
}
