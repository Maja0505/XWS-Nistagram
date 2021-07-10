package mapper

import (
	"XWS-Nistagram/AgentApplication/dto"
	"XWS-Nistagram/AgentApplication/model"
	"github.com/google/uuid"
)

func MapDTOToPurchase(dto *dto.PurchaseDTO) *model.Purchase{
	var purchase model.Purchase
	purchase.Orders= *MapToListOrders(dto.Orders)
	purchase.AddressID,_ = uuid.Parse(dto.AddressID)
	purchase.UserID,_=uuid.Parse(dto.UserID)
	purchase.TotalPrice,_=dto.TotalPrice.Float64()
	return &purchase
}

