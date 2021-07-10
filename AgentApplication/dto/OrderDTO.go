package dto

import "encoding/json"

type OrderDTO struct {
	ID  string	`json:"id"`
	ProductID string	`json:"productId"`
	Amount json.Number `json:"amount"`
	ShoppingCartID string `json:"shoppingCartId"`
}

