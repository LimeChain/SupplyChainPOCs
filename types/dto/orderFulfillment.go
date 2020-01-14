package dto

type OrderFulfillmentDto struct {
	Id      string            `json:"id,omitempty"`
	Records RecordQuantityDto `json:"records,omitempty"`
	Status  bool              `json:"status,omitempty"`
}
