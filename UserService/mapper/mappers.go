package mapper

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
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
	vq.User,_ = primitive.ObjectIDFromHex(requestDTO.User)
	vq.Approved = false
	vq.Admin = "admin"
	vq.Username = requestDTO.Username
	vq.FullName = requestDTO.FullName
	vq.KnownAs = requestDTO.KnowAs
	vq.Image = requestDTO.Image
	switch requestDTO.Category {
	case "Blogger/Influencer":
		vq.Category = 0
	case "Sports":
		vq.Category = 1
	case "News/Media":
		vq.Category = 2
	case "Business/Brand/Organization":
		vq.Category = 3
	case "Government/Politics":
		vq.Category = 4
	case "Music":
		vq.Category = 5
	case "Fashion":
		vq.Category = 6
	case "Entertainment":
		vq.Category = 7
	case "Other":
		vq.Category = 8
	default:
		vq.Category = 8
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

