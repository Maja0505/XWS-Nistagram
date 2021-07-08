package model

type AgentRegistrationRequest struct {
	User `bson:",inline"`
	WebSite string `bson:"web_site,omitempty"`
	Approved bool `bson:"approved"`
}
