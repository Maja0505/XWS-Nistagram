package dto

import "userService/model"

type UserForRegistrationDTO struct {
	model.User
	ConfirmedPassword string `bson:"confirmed_password,omitempty"`
}
