package model

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"time"
)

type Product struct {
	ID  uuid.UUID `json:"id"`
	Name string `json:"name" gorm:"not null"`
	Description string `json:"description" gorm:"not null"`
	Image string `json:"image"`
	AvailableQuantity int64 `json:"availableQuantity" gorm:"not null"`
	Price float64 `json:"price" gorm:"not null"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

func (product *Product) BeforeCreate(scope *gorm.DB) error {
	product.ID = uuid.New()
	return nil
}
