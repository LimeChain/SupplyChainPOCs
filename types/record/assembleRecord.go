package record

type AssembleRecord []struct {
	Id       string `json:"id,omitempty"`
	Quantity uint64 `json:"quantity,omitempty"`
}
