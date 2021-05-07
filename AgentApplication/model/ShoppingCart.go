package model

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type ShoppingCart struct {
	ID  uuid.UUID `json:"id"`
	Orders []Order `json:"orders" gorm:"not null"`
	TotalPrice float32 `json:"totalPrice" gorm:"not null"`
	Customer User `json:"customer" gorm:"not null"`
	PaymentDetails PaymentDetails `json:"paymentDetails" gorm:"not null"`
}

func (shoppingCart *ShoppingCart) BeforeCreate(scope *gorm.DB) error {
	shoppingCart.ID = uuid.New()
	return nil
}