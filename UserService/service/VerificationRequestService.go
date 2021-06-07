package service

import (
	"userService/dto"
	"userService/mapper"
	"userService/repository"
)

type VerificationRequestService struct {
	Repo *repository.VerificationRequestRepository
	UserService *UserService
}


func (service *VerificationRequestService) Create(verificationRequestDTO *dto.VerificationRequestDTO) error{
	vq := mapper.ConvertVerificationRequestDTOToVerificationRequest(verificationRequestDTO)
	err := service.Repo.Create(vq)
	if err != nil{
		return err
	}
	return nil
}



