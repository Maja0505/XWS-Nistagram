package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"userService/dto"
	"userService/service"
)

type VerificationRequestHandler struct {
	Service *service.VerificationRequestService
}


func (handler *VerificationRequestHandler) Create(w http.ResponseWriter,r *http.Request){
	w.Header().Set("Content-Type", "application/json")
	var vrDTO dto.VerificationRequestDTO
	err := json.NewDecoder(r.Body).Decode(&vrDTO)
	if err != nil{
		fmt.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	err = handler.Service.Create(&vrDTO)
	if err != nil{
		fmt.Println(err)
		w.WriteHeader(http.StatusExpectationFailed)
		return
	}

	w.WriteHeader(http.StatusCreated)
}