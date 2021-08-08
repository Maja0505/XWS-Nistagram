package model

import "go.mongodb.org/mongo-driver/bson/primitive"

type Gender int

const (
	Female Gender = iota
	Male
)

type User struct {
	ID	primitive.ObjectID `bson:"_id,omitempty"`
	IdString string `bson:"id_string,omitempty"`
	FirstName string  `bson:"first_name,omitempty"`
	LastName string  `bson:"last_name,omitempty"`
	Username string  `bson:"username,omitempty"`
	Password string  `bson:"password,omitempty"`
	Email string  `bson:"email,omitempty"`
	PhoneNumber string  `bson:"phone_number,omitempty"`
	DateOfBirth *primitive.DateTime `bson:"date_of_birth,omitempty"`
	Gender *Gender `bson:"gender,omitempty"`
}

