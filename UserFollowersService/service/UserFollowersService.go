package service

import (
	"XWS-Nistagram/UserFollowersService/dto"
	"XWS-Nistagram/UserFollowersService/mapper"
	"XWS-Nistagram/UserFollowersService/repository"
	"XWS-Nistagram/UserFollowersService/saga"
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"sort"
	"github.com/go-redis/redis"
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

func (service *UserFollowersService) FollowSuggestions(userId string) (*[]dto.UserByUsernameDTO, error) {

	var followSuggestions []interface{}
	mapOfUsers := make(map[string]float64)

	allFollowedUsers, err := service.Repository.GetAllFollowedUsersByUser(userId)
	if err != nil{
		return nil,err
	}
	var listOfFollowedUserFollowed []interface{}
	for _,followerUser := range *allFollowedUsers {
		strFollowerUser := fmt.Sprintf("%v",followerUser)
		if userId != strFollowerUser {
			list,err := service.Repository.GetAllFollowedUsersByUser(strFollowerUser)
			if err != nil{
				return nil,err
			}
			s := strFollowerUser + "->" + fmt.Sprintf("%v",list)
			fmt.Println(s)
			listOfFollowedUserFollowed = append(listOfFollowedUserFollowed,*list...)
		}else{
			continue
		}
	}

	fmt.Println()

	var listOfFollowedUserFollowedUserFollowedUser []interface{}
	for _,f := range listOfFollowedUserFollowed {

		strF := fmt.Sprintf("%v",f)
		following, err := service.Repository.CheckFollowing(userId,strF)
		if err != nil {
			return nil, err
		}
		strFollowing := fmt.Sprintf("%v",*following)
		if strF != userId && strFollowing == "false" {
			_,existInMap := mapOfUsers[strF]
			if existInMap{
				mapOfUsers[strF] = mapOfUsers[strF] + 2
			}else{
				mapOfUsers[strF] = 1
			}

			list,err := service.Repository.GetAllFollowedUsersByUser(strF)
			if err != nil{
				return nil,err
			}
			s := strF + "->" + fmt.Sprintf("%v",list)
			fmt.Println(s)
			listOfFollowedUserFollowedUserFollowedUser = append(listOfFollowedUserFollowedUserFollowedUser,*list...)
		}else{
			continue
		}

	}

	for _,f := range listOfFollowedUserFollowedUserFollowedUser{
		strF := fmt.Sprintf("%v",f)
		following, err := service.Repository.CheckFollowing(userId,strF)
		if err != nil {
			return nil, err
		}
		strFollowing := fmt.Sprintf("%v",*following)
		if strF != userId && strFollowing == "false" {
		_,existInMap := mapOfUsers[strF]
		if existInMap{
			mapOfUsers[strF] = mapOfUsers[strF] + 0.5
		}else{
				mapOfUsers[strF] = 0.25
			}
		}else{
			continue
		}
	}
	sorted := make(PairList,len(mapOfUsers))
	var i int
	for key, value := range mapOfUsers {
		sorted[i] = Pair{key, value}
		i++
	}
	sort.Sort(sort.Reverse(sorted))
	var number = 0
	for _, v := range sorted {
		if number == 15 {
			break
		}
		followSuggestions = append(followSuggestions,v.Key)
		number++
	}

	if len(followSuggestions) == 0 {
		allUsers, err := service.Repository.GetAllUsers(userId)
		if err != nil{
			return nil, err
		}
		usernamesDTOList,err := GetUsernamesByUserIdsFromUserService(allUsers)

		if err != nil{
			return nil, err
		}


		return usernamesDTOList,nil
	}

	usernamesDTOList,err := GetUsernamesByUserIdsFromUserService(&followSuggestions)

	if err != nil{
		return nil, err
	}


	return usernamesDTOList,nil

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

type Pair struct {
	Key   string
	Value float64
}

type PairList []Pair

func (p PairList) Len() int           { return len(p) }
func (p PairList) Less(i, j int) bool { return p[i].Value < p[j].Value }
func (p PairList) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }


//saga deo


func (service *UserFollowersService) RedisConnection() {
	// create client and ping redis
	var err error
	client := redis.NewClient(&redis.Options{Addr: "localhost:6379"})
	if _, err = client.Ping().Result(); err != nil {
		log.Fatalf("error creating redis client %s", err)
	}

	// subscribe to the required channels
	pubsub := client.Subscribe(saga.UserFollowerChannel, saga.ReplyChannel)
	if _, err = pubsub.Receive(); err != nil {
		log.Fatalf("error subscribing %s", err)
	}
	defer func() { _ = pubsub.Close() }()
	ch := pubsub.Channel()

	log.Println("starting the user-follower service")
	for {
		select {
		case msg := <-ch:
			m := saga.Message{}
			err := json.Unmarshal([]byte(msg.Payload), &m)
			if err != nil {
				log.Println(err)
				continue
			}

			switch msg.Channel {
			case saga.UserFollowerChannel:

				// Happy Flow
				if m.Action == saga.ActionStart {

					err := service.Repository.CreateUserNodeIfNotExist(m.UserId)
					if err != nil{
						fmt.Println("Neuspesno upisao u neo4j")
						sendToReplyChannel(client, &m, saga.ActionRollback, saga.ServiceUser, saga.ServiceUserFollower)
					}else{
						fmt.Println("Uspesno upisao u neo4j")
						sendToReplyChannel(client, &m, saga.ActionDone, saga.ServiceUser, saga.ServiceUserFollower)
					}

				}

				// Rollback flow
				if m.Action == saga.ActionRollback {
					log.Printf("rolling back transaction")
				}

			}
		}
	}
}

func sendToReplyChannel(client *redis.Client, m *saga.Message, action string, service string, senderService string) {
	var err error
	m.Action = action
	m.Service = service
	m.SenderService = senderService
	if err = client.Publish(saga.ReplyChannel, m).Err(); err != nil {
		log.Printf("error publishing done-message to %s channel", saga.ReplyChannel)
	}
	log.Printf("done message published to channel :%s", saga.ReplyChannel)
}


