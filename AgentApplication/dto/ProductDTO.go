package dto

import "encoding/json"

type ProductDTO struct {
	ID  string `json:"id"`
	Name string `json:"name"`
	Description string `json:"description"`
	Image string `json:"image"`
	AvailableQuantity json.Number `json:"availableQuantity"`
	Price json.Number  `json:"price"`
}
