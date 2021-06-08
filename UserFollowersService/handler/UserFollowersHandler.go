package handler

import (
	"XWS-Nistagram/UserFollowersService/events"
	"XWS-Nistagram/UserFollowersService/model"
	"encoding/json"
	"log"
	"net/http"
)

type UserFollowersHandler struct{}

func (handler *UserFollowersHandler) CreateEvent(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
	var data events.Event
	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		log.Println(err)
	}
	ce := make(chan error)
	// goroutine for invoking the model layer event create function
	go model.CreateEvent(data, ce)
	if err = <-ce; err != nil {
		log.Println(err)
		json.NewEncoder(w).Encode(struct {
			Status  bool   `json:"status"`
			Message string `json:"message"`
		}{false, "some error occurreed"})
		return
	}
	json.NewEncoder(w).Encode(struct {
		Status  bool   `json:"status"`
		Message string `json:"message"`
	}{true, "new node created successfully"})

}
