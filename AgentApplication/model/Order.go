package model

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Order struct {
	ID  uuid.UUID `json:"id"`
	ProductID uuid.UUID `json:"product_id"`
	Product Product
	ProductType string
	Amount int32 `json:"amount" gorm:"not null"`
}

func (order *Order) BeforeCreate(scope *gorm.DB) error {
	order.ID = uuid.New()
	return nil
}
