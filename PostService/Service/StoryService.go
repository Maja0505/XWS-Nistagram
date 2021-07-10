package Service

import (
	"XWS-Nistagram/PostService/DTO"
	"XWS-Nistagram/PostService/Mapper"
	"XWS-Nistagram/PostService/Repository"
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/gocql/gocql"
	"io/ioutil"
	"net/http"
	"os"
)

type StoryService struct {
	Repo Repository.StoryRepository
}

func (service *StoryService) Create(storyDTO *DTO.CreateStoryDTO) (gocql.UUID, error) {
	story := Mapper.ConvertCreateStoryDTOToPost(storyDTO)
	id, err := service.Repo.Create(story)
	if err != nil{
		return id, err
	}
	return id, nil
}

func (service *StoryService) UpdateStoryAvailabilityAndDate(dto *DTO.UpdateStoryAgentDTO) error {
	err := service.Repo.UpdateStoryAvailabilityAndDate(dto.ID, dto.UserID, dto.CreatedAt)
	if err != nil{
		return  err
	}
	return nil
}

func (service *StoryService) SetStoryForHighlights(stringId string) error {
	id,err := gocql.ParseUUID(stringId)
	if err != nil{
		return err
	}

	err = service.Repo.SetStoryForHighlights(id)
	if err != nil{
		return  err
	}
	return nil
}

func (service *StoryService) GetAllStoriesByUser(userId string) (*[]DTO.StoryDTO,error){
	stories,err := service.Repo.GetAllStoriesByUser(userId)
	if err != nil {
		return nil, err
	}
	storiesDtos := Mapper.ConvertStoryListToStoryDTOList(stories)
	return storiesDtos, nil
}

func (service *StoryService) GetAllNotExpiredStoriesByUser(userId string) (*[]DTO.StoryDTO,error){
	stories,err := service.Repo.GetAllNotExpiredStoriesByUser(userId)
	if err != nil {
		return nil, err
	}
	storiesDtos := Mapper.ConvertStoryListToStoryDTOList(stories)
	return storiesDtos, nil
}

func (service *StoryService) GetAllStoriesForCloseFriendsByUser(userId string) (*[]DTO.StoryDTO,error){
	fmt.Println("GetAllStoriesForCloseFriendsByUser")
	stories,err := service.Repo.GetAllStoriesForCloseFriendsByUser(userId)
	if err != nil {
		return nil, err
	}
	storiesDtos := Mapper.ConvertStoryListToStoryDTOList(stories)
	return storiesDtos, nil
}

func (service *StoryService) GetAllHighlightsStoriesByUser(userId string) (*[]DTO.StoryDTO,error){
	stories,err := service.Repo.GetAllHighlightsStoriesByUser(userId)
	if err != nil {
		return nil, err
	}
	storiesDtos := Mapper.ConvertStoryListToStoryDTOList(stories)
	return storiesDtos, nil
}

func (service *StoryService) GetAllFollowsWithStories(userid string) (*[]DTO.UserByUsernameDTO,error){

	var followsWithStories []string

	reqUrl1 := fmt.Sprintf("http://" + os.Getenv("USER_FOLLOWERS_SERVICE_DOMAIN") + ":" + os.Getenv("USER_FOLLOWERS_SERVICE_PORT") + "/allAllFollowsWhomUserIsNotCloseFriend/" + userid)
	reqUrl2 := fmt.Sprintf("http://" + os.Getenv("USER_FOLLOWERS_SERVICE_DOMAIN") + ":" + os.Getenv("USER_FOLLOWERS_SERVICE_PORT") + "/allAllFollowsWhomUserIsCloseFriend/" + userid)

	fmt.Println("Usao u metodu")

	resp, err := http.Get(reqUrl1)
	if err != nil || resp.StatusCode == 404 {
		return nil,err
	}
	body, err := ioutil.ReadAll(resp.Body)
	var followsWhomUserIsNotCloseFriend []string
	err = json.Unmarshal(body, &followsWhomUserIsNotCloseFriend)
	fmt.Println(followsWhomUserIsNotCloseFriend)
	if err != nil{
		return nil, err
	}

	resp, err = http.Get(reqUrl2)
	if err != nil || resp.StatusCode == 404 {
		return nil,err
	}
	body, err = ioutil.ReadAll(resp.Body)
	var followsWhomUserIsCloseFriend []string
	err = json.Unmarshal(body, &followsWhomUserIsCloseFriend)
	fmt.Println(followsWhomUserIsCloseFriend)
	if err != nil{
		return nil, err
	}

	for _, userID1 := range followsWhomUserIsNotCloseFriend {
		if service.Repo.CheckDoesUserHaveAnyNotExpiredStory(userID1){
			followsWithStories = append(followsWithStories,userID1)
		}
	}

	for _, userID2 := range followsWhomUserIsCloseFriend {
		if service.Repo.CheckDoesUserHaveAnyNotExpiredStoryForCloseFriends(userID2) || service.Repo.CheckDoesUserHaveAnyNotExpiredStory(userID2){
			followsWithStories = append(followsWithStories,userID2)
		}
	}

	usernamesDTOList,err := GetUsernamesByUserIdsFromUserService(&followsWithStories)

	if err != nil{
		return nil, err
	}

	return usernamesDTOList,nil


}

func GetUsernamesByUserIdsFromUserService (users *[]string) (*[]DTO.UserByUsernameDTO,error){

	reqUrl := fmt.Sprintf("http://" +os.Getenv("USER_SERVICE_DOMAIN") + ":" + os.Getenv("USER_SERVICE_PORT")+ "/convert-user-ids")

	type UserIdsDTO struct {
		UserIds []string
	}

	userIdsDto := UserIdsDTO{}
	userIdsDto.UserIds = *users
	jsonUserids,_ := json.Marshal(userIdsDto)
	resp, err := http.Post(reqUrl,"appliation/json",bytes.NewBuffer(jsonUserids))

	if err != nil || resp.StatusCode == 404 {
		return nil,err
	}
	body, err := ioutil.ReadAll(resp.Body)
	var data []DTO.UserByUsernameDTO
	err = json.Unmarshal(body, &data)
	if err != nil{
		return nil, err
	}
	return &data, nil
}