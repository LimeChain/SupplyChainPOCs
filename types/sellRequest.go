package types

type SellRequest struct {
	RecordId string `json:"id,omitempty"`
	Quantity uint64 `json:"quantity,omitempty"`
}
