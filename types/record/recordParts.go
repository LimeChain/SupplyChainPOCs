package record

type RecordParts []struct {
	Id       string `json:"id,omitempty"`
	Quantity uint64 `json:"quantity,omitempty"`
}
