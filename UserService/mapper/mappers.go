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

func ConvertVerificationRequestsToVerificationRequestDTO(vr *model.VerificationRequest) *dto.VerificationRequestDTO{
	return &dto.VerificationRequestDTO{Username: vr.Username,
		User: vr.User.Hex(),
		FullName: vr.FullName,
		KnowAs: vr.KnownAs,
		Category: pickCategory(vr.Category),
		Image: vr.Image,
	}
}

func ConvertVerificationRequestsListToVerificationRequestDTOList(verificationRequests *[]model.VerificationRequest) *[]dto.VerificationRequestDTO{
	var verificationRequestDTOList []dto.VerificationRequestDTO
	for _,vr := range *verificationRequests{
		verificationRequestDTOList = append(verificationRequestDTOList, *ConvertVerificationRequestsToVerificationRequestDTO(&vr))
	}
	return &verificationRequestDTOList
}

func pickCategory(category model.Category) string{
	switch int(category) {
	case 0:
		return "Blogger/Influencer"
	case 1:
		return "Sports"
	case 2:
		return "News/Media"
	case 3:
		return "Business/Brand/Organization"
	case 4:
		return "Government/Politics"
	case 5:
		return "Music"
	case 6:
		return "Fashion"
	case 7:
		return "Entertainment"
	case 8:
		return "Other"

	}
	return "Other"
}

func ConvertUsersListTOUserFromSearchDTOList(usersList *[]model.RegisteredUser) *[]dto.UserFromSearchDTO{
	var userFromSearchDTO []dto.UserFromSearchDTO
	for _,user := range *usersList {
		userFromSearchDTO = append(userFromSearchDTO, dto.UserFromSearchDTO{Username: user.Username})
	}
	return &userFromSearchDTO
}

