package mapper

import (
	"XWS-Nistagram/AgentApplication/dto"
	"XWS-Nistagram/AgentApplication/model"
)

func MapDTOToUser(dto *dto.UserDTO) *model.User {
	var user model.User
	user.Email = dto.Email
	user.FirstName = dto.FirstName
	user.LastName  = dto.LastName
	user.Password = dto.Password
	user.Role = dto.Role
	return &user
}