package dto

type RecordPartsDto []struct {
	Id       string `json:"id,omitempty"`
	Quantity uint64 `json:"quantity,omitempty"`
}
