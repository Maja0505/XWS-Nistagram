package model

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type PaymentDetails struct {
	ID  uuid.UUID `json:"id"`
	Address Address `json:"address" gorm:"not null"`
	PhoneNumber string `json:"phoneNumber" gorm:"not null"`
}

type Address struct {
	StreetName string `json:"streetName" gorm:"not null"`
	StreetNumber string `json:"streetNumber" gorm:"not null"`
	City string `json:"city" gorm:"not null"`
	Longitude string `json:"longitude" gorm:"not null"`
	Latitude string `json:"latitude" gorm:"not null"`

}

func (paymentDetails *PaymentDetails) BeforeCreate(scope *gorm.DB) error {
	paymentDetails.ID = uuid.New()
	return nil
}
