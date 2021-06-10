package service

import (
	"XWS-Nistagram/UserFollowersService/model"
	"XWS-Nistagram/UserFollowersService/repository"
)

type UserFollowersService struct{
	Repository *repository.UserFollowersRepository
}

func (service *UserFollowersService) FollowUser(fr model.FollowRelationship, ce chan error){
	service.Repository.FollowUser(fr,ce)
}
