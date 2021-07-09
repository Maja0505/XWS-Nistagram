package model

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"time"
)

type ShoppingCart struct {
	ID  uuid.UUID `json:"id"`
	Orders []Order `json:"orders" gorm:"foreignKey:ShoppingCartID"`
	TotalPrice float64 `json:"totalPrice" gorm:"not null"`
	UserID uuid.UUID `json:"user_id"`
	User User
	AddressID uuid.UUID `json:"paymentDetails_id"`
	Address Address
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`

 }

func (shoppingCart *ShoppingCart) BeforeCreate(scope *gorm.DB) error {
	shoppingCart.ID = uuid.New()
	return nil
}