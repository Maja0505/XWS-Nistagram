package mapper

import (
	"XWS-Nistagram/AgentApplication/dto"
	"XWS-Nistagram/AgentApplication/model"
	"github.com/google/uuid"
)

func MapDTOToShoppingCart(dto *dto.ShoppingCartDTO) *model.ShoppingCart{
	var shoppingCart model.ShoppingCart
	shoppingCart.UserID,_ = uuid.Parse(dto.UserID)
	shoppingCart.AddressID,_=uuid.Parse(dto.AddressID)
	shoppingCart.TotalPrice,_=dto.TotalPrice.Float64()
	return &shoppingCart
}
