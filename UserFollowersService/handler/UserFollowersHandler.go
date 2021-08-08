package handler

import (
	"XWS-Nistagram/UserFollowersService/dto"
	"XWS-Nistagram/UserFollowersService/service"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/neo4j/neo4j-go-driver/v4/neo4j"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"time"
)

type UserFollowersHandler struct{
	Service *service.UserFollowersService
}

func (handler *UserFollowersHandler) CheckAuthorize(w http.ResponseWriter,r *http.Request) bool {
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

func (handler *UserFollowersHandler) FollowUser(w http.ResponseWriter, r *http.Request) {
	if !handler.CheckAuthorize(w,r){
		return
	}

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
	if !handler.CheckAuthorize(w,r){
		return
	}

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
	if !handler.CheckAuthorize(w,r){
		return
	}

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
	if !handler.CheckAuthorize(w,r){
		return
	}

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
	if !handler.CheckAuthorize(w,r){
		return
	}

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
	if !handler.CheckAuthorize(w,r){
		return
	}

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
	followRequests,err := handler.Service.GetAllFollowRequests(userId)

	if err != nil{
		w.WriteHeader(http.StatusExpectationFailed)
		return
	}

	json.NewEncoder(w).Encode(followRequests)
	w.WriteHeader(http.StatusOK)
}

func (handler *UserFollowersHandler) GetAllCloseFriends(w http.ResponseWriter, r *http.Request) {
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
	closeFriends,err := handler.Service.GetAllCloseFriends(userId)

	if err != nil{
		w.WriteHeader(http.StatusExpectationFailed)
		return
	}

	json.NewEncoder(w).Encode(closeFriends)
	w.WriteHeader(http.StatusOK)
}

func (handler *UserFollowersHandler) GetAllMuteFriends(w http.ResponseWriter, r *http.Request) {
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

func (handler *UserFollowersHandler) GetFollowSuggestions(w http.ResponseWriter, r *http.Request) {
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

	users,err := handler.Service.FollowSuggestions(userId)

	if err != nil{
		w.WriteHeader(http.StatusExpectationFailed)
		return
	}

	json.NewEncoder(w).Encode(users)
	w.WriteHeader(http.StatusOK)
}

func (handler *UserFollowersHandler) TestHttp(w http.ResponseWriter, r *http.Request) {

	session := handler.Service.Repository.Driver.NewSession(neo4j.SessionConfig{
		AccessMode:   neo4j.AccessModeRead,
	})
	defer session.Close()

	query := `match (u1:User)-[r:follow]->(u2:User)
 			  where u1.UserId='nemanja' and u2.UserId='pera' return r`

	fmt.Println("Zapoceto izvrsavanje sypher-shell-a : ",time.Now())
	_, err := session.ReadTransaction(func(tx neo4j.Transaction) (interface{}, error) {
		records, err := tx.Run(query, map[string]interface{}{})
		if err != nil {
			return nil, err
		}
		if records.Next() {
			return true,nil
		}
		return nil, nil
	})
	fmt.Println("Zavrseno izvrsavanje sypher-shell-a : ",time.Now())

	if err != nil {
		log.Println("error querying graph:", err)
		return
	}
	/*err = json.NewEncoder(w).Encode(d3Resp)
	if err != nil {
		log.Println("error writing graph response:", err)
	}
	w.WriteHeader(http.StatusOK)*/
}


func (handler *UserFollowersHandler) Test() {

	session := handler.Service.Repository.Driver.NewSession(neo4j.SessionConfig{
		AccessMode:   neo4j.AccessModeRead,
	})
	defer session.Close()

	query := `match (u1:User)-[r:follow]->(u2:User)
 			  where u1.UserId='nemanja' and u2.UserId='pera' return r`

	fmt.Println("Zapoceto izvrsavanje sypher-shell-a : ",time.Now())
	_, err := session.ReadTransaction(func(tx neo4j.Transaction) (interface{}, error) {
		records, err := tx.Run(query, map[string]interface{}{})
		if err != nil {
			return nil, err
		}
		if records.Next() {
			return true,nil
		}
		return nil, nil
	})
	fmt.Println("Zavrseno izvrsavanje sypher-shell-a : ",time.Now())

	if err != nil {
		log.Println("error querying graph:", err)
		return
	}
	/*err = json.NewEncoder(w).Encode(d3Resp)
	if err != nil {
		log.Println("error writing graph response:", err)
	}
	w.WriteHeader(http.StatusOK)*/
}