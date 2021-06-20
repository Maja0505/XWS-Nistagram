package Service

import (
	"XWS-Nistagram/PostService/DTO"
	"XWS-Nistagram/PostService/Mapper"
	"XWS-Nistagram/PostService/Model"
	"XWS-Nistagram/PostService/Repository"
	"github.com/gocql/gocql"
)

type StoryService struct {
	Repo Repository.StoryRepository
}

func (service *StoryService) Create(storyDTO *DTO.CreateStoryDTO) error {
	story := Mapper.ConvertCreateStoryDTOToPost(storyDTO)
	err := service.Repo.Create(story)
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

func (service *StoryService) GetAllStoriesByUser(userId string) (*[]Model.Story,error){
	stories,err := service.Repo.GetAllStoriesByUser(userId)
	if err != nil {
		return nil, err
	}
	return stories, err
}

func (service *StoryService) GetAllNotExpiredStoriesByUser(userId string) (*[]Model.Story,error){
	stories,err := service.Repo.GetAllNotExpiredStoriesByUser(userId)
	if err != nil {
		return nil, err
	}
	return stories, err
}

func (service *StoryService) GetAllStoriesForCloseFriendsByUser(userId string) (*[]Model.Story,error){
	stories,err := service.Repo.GetAllStoriesForCloseFriendsByUser(userId)
	if err != nil {
		return nil, err
	}
	return stories, err
}

func (service *StoryService) GetAllHighlightsStoriesByUser(userId string) (*[]Model.Story,error){
	stories,err := service.Repo.GetAllHighlightsStoriesByUser(userId)
	if err != nil {
		return nil, err
	}
	return stories, err
}

