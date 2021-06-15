package model

import "go.mongodb.org/mongo-driver/bson/primitive"

type Category int

const (
	Blogger_Influencer Category = iota
	Sports
	News_Media
	Business_Brand_Organization
	Government_Politics
	Music
	Fashion
	Entertainment
	Other
)

type VerificationRequest struct {
	ID	primitive.ObjectID `bson:"_id,omitempty"`
	User primitive.ObjectID `bson:"user,omitempty"`
	Username string `bson:"username,omitempty"`
	FullName string `bson:"full_name,omitempty"`
	KnownAs string `bson:"known_as,omitempty"`
	Admin string `bson:"admin,omitempty"`
	Category Category `bson:"category,omitempty"`
	Approved bool `bson:"approved,omitempty"`
	Image string `bson:"image,omitempty"`
}
