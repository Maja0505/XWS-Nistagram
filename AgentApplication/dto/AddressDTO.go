package dto

import "github.com/google/uuid"

type AddressDTO struct {
	ID uuid.UUID `json:"id"`
	Address string `json:"address" gorm:"not null"`
	City string `json:"city" gorm:"not null"`
	ZIPCode string `json:"zip" gorm:"not null"`
	Country string `json:"country" gorm:"not null"`

}
