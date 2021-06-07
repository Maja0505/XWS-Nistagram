package mapper

import (
	"userService/dto"
	"userService/model"
)

func ConvertRegisteredUserDtoToRegisteredUser(registeredUserDto *dto.RegisteredUserProfileInfoDTO) *model.RegisteredUser{

	var registeredUser model.RegisteredUser
	registeredUser.Username = registeredUserDto.Username
	registeredUser.FirstName = registeredUserDto.FirstName
	registeredUser.LastName = registeredUserDto.LastName
	registeredUser.DateOfBirth = registeredUserDto.DateOfBirth
	registeredUser.Gender = registeredUserDto.Gender
	registeredUser.Email = registeredUserDto.Email
	registeredUser.PhoneNumber = registeredUserDto.PhoneNumber
	registeredUser.Biography = registeredUserDto.Biography
	registeredUser.WebSite = registeredUserDto.WebSite

	return &registeredUser
}

func ConvertUserForRegistrationDTOToRegisteredUser(registrationDTO *dto.UserForRegistrationDTO) *model.RegisteredUser {
	var userForRegistration model.RegisteredUser
	userForRegistration.FirstName = registrationDTO.FirstName
	userForRegistration.LastName = registrationDTO.LastName
	userForRegistration.Username = registrationDTO.Username
	userForRegistration.Password = registrationDTO.Password
	userForRegistration.PhoneNumber = registrationDTO.PhoneNumber
	userForRegistration.Email = registrationDTO.Email
	userForRegistration.DateOfBirth = registrationDTO.DateOfBirth
	userForRegistration.Gender = registrationDTO.Gender
	return &userForRegistration
}

func ConvertVerificationRequestDTOToVerificationRequest(requestDTO *dto.VerificationRequestDTO) *model.VerificationRequest {
	var vq model.VerificationRequest
	vq.Username = requestDTO.Username
	vq.FirstName = requestDTO.FirstName
	vq.LastName = requestDTO.LastName
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
	return &vq
}



func ConvertUsersListTOUserFromSearchDTOList(usersList *[]model.RegisteredUser) *[]dto.UserFromSearchDTO{
	var userFromSearchDTO []dto.UserFromSearchDTO
	for _,user := range *usersList {
		userFromSearchDTO = append(userFromSearchDTO, dto.UserFromSearchDTO{Username: user.Username})
	}
	return &userFromSearchDTO
}

