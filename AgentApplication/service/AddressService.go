package service

import (
	"XWS-Nistagram/AgentApplication/dto"
	"XWS-Nistagram/AgentApplication/mapper"
	"XWS-Nistagram/AgentApplication/model"
	"XWS-Nistagram/AgentApplication/repository"
)

type  AddressService struct{
	Repository *repository.AddressRepository
}

func (service *AddressService) CreateAddress(address dto.AddressDTO) (error,model.Address){
	return  service.Repository.CreateAddress(mapper.MapDTOToAddress(&address))
}
