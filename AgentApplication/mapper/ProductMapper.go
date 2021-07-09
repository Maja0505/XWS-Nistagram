package mapper

import (
	"XWS-Nistagram/AgentApplication/dto"
	"XWS-Nistagram/AgentApplication/model"
)

func MapDTOToProduct(dto *dto.ProductDTO) *model.Product{
	var product model.Product
	product.Name = dto.Name
	product.Description =dto.Description
	product.Price,_ = dto.Price.Float64()
	product.Image = dto.Image
	product.AvailableQuantity,_ = dto.AvailableQuantity.Int64()
	return &product
}


