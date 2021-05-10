package model

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type ShoppingCart struct {
	ID  uuid.UUID `json:"id"`
	Orders []Order `json:"orders" gorm:"polymorphic:Product;"`
	TotalPrice float32 `json:"totalPrice" gorm:"not null"`
	UserID uuid.UUID `json:"user_id"`
	User User
	PaymentDetailsID uuid.UUID `json:"paymentDetails_id"`
	PaymentDetails PaymentDetails

}

func (shoppingCart *ShoppingCart) BeforeCreate(scope *gorm.DB) error {
	shoppingCart.ID = uuid.New()
	return nil
}