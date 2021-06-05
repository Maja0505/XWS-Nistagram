package service

import (
	"errors"
	"userService/dto"
	"userService/model"
	"userService/repository"
)

type VerificationRequestService struct {
	Repo *repository.VerificationRequestRepository
	UserService *UserService
}


func (service *VerificationRequestService) Create(verificationRequestDTO *dto.VerificationRequestDTO) error{



	if verificationRequestDTO.ConfirmedPassword != verificationRequestDTO.Password{
		return errors.New("Password and confirmed password are not same !")
	}
	user,_ := service.UserService.FindUserByUsername(verificationRequestDTO.Username)
	if user != nil{
		return errors.New("User already exist !")
	}

	vq := convertVerificationRequestDTOToVerificationRequest(verificationRequestDTO)

	err := service.Repo.Create(vq)
	if err != nil{
		return err
	}
	return nil
}

func convertVerificationRequestDTOToVerificationRequest(requestDTO *dto.VerificationRequestDTO) *model.VerificationRequest {
	var vq model.VerificationRequest
	vq.Username = requestDTO.Username
	vq.Password = requestDTO.Password
	vq.FirstName = requestDTO.FirstName
	vq.LastName = requestDTO.LastName
	vq.Email = requestDTO.Email
	vq.PhoneNumber = requestDTO.PhoneNumber
	vq.Gender = requestDTO.Gender
	switch requestDTO.Category {
	case "influencer":
		vq.Category = 0
	case "sports":
		vq.Category = 1
	case "new/media":
		vq.Category = 2
	case "business":
		vq.Category = 3
	case "brand":
		vq.Category = 4
	case "organization":
		vq.Category = 5

	}
	vq.DateOfBirth = requestDTO.DateOfBirth
	vq.Approved = false
	return &vq
}

