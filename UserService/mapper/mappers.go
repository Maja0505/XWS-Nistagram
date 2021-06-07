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