package dto

type SellDto struct {
	RecordId string `json:"id,omitempty"`
	Quantity uint64 `json:"quantity,omitempty"`
}
