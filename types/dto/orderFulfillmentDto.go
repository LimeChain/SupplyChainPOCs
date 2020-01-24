package dto

type OrderFulfillmentDto struct {
	Id      string         `json:"id,omitempty"`
	Records RecordPartsDto `json:"records,omitempty"`
	Status  bool           `json:"status,omitempty"`
}
