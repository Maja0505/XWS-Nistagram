package service

import (
	"XWS-Nistagram/UserFollowersService/dto"
	"XWS-Nistagram/UserFollowersService/mapper"
	"XWS-Nistagram/UserFollowersService/repository"
)

type UserFollowersService struct{
	Repository *repository.UserFollowersRepository
}

func (service *UserFollowersService) FollowUser(dto *dto.FollowRelationshipDTO) error {

	fr := mapper.ConvertFollowRelationshipDTOTOFollowRelationship(dto)

	var err error
	if dto.Private == false {
		err = service.Repository.FollowUser(fr)
	}else{
		err = service.Repository.SendFollowRequest(fr)
	}

	if err != nil {
		return err
	}
	return nil
}

func (service *UserFollowersService) UnfollowUser(dto *dto.UnfollowRelationshipDTO) error {

	fr := mapper.ConvertUnFollowRelationshipDTOTOFollowRelationship(dto)
	err := service.Repository.UnfollowUser(fr)
	if err != nil {
		return err
	}
	return nil

}

func (service *UserFollowersService) AcceptFollowRequest(dto *dto.FollowRequestDTO) error {

	err := service.Repository.AcceptFollowRequest(dto.User,dto.UserWitchSendRequest)
	if err != nil {
		return err
	}
	return nil
}

func (service *UserFollowersService) CancelFollowRequest(dto *dto.FollowRequestDTO) error {

	err := service.Repository.CancelFollowRequest(dto.User,dto.UserWitchSendRequest)
	if err != nil {
		return err
	}
	return nil
}

func (service *UserFollowersService) SetCloseFriend(dto *dto.CloseFriendDTO) error{
	err := service.Repository.SetFriendForClose(dto.User,dto.Friend,dto.Close)
	if err != nil{
		return err
	}
	return nil
}

func (service *UserFollowersService) SetMuteFriend(dto *dto.MuteFriendDTO) error{
	err := service.Repository.SetFriendForMute(dto.User,dto.Friend,dto.Mute)
	if err != nil{
		return err
	}
	return nil
}

func (service *UserFollowersService) GetAllFollowedUsers(userId string) ( *[]interface{},error) {

	users,err := service.Repository.GetAllFollowedUsersByUser(userId)

	if err != nil{
		return nil, err
	}

	return users,nil
}

func (service *UserFollowersService) GetAllFollowersByUser(userId string) (*[]interface{}, error) {

	users,err := service.Repository.GetAllFollowersByUser(userId)

	if err != nil{
		return nil, err
	}

	return users,nil
}

func (service *UserFollowersService) GetAllFollowRequests(userId string) (*[]interface{}, error) {

	users,err := service.Repository.GetAllFollowRequests(userId)

	if err != nil{
		return nil, err
	}

	return users,nil
}

func (service *UserFollowersService) GetAllCloseFriends(userId string) (*[]interface{}, error) {

	closeFriends,err := service.Repository.GetAllCloseFriends(userId)

	if err != nil{
		return nil, err
	}

	return closeFriends,nil
}

func (service *UserFollowersService) GetAllMuteFriends(userId string) (*[]interface{}, error) {

	muteFriends,err := service.Repository.GetAllMuteFriends(userId)

	if err != nil{
		return nil, err
	}

	return muteFriends,nil
}


func (service *UserFollowersService) CheckFollowing(userId string, followedUserId string) (*interface{}, error) {
	following,err := service.Repository.CheckFollowing(userId , followedUserId)
	if err != nil{
		return nil, err
	}

	return following,err
}
