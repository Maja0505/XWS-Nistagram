package model

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)


type Address struct {
	ID uuid.UUID `json:"id"`
	Address string `json:"address" gorm:"not null"`
	City string `json:"city" gorm:"not null"`
	ZipCode string `json:"zip" gorm:"not null"`
	Country string `json:"country" gorm:"not null"`

}

func (address *Address) BeforeCreate(scope *gorm.DB) error {
	address.ID = uuid.New()
	return nil
}