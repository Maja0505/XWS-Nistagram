package mapper

import (
	"XWS-Nistagram/AgentApplication/dto"
	"XWS-Nistagram/AgentApplication/model"
	"github.com/google/uuid"
)

func MapDTOToOrder(dto *dto.OrderDTO) *model.Order{
	var order model.Order
	order.ShoppingCartID,_ = uuid.Parse(dto.ShoppingCartID)
	order.Amount,_ = dto.Amount.Int64()
	order.ProductID,_=uuid.Parse(dto.ProductID)
	return &order
}

func MapToListOrders(dtos []dto.OrderDTO) *[]model.Order{
	var orders []model.Order
	for i := 1; i < len(dtos); i++ {
		orders=append(orders, *MapDTOToOrder(&dtos[i]))
	}
	return &orders
}



