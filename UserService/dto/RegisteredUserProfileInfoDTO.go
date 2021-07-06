package dto

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"userService/model"
)

type RegisteredUserProfileInfoDTO struct {
	Username string
	FirstName string
	LastName string
	DateOfBirth *primitive.DateTime
	Email string
	PhoneNumber string
	Gender *model.Gender
	Biography string
	WebSite string
	ProfilePicture string
}