 package model

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"time"
)

type Order struct {
	ID  uuid.UUID `json:"id"`
	ProductID uuid.UUID `json:"product_id"`
	Product Product
	Amount int64 `json:"amount" gorm:"not null"`
	ShoppingCartID uuid.UUID
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

func (order *Order) BeforeCreate(scope *gorm.DB) error {
	order.ID = uuid.New()
	return nil
}
