package model

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"image"
)

type Product struct {
	ID  uuid.UUID `json:"id"`
	Name string `json:"name" gorm:"not null"`
	Description string `json:"description" gorm:"not null"`
	Image image.Image `json:"image"`
	AvailableQuantity int64 `json:"availableQuantity" gorm:"not null"`
	Price float32 `json:"price" gorm:"not null"`
}

func (product *Product) BeforeCreate(scope *gorm.DB) error {
	product.ID = uuid.New()
	return nil
}
