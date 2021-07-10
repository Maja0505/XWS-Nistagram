package model

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"time"
)

type Purchase struct {
	ID  uuid.UUID `json:"id"`
	Orders []Order `json:"orders" gorm:"-"`
	TotalPrice float64 `json:"totalPrice" gorm:"not null"`
	UserID uuid.UUID `json:"user_id"`
	User User
	AddressID uuid.UUID `json:"paymentDetails_id"`
	Address Address
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

func (purchase *Purchase) BeforeCreate(scope *gorm.DB) error {
	purchase.ID = uuid.New()
	return nil
}