package dto

import (
	"userService/model"
)

type VerificationRequestDTO struct {
	model.User
	ConfirmedPassword string `bson:"confirmed_password,omitempty"`
	Category string `bson:"category,omitempty"`
}
