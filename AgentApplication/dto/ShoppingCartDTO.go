package dto

import "encoding/json"

type ShoppingCartDTO struct {
	ID  string `json:"id"`
	TotalPrice json.Number `json:"totalPrice" gorm:"not null"`
	UserID string `json:"user_id"`
	AddressID string `json:"address_id"`
}

