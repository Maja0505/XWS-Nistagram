package dto

type VerificationRequestDTO struct {
	User string `bson:"user,omitempty"`
	Username string `bson:"username,omitempty"`
	FullName string `bson:"full_name,omitempty"`
	KnowAs string `bson:"know_as,omitempty"`
	Image string `bson:"image,omitempty"`
	Category string `bson:"category,omitempty"`
}
