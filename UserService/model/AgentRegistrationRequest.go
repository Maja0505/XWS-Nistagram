package model

import "go.mongodb.org/mongo-driver/bson/primitive"

type AgentRegistrationRequest struct {
	ID	primitive.ObjectID `bson:"_id,omitempty"`
	User primitive.ObjectID `bson:"user,omitempty"`
	Admin primitive.ObjectID `bson:"admin,omitempty"`
	WebSite string `bson:"web_site,omitempty"`
	Approved bool `bson:"approved,omitempty"`
}
