package dto

import (
	"encoding/json"
)

type PurchaseDTO struct {
	ID  string `json:"id"`
	Orders []OrderDTO `json:"orders"`
	TotalPrice json.Number `json:"totalPrice" `
	UserID string `json:"user_id"`
	AddressID string `json:"address_id"`
}
