package mapper

import (
	"XWS-Nistagram/UserFollowersService/dto"
	"XWS-Nistagram/UserFollowersService/model"
)

func ConvertFollowRelationshipDTOTOFollowRelationship(dto *dto.FollowRelationshipDTO) *model.FollowRelationship{
	var fr model.FollowRelationship
	fr.User = dto.User
	fr.FollowedUser = dto.FollowedUser
	fr.CloseFriend = false
	fr.Muted = false
	return &fr
}

func ConvertUnFollowRelationshipDTOTOFollowRelationship(dto *dto.UnfollowRelationshipDTO) *model.FollowRelationship {
	var fr model.FollowRelationship
	fr.User = dto.User
	fr.FollowedUser = dto.UnfollowedUser
	return &fr
}

