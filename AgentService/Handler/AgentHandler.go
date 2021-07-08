package Handler

import (
	"XWS-Nistagram/AgentService/DTO"
	"XWS-Nistagram/AgentService/Service"
	"encoding/json"
	"fmt"
	"net/http"
)

type AgentHandler struct {
	Service *Service.AgentService
}

func (handler *AgentHandler) CreateCampaign(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var campaignDTO DTO.CampaignDTO
	err := json.NewDecoder(r.Body).Decode(&campaignDTO)
	if err != nil{
		fmt.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	err = handler.Service.CreateCampaign(&campaignDTO)
	if err != nil{
		fmt.Println(err)
		w.WriteHeader(http.StatusExpectationFailed)
		return
	}
	w.WriteHeader(http.StatusCreated)
}

func (handler *AgentHandler) DeleteCampaign(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var campaignDTO DTO.CampaignDTO
	err := json.NewDecoder(r.Body).Decode(&campaignDTO)
	if err != nil{
		fmt.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	err = handler.Service.DeleteCampaign(&campaignDTO)
	if err != nil{
		fmt.Println(err)
		w.WriteHeader(http.StatusExpectationFailed)
		return
	}
	w.WriteHeader(http.StatusCreated)
}

func (handler *AgentHandler) AddCampaignInfluencer(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var influencerDTO DTO.AddInfluencerDTO
	err := json.NewDecoder(r.Body).Decode(&influencerDTO)
	if err != nil{
		fmt.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	err = handler.Service.AddCampaignInfluencer(&influencerDTO)
	if err != nil{
		fmt.Println(err)
		w.WriteHeader(http.StatusExpectationFailed)
		return
	}
	w.WriteHeader(http.StatusCreated)
}