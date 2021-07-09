package DTO

type UserByUsernameDTO struct {

	IdString string `bson:"id_string,omitempty"`
	Username string  `bson:"username,omitempty"`

}
