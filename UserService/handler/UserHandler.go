package handler

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"userService/model"
	"userService/service"
)

type UserHandler struct {
	Service *service.UserService
}



func (handler *UserHandler) FindAll(w http.ResponseWriter,r *http.Request){
	w.Header().Set("Content-Type", "application/json")
	users,err := handler.Service.FindAll()
	if err != nil{
		fmt.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(users)
}

func (handler *UserHandler) Create(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var user model.User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil{
		fmt.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	err = handler.Service.Create(&user)
	if err != nil{
		fmt.Println(err)
		w.WriteHeader(http.StatusExpectationFailed)
		return
	}

	w.WriteHeader(http.StatusCreated)

}

func (handler *UserHandler) Update(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	if id == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	var user model.User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil{
		fmt.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	err = handler.Service.Update(id,&user)
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusExpectationFailed)
		return
	}
	w.WriteHeader(http.StatusOK)
}

func (handler *UserHandler) FindUserByUsername(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")
	vars := mux.Vars(r)
	username := vars["username"]
	if username == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	user,err := handler.Service.FindUserByUsername(username)
	if err != nil{
		w.WriteHeader(http.StatusExpectationFailed)
		return
	}
	if user == nil{
		w.WriteHeader(http.StatusNotFound)
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(user)
}