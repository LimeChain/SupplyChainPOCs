package record

import (
	"github.com/LimeChain/SupplyChainPOCs/types/dto"
	"time"
)

type Record struct {
	Id              string    `json:"id,omitempty"`
	BatchId         string    `json:"batchId,omitempty"`
	Owner           string    `json:"owner,omitempty"`
	Quantity        uint64    `json:"quantity,omitempty"`
	DateCreated     time.Time `json:"dateCreated,omitempty"`
	LastUpdated     time.Time `json:"lastUpdated,omitempty"`
	CreationOrderId string    `json:"creationOrderId,omitempty"`
}

func NewRecord(id string, dto *dto.RecordDto) *Record {
	return &Record{
		Id:          id,
		BatchId:     dto.BatchId,
		Owner:       dto.Owner,
		Quantity:    dto.Quantity,
		DateCreated: time.Now(),
	}
}

func (record *Record) DecreaseQuantity(quantity uint64) {
	record.Quantity -= quantity
	record.update()
}

func (record *Record) SetQuantity(quantity uint64) {
	record.Quantity = quantity
	record.update()
}

func (record *Record) GetNewQuantity(subtractAmount uint64) uint64 {
	return record.Quantity - subtractAmount
}

func (record *Record) update() {
	record.LastUpdated = time.Now()
}
