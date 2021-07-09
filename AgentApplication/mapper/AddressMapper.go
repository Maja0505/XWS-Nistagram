package mapper

import (
	"XWS-Nistagram/AgentApplication/dto"
	"XWS-Nistagram/AgentApplication/model"
)

func MapDTOToAddress(dto *dto.AddressDTO) *model.Address{
	var address model.Address
	address.Address = dto.Address
	address.Country = dto.Country
	address.ZipCode = dto.ZIPCode
	address.City = dto.City
	return &address
}

