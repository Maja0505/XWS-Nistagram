package Handler

import (
	"XWS-Nistagram/AgentService/DTO"
	"XWS-Nistagram/AgentService/Service"
	"encoding/json"
	"fmt"
	"github.com/gocql/gocql"
	"github.com/gorilla/mux"
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

func (handler *AgentHandler) CreateCampaignRequest(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var requestDTO DTO.RequestDTO
	err := json.NewDecoder(r.Body).Decode(&requestDTO)
	if err != nil{
		fmt.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	err = handler.Service.CreateCampaignRequest(&requestDTO)
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

func (handler *AgentHandler) GetCampaignsForUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	vars := mux.Vars(r)
	userid := vars["id"]
	if userid == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	camp, err := handler.Service.GetCampaignsForUser(userid)
	if err != nil{
		fmt.Println(err)
		w.WriteHeader(http.StatusExpectationFailed)
		return
	}
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(camp)
}

func (handler *AgentHandler) GetCampaignRequests(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	vars := mux.Vars(r)
	userid := vars["id"]
	if userid == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	camp, err := handler.Service.GetCampaignRequests(userid)
	if err != nil{
		fmt.Println(err)
		w.WriteHeader(http.StatusExpectationFailed)
		return
	}
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(camp)
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


func (handler *AgentHandler) GetCampaignById(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	vars := mux.Vars(r)
	camapignIdString := vars["id"]
	if camapignIdString == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	camp, err := handler.Service.GetCampaignByID(camapignIdString)
	if err != nil{
		fmt.Println(err)
		w.WriteHeader(http.StatusExpectationFailed)
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(camp)
}

func ParseUUID(input string) (gocql.UUID, error) {
	var u gocql.UUID
	j := 0
	for _, r := range input {
		switch {
		case r == '-' && j&1 == 0:
			continue
		case r >= '0' && r <= '9' && j < 32:
			u[j/2] |= byte(r-'0') << uint(4-j&1*4)
		case r >= 'a' && r <= 'f' && j < 32:
			u[j/2] |= byte(r-'a'+10) << uint(4-j&1*4)
		case r >= 'A' && r <= 'F' && j < 32:
			u[j/2] |= byte(r-'A'+10) << uint(4-j&1*4)
		default:
			return gocql.UUID{}, fmt.Errorf("invalid UUID %q", input)
		}
		j += 1
	}
	if j != 32 {
		return gocql.UUID{}, fmt.Errorf("invalid UUID %q", input)
	}
	return u, nil
}
