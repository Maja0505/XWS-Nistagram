package model

type FollowRelationship struct{
	User       		string
	FollowedUser	string
	CloseFriend		bool
	Muted			bool
}
