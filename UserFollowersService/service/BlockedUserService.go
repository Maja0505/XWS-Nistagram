package service

import (
	"XWS-Nistagram/UserFollowersService/dto"
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

func (service *BlockedUserService) GetAllBlockedUser(userId string) (*[]dto.UserByUsernameDTO,error){
	blockedUsers,err := service.Repository.GetAllBlockedUsers(userId)
	if err != nil{
		return nil,err
	}

	usernamesDTOList,err := GetUsernamesByUserIdsFromUserService(blockedUsers)

	if err != nil{
		return nil, err
	}

	return usernamesDTOList,nil

}

func (service *BlockedUserService) CheckBlock(userId string, blockedUserId string) (interface{},error){
	block,err := service.Repository.CheckBlock(userId,blockedUserId)
	if err !=nil{
		return nil,err
	}

	return block, nil
}

func (service *BlockedUserService) UnblockUser(m *model.BlockRelationship) error {
	err := service.Repository.UnblockUser(m)
	if err != nil {
		return err
	}
	return nil
}

