package model

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type PaymentDetails struct {
	ID  uuid.UUID `json:"id"`
	AddressID uuid.UUID `json:"address_id" `
	Address Address
	PhoneNumber string `json:"phoneNumber" gorm:"not null"`
}

type Address struct {
	ID  uuid.UUID `json:"id"`
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

func (address *Address) BeforeCreate(scope *gorm.DB) error {
	address.ID = uuid.New()
	return nil
}