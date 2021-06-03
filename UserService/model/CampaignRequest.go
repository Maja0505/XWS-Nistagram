package model

import "go.mongodb.org/mongo-driver/bson/primitive"

type CampaignRequest struct {
	ID	primitive.ObjectID `bson:"_id,omitempty"`
	Agent primitive.ObjectID `bson:"agent,omitempty"`
	Admin primitive.ObjectID `bson:"admin,omitempty"`
	Approved bool `bson:"approved,omitempty"`
}
