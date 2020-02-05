package dto

type QueryPrivateDto struct {
	Collection string `json:"collection,omitempty"`
	Key        string `json:"key,omitempty"`
}
