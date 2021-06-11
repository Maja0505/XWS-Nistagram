package handler

import (
	"XWS-Nistagram/UserFollowersService/dto"
	"XWS-Nistagram/UserFollowersService/service"
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
)

type UserFollowersHandler struct{
	Service *service.UserFollowersService
}

func (handler *UserFollowersHandler) FollowUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var data dto.FollowRelationshipDTO
	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		w.WriteHeader(http.StatusExpectationFailed)
		return
	}

	err = handler.Service.FollowUser(&data)
	if err != nil{
		w.WriteHeader(http.StatusExpectationFailed)
		return
	}

	w.WriteHeader(http.StatusCreated)


}

func (handler *UserFollowersHandler) UnfollowUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var data dto.UnfollowRelationshipDTO
	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		w.WriteHeader(http.StatusExpectationFailed)
		return
	}

	err = handler.Service.UnfollowUser(&data)
	if err != nil{
		w.WriteHeader(http.StatusExpectationFailed)
		return
	}

	w.WriteHeader(http.StatusOK)


}

func (handler *UserFollowersHandler) GetAllFollowedUsers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	vars := mux.Vars(r)
	userId := vars["userId"]

	if userId == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	//potrebno je pozvati http metodu iz userService
	users,err := handler.Service.GetAllFollowedUsers(userId)

	if err != nil{
		w.WriteHeader(http.StatusExpectationFailed)
		return
	}

	json.NewEncoder(w).Encode(users)
	w.WriteHeader(http.StatusOK)

}

func (handler *UserFollowersHandler) GetAllFollowersByUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	vars := mux.Vars(r)
	userId := vars["userId"]

	if userId == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	//potrebno je pozvati http metodu iz userService
	users,err := handler.Service.GetAllFollowersByUser(userId)

	if err != nil{
		w.WriteHeader(http.StatusExpectationFailed)
		return
	}

	json.NewEncoder(w).Encode(users)
	w.WriteHeader(http.StatusOK)
}

func (handler *UserFollowersHandler) AcceptFollowRequest(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var data dto.AcceptFollowRequestDTO
	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		w.WriteHeader(http.StatusExpectationFailed)
		return
	}

	err = handler.Service.AcceptFollowRequest(&data)
	if err != nil{
		w.WriteHeader(http.StatusExpectationFailed)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

func (handler *UserFollowersHandler) GetAllFollowRequsts(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	vars := mux.Vars(r)
	userId := vars["userId"]

	if userId == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	//potrebno je pozvati http metodu iz userService
	followRequests,err := handler.Service.GetAllFollowRequests(userId)

	if err != nil{
		w.WriteHeader(http.StatusExpectationFailed)
		return
	}

	json.NewEncoder(w).Encode(followRequests)
	w.WriteHeader(http.StatusOK)
}

func (handler *UserFollowersHandler) CheckFollowing(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	vars := mux.Vars(r)
	userId := vars["userId"]
	followedUserId := vars["followedUserId"]

	if userId == "" || followedUserId == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	following,err := handler.Service.CheckFollowing(userId,followedUserId)

	if err != nil{
		w.WriteHeader(http.StatusExpectationFailed)
		return
	}

	json.NewEncoder(w).Encode(following)
	w.WriteHeader(http.StatusOK)
}