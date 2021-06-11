package dto

type VerificationRequestDTO struct {
	Username string `bson:"username,omitempty"`
	FirstName string `bson:"first_name,omitempty"`
	LastName string `bson:"last_name,omitempty"`
	Image string `bson:"image,omitempty"`
	Category string `bson:"category,omitempty"`
}
