package model

type FollowRelationship struct{
	User       string `json:"user"`
	FollowedUser             string `json:"followedUser"`
}
