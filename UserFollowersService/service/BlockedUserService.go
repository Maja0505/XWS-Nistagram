package service

import (
	"XWS-Nistagram/UserFollowersService/model"
	"XWS-Nistagram/UserFollowersService/repository"
)

type BlockedUserService struct{
	Repository *repository.BlockedUserRepository
}

func (service *BlockedUserService) BlockUser(br *model.BlockRelationship) error{
	err := service.Repository.BlockUser(br)
	if err != nil{
		return err
	}
	return nil
}

func (service *BlockedUserService) GetAllBlockedUser(userId string) (*[]interface{},error){
	blockedUsers,err := service.Repository.GetAllBlockedUsers(userId)
	if err != nil{
		return nil,err
	}
	return blockedUsers,nil
}

func (service *BlockedUserService) CheckBlock(userId string, blockedUserId string) (interface{},error){
	block,err := service.Repository.CheckBlock(userId,blockedUserId)
	if err !=nil{
		return nil,err
	}

	return block, nil
}
