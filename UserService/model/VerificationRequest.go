package model

import "go.mongodb.org/mongo-driver/bson/primitive"

type Category int

const (
	Influencer Category = iota
	Sports
	New_Media
	Business
	Brand
	Organization
)

type VerificationRequest struct {
	ID	primitive.ObjectID `bson:"_id,omitempty"`
	User primitive.ObjectID `bson:"user,omitempty"`
	Admin primitive.ObjectID `bson:"admin,omitempty"`
	Category Category `bson:"category,omitempty"`
	Approved bool `bson:"approved,omitempty"`
	Image string `bson:"image,omitempty"`
}
