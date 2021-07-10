package dto

import "userService/model"

type UserForRegistrationDTO struct {
	model.User
	IsAgent bool `bson:"is_agent,omitempty"`
	ConfirmedPassword string `bson:"confirmed_password,omitempty"`
	WebSite string `bson:"web_site,omitempty"`
}
