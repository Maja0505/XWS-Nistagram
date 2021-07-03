package service

import (
	"XWS-Nistagram/UserFollowersService/dto"
	"XWS-Nistagram/UserFollowersService/mapper"
	"XWS-Nistagram/UserFollowersService/repository"
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
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

func (service *UserFollowersService) GetAllFollowedUsers(userId string) ( *[]dto.UserByUsernameDTO,error) {

	users,err := service.Repository.GetAllFollowedUsersByUser(userId)

	if err != nil{
		return nil, err
	}

	usernamesDTOList,err := GetUsernamesByUserIdsFromUserService(users)

	if err != nil{
		return nil, err
	}

	return usernamesDTOList,nil

}

func (service *UserFollowersService) GetAllFollowersByUser(userId string) (*[]dto.UserByUsernameDTO, error) {

	users,err := service.Repository.GetAllFollowersByUser(userId)

	if err != nil{
		return nil, err
	}

	usernamesDTOList,err := GetUsernamesByUserIdsFromUserService(users)

	if err != nil{
		return nil, err
	}

	return usernamesDTOList,nil
}

func (service *UserFollowersService) GetAllNotMutedFollowedUsersByUser(userId string)(*[]interface{}, error){

	users,err := service.Repository.GetAllNotMutedFollowedUsersByUser(userId)

	if err !=nil {
		return nil,err
	}

	return users,nil

}

func (service *UserFollowersService) GetAllFollowsWhomUserIsCloseFriend(userId string)(*[]interface{}, error){

	notMutedFollows,err := service.Repository.GetAllNotMutedFollowedUsersByUser(userId)
	follows,err := service.Repository.GetAllFollowsWhomUserIsCloseFriend(userId)

	var users []interface{}
	

	if err !=nil {
		return nil,err
	}

	for _, notMutedFollowUserId := range *notMutedFollows {
		for _, followUserId := range *follows {
			if notMutedFollowUserId == followUserId{
				users = append(users, notMutedFollowUserId)
				break
			}
		}
	}
	return &users,nil

}

func (service *UserFollowersService) GetAllFollowsWhomUserIsNotCloseFriend(userId string)(*[]interface{}, error){

	notMutedFollows,err := service.Repository.GetAllNotMutedFollowedUsersByUser(userId)
	follows,err := service.Repository.GetAllFollowsWhomUserIsCloseFriend(userId)

	var users []interface{}

	if err !=nil {
		return nil,err
	}

	for _, notMutedFollowUserId := range *notMutedFollows {
		add := true
		for _, followUserId := range *follows {
			if notMutedFollowUserId == followUserId{
				add = false
				break
			}
		}
		if add{
			users = append(users,notMutedFollowUserId)
		}
	}
	return &users,nil

}

func (service *UserFollowersService) GetAllFollowsWithoutFollowsWhomUserIsCloseFriend(userId string)(*[]interface{}, error){

	notMutedFollows,err := service.Repository.GetAllNotMutedFollowedUsersByUser(userId)
	follows,err := service.Repository.GetAllFollowsWhomUserIsCloseFriend(userId)

	var users []interface{}


	if err !=nil {
		return nil,err
	}

	for _, notMutedFollowUserId := range *notMutedFollows {
		for _, followUserId := range *follows {
			if notMutedFollowUserId == followUserId{
				users = append(users, notMutedFollowUserId)
				break
			}
		}
	}
	return &users,nil

}

func (service *UserFollowersService) GetAllFollowRequests(userId string) (*[]dto.UserByUsernameDTO, error) {

	users,err := service.Repository.GetAllFollowRequests(userId)

	if err != nil{
		return nil, err
	}

	usernamesDTOList,err := GetUsernamesByUserIdsFromUserService(users)

	if err != nil{
		return nil, err
	}

	return usernamesDTOList,nil
}

func (service *UserFollowersService) GetAllCloseFriends(userId string) (*[]dto.UserByUsernameDTO, error) {

	closeFriends,err := service.Repository.GetAllCloseFriends(userId)

	if err != nil{
		return nil, err
	}

	usernamesDTOList,err := GetUsernamesByUserIdsFromUserService(closeFriends)

	if err != nil{
		return nil, err
	}

	return usernamesDTOList,nil
}

func (service *UserFollowersService) GetAllMuteFriends(userId string) (*[]dto.UserByUsernameDTO, error) {

	muteFriends,err := service.Repository.GetAllMuteFriends(userId)

	if err != nil{
		return nil, err
	}

	usernamesDTOList,err := GetUsernamesByUserIdsFromUserService(muteFriends)

	if err != nil{
		return nil, err
	}

	return usernamesDTOList,nil
}

func (service *UserFollowersService) CheckFollowing(userId string, followedUserId string) (*interface{}, error) {
	following,err := service.Repository.CheckFollowing(userId , followedUserId)
	if err != nil{
		return nil, err
	}

	return following,err
}

func (service *UserFollowersService) CheckRequested(userId string, followedUserId string) (*interface{}, error) {
	requested,err := service.Repository.CheckRequested(userId , followedUserId)
	if err != nil{
		return nil, err
	}

	return requested,err
}

func (service *UserFollowersService) CheckMuted(userId string, mutedUserId string) (*interface{}, error) {
	muted,err := service.Repository.CheckMuted(userId , mutedUserId)
	if err != nil{
		return nil, err
	}

	return muted,err
}

func (service *UserFollowersService) CheckClosed(userId string, closedUserId string) (*interface{}, error) {
	closed,err := service.Repository.CheckClosed(userId , closedUserId)
	if err != nil{
		return nil, err
	}

	return closed,err
}


func GetUsernamesByUserIdsFromUserService (users *[]interface{}) (*[]dto.UserByUsernameDTO,error){

	reqUrl := fmt.Sprintf("http://" +os.Getenv("USER_SERVICE_DOMAIN") + ":" + os.Getenv("USER_SERVICE_PORT")+ "/convert-user-ids")

	type UserIdsDTO struct {
		UserIds []interface{}
	}

	userIdsDto := UserIdsDTO{}
	userIdsDto.UserIds = *users
	jsonUserids,_ := json.Marshal(userIdsDto)
	resp, err := http.Post(reqUrl,"appliation/json",bytes.NewBuffer(jsonUserids))

	if err != nil || resp.StatusCode == 404 {
		return nil,err
	}
	body, err := ioutil.ReadAll(resp.Body)
	var data []dto.UserByUsernameDTO
	err = json.Unmarshal(body, &data)
	if err != nil{
		return nil, err
	}
	return &data, nil
}