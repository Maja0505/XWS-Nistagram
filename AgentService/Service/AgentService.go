package Service

import (
	"XWS-Nistagram/AgentService/DTO"
	"XWS-Nistagram/AgentService/Mapper"
	"XWS-Nistagram/AgentService/Model"
	"XWS-Nistagram/AgentService/Repository"
	"fmt"
)

type AgentService struct {
	Repo Repository.AgentRepository
}

func (service *AgentService) CreateCampaign(campaignDTO *DTO.CampaignDTO) error {
	campaign := Mapper.ConvertCampaignDTOToCampaign(campaignDTO)
	err := service.Repo.CreateCampaign(campaign)
	if err != nil{
		fmt.Println(err)
		return  err
	}
	return nil
}

func (service *AgentService) CreateCampaignRequest(request *DTO.RequestDTO) error {
	err := service.Repo.CreateCampaignRequest(request.UserID, request.CampaignID)
	if err != nil{
		fmt.Println(err)
		return  err
	}
	return nil
}

func (service *AgentService) DeleteCampaign(campaignDTO *DTO.CampaignDTO) error {
	campaign := Mapper.ConvertCampaignDTOToCampaign(campaignDTO)
	err := service.Repo.DeleteCampaign(campaign)
	if err != nil{
		fmt.Println(err)
		return  err
	}
	return nil
}

func (service *AgentService) AddCampaignInfluencer(influencerDTO *DTO.AddInfluencerDTO) error {
	err := service.Repo.AddCampaignInfluencer(influencerDTO.InfluencerID, influencerDTO.UserID, influencerDTO.ID)
	if err != nil{
		fmt.Println(err)
		return  err
	}
	return nil
}

func (service *AgentService) GetCampaignsForUser(userid string) ( *[]Model.Campaign, error) {
	camp, err := service.Repo.GetCampaignsForUser(userid)
	if err != nil{
		fmt.Println(err)
		return nil, err
	}
	return camp, nil
}

func (service *AgentService) GetCampaignRequests(userid string) ( *[]DTO.RequestDTO, error) {
	camp, err := service.Repo.GetCampaignRequests(userid)
	if err != nil{
		fmt.Println(err)
		return nil, err
	}
	return camp, nil
}
