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
		http.Error(w,err.Error(),417)
		return
	}

	w.WriteHeader(http.StatusOK)


}

func (handler *UserFollowersHandler) AcceptFollowRequest(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var data dto.FollowRequestDTO
	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		w.WriteHeader(http.StatusExpectationFailed)
		return
	}

	err = handler.Service.AcceptFollowRequest(&data)
	if err != nil{
		http.Error(w,err.Error(),417)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (handler *UserFollowersHandler) CancelFollowRequest(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var data dto.FollowRequestDTO
	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		w.WriteHeader(http.StatusExpectationFailed)
		return
	}

	err = handler.Service.CancelFollowRequest(&data)
	if err != nil{
		http.Error(w,err.Error(),417)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (handler *UserFollowersHandler) SetCloseFriend(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var data dto.CloseFriendDTO
	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		w.WriteHeader(http.StatusExpectationFailed)
		return
	}

	err = handler.Service.SetCloseFriend(&data)
	if err != nil{
		http.Error(w,err.Error(),417)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (handler *UserFollowersHandler) SetMuteFriend(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var data dto.MuteFriendDTO
	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		w.WriteHeader(http.StatusExpectationFailed)
		return
	}

	err = handler.Service.SetMuteFriend(&data)
	if err != nil{
		http.Error(w,err.Error(),417)
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

func (handler *UserFollowersHandler) GetAllNotMutedFollowedUsers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	vars := mux.Vars(r)
	userId := vars["userId"]

	if userId == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	users,err := handler.Service.GetAllNotMutedFollowedUsersByUser(userId)

	if err != nil{
		w.WriteHeader(http.StatusExpectationFailed)
		return
	}

	json.NewEncoder(w).Encode(users)
	w.WriteHeader(http.StatusOK)

}

func (handler *UserFollowersHandler) GetAllFollowsWhomUserIsCloseFriend(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	vars := mux.Vars(r)
	userId := vars["userId"]

	if userId == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	users,err := handler.Service.GetAllFollowsWhomUserIsCloseFriend(userId)

	if err != nil{
		w.WriteHeader(http.StatusExpectationFailed)
		return
	}

	json.NewEncoder(w).Encode(users)
	w.WriteHeader(http.StatusOK)

}

func (handler *UserFollowersHandler) GetAllFollowsWhomUserIsNotCloseFriend(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	vars := mux.Vars(r)
	userId := vars["userId"]

	if userId == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	users,err := handler.Service.GetAllFollowsWhomUserIsNotCloseFriend(userId)

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

func (handler *UserFollowersHandler) GetAllFollowRequests(w http.ResponseWriter, r *http.Request) {
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

func (handler *UserFollowersHandler) GetAllCloseFriends(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	vars := mux.Vars(r)
	userId := vars["userId"]

	if userId == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	//potrebno je pozvati http metodu iz userService
	closeFriends,err := handler.Service.GetAllCloseFriends(userId)

	if err != nil{
		w.WriteHeader(http.StatusExpectationFailed)
		return
	}

	json.NewEncoder(w).Encode(closeFriends)
	w.WriteHeader(http.StatusOK)
}

func (handler *UserFollowersHandler) GetAllMuteFriends(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	vars := mux.Vars(r)
	userId := vars["userId"]

	if userId == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	//potrebno je pozvati http metodu iz userService
	muteFriends,err := handler.Service.GetAllMuteFriends(userId)

	if err != nil{
		w.WriteHeader(http.StatusExpectationFailed)
		return
	}

	json.NewEncoder(w).Encode(muteFriends)
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

func (handler *UserFollowersHandler) CheckRequested(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	vars := mux.Vars(r)
	userId := vars["userId"]
	requestedUserId := vars["requestedUserId"]

	if userId == "" || requestedUserId == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	requested,err := handler.Service.CheckRequested(userId,requestedUserId)

	if err != nil{
		w.WriteHeader(http.StatusExpectationFailed)
		return
	}

	json.NewEncoder(w).Encode(requested)
	w.WriteHeader(http.StatusOK)
}

func (handler *UserFollowersHandler) CheckMuted(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	vars := mux.Vars(r)
	userId := vars["userId"]
	mutedUserId := vars["mutedUserId"]

	if userId == "" || mutedUserId == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	muted,err := handler.Service.CheckMuted(userId,mutedUserId)

	if err != nil{
		w.WriteHeader(http.StatusExpectationFailed)
		return
	}

	json.NewEncoder(w).Encode(muted)
	w.WriteHeader(http.StatusOK)
}

func (handler *UserFollowersHandler) CheckClosed(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	vars := mux.Vars(r)
	userId := vars["userId"]
	closedUserId := vars["closedUserId"]

	if userId == "" || closedUserId == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	muted,err := handler.Service.CheckClosed(userId,closedUserId)

	if err != nil{
		w.WriteHeader(http.StatusExpectationFailed)
		return
	}

	json.NewEncoder(w).Encode(muted)
	w.WriteHeader(http.StatusOK)
}