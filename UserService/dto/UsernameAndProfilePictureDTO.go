package dto

type UsernameAndProfilePictureDTO struct {
	ProfilePicture string `bson:"profile_picture,omitempty"`
	Username string `bson:"username,omitempty"`
}
