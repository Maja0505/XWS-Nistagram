package handler

import (
	"XWS-Nistagram/UserFollowersService/model"
	"XWS-Nistagram/UserFollowersService/service"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/gorilla/mux"
	"io/ioutil"
	"net/http"
	"os"
)

type BlockedUserHandler struct{
	Service *service.BlockedUserService
}

func (handler *BlockedUserHandler) CheckAuthorize(w http.ResponseWriter,r *http.Request) bool {
	client := &http.Client{}
	reqUrl := fmt.Sprintf("http://" +os.Getenv("AUTHENTICATION_SERVICE_DOMAIN") + ":" + os.Getenv("AUTHENTICATION_SERVICE_PORT")+ "/authorize")
	req,err := http.NewRequest("POST",reqUrl,nil)
	req.Header.Add("Authorization",r.Header.Get("Authorization"))
	req.Header.Add("path","/api/user-follow" + r.URL.Path)
	req.Header.Add("method",r.Method)

	fmt.Println(r.Method)
	resp,err := client.Do(req)
	if err != nil{
		fmt.Println(err)
	}
	fmt.Println(resp.Body)
	fmt.Println(resp.Status)
	fmt.Println(resp.Header)

	if resp.StatusCode != 200 {
		var errorText string
		body, _ := ioutil.ReadAll(resp.Body)
		respBodyInErrorCase := json.Unmarshal(body, &errorText)
		respBodyInErrorCase = errors.New(errorText)
		http.Error(w,respBodyInErrorCase.Error(),resp.StatusCode)
		return false
	}
	return true

}

func (handler *BlockedUserHandler) BlockUser(w http.ResponseWriter, r *http.Request){
	if !handler.CheckAuthorize(w,r){
		return
	}

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
	if !handler.CheckAuthorize(w,r){
		return
	}

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
	if !handler.CheckAuthorize(w,r){
		return
	}

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