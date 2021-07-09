package handler

import (
	"XWS-Nistagram/UserFollowersService/model"
	"XWS-Nistagram/UserFollowersService/service"
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
)

type BlockedUserHandler struct{
	Service *service.BlockedUserService
}

func (handler *BlockedUserHandler) BlockUser(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")

	var br model.BlockRelationship
	err := json.NewDecoder(r.Body).Decode(&br)
	if err != nil {
		w.WriteHeader(http.StatusExpectationFailed)
		return
	}

	err = handler.Service.BlockUser(&br)
	if err != nil{
		w.WriteHeader(http.StatusExpectationFailed)
		return
	}

	w.WriteHeader(http.StatusCreated)

}

func (handler *BlockedUserHandler) GetAllBlockedUsers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	vars := mux.Vars(r)
	userId := vars["userId"]

	if userId == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	//potrebno je pozvati http metodu iz userService
	blockedUsers,err := handler.Service.GetAllBlockedUser(userId)

	if err != nil{
		w.WriteHeader(http.StatusExpectationFailed)
		return
	}

	json.NewEncoder(w).Encode(blockedUsers)
	w.WriteHeader(http.StatusOK)

}

func (handler *BlockedUserHandler) CheckBlock(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	vars := mux.Vars(r)
	userId := vars["userId"]
	blockedUserId := vars["blockedUserId"]

	if userId == "" || blockedUserId == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	block,err := handler.Service.CheckBlock(userId,blockedUserId)

	if err != nil{
		w.WriteHeader(http.StatusExpectationFailed)
		return
	}

	json.NewEncoder(w).Encode(block)
	w.WriteHeader(http.StatusOK)
}

func (handler *BlockedUserHandler) UnblockUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var data model.BlockRelationship
	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		w.WriteHeader(http.StatusExpectationFailed)
		return
	}

	err = handler.Service.UnblockUser(&data)
	if err != nil{
		http.Error(w,err.Error(),417)
		return
	}

	w.WriteHeader(http.StatusOK)
}