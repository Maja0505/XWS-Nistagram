package handler

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/gorilla/mux"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"userService/dto"
	"userService/model"
	"userService/service"
)

type VerificationRequestHandler struct {
	Service *service.VerificationRequestService
}

func (handler *VerificationRequestHandler) CheckAuthorize(w http.ResponseWriter,r *http.Request) {
	client := &http.Client{}
	reqUrl := fmt.Sprintf("http://" +os.Getenv("AUTHENTICATION_SERVICE_DOMAIN") + ":" + os.Getenv("AUTHENTICATION_SERVICE_PORT")+ "/authorize")
	req,err := http.NewRequest("POST",reqUrl,nil)
	req.Header.Add("Authorization",r.Header.Get("Authorization"))
	req.Header.Add("path","/api/user" + r.URL.Path)
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
		return
	}


}


func (handler *VerificationRequestHandler) Create(w http.ResponseWriter,r *http.Request){
	handler.CheckAuthorize(w,r)
	w.Header().Set("Content-Type", "application/json")
	var vrDTO dto.VerificationRequestDTO
	err := json.NewDecoder(r.Body).Decode(&vrDTO)
	if err != nil{
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	err = handler.Service.Create(&vrDTO)
	if err != nil{
		http.Error(w,err.Error(),417)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

func (handler *VerificationRequestHandler) Update(w http.ResponseWriter,r *http.Request){
	handler.CheckAuthorize(w,r)

	w.Header().Set("Content-Type", "application/json")
	var vrDTO dto.VerificationRequestDTO
	vars := mux.Vars(r)

	user := vars["user"]
	if user == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	err := json.NewDecoder(r.Body).Decode(&vrDTO)
	if err != nil{
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	err = handler.Service.Update(user,&vrDTO)
	if err != nil{
		fmt.Println(err)
		w.WriteHeader(http.StatusExpectationFailed)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (handler *VerificationRequestHandler) GetAllVerificationRequest(w http.ResponseWriter,r *http.Request){
	handler.CheckAuthorize(w,r)

	w.Header().Set("Content-Type", "application/json")

	vrDto,err := handler.Service.GetAllVerificationRequests()
	if err != nil{
		fmt.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(vrDto)
}

func (handler *VerificationRequestHandler) GetVerificationRequestByUser(w http.ResponseWriter,r *http.Request){
	handler.CheckAuthorize(w,r)

	w.Header().Set("Content-Type", "application/json")
	vars := mux.Vars(r)
	user := vars["user"]

	vrDto,err := handler.Service.GetVerificationRequestByUser(user)
	if vrDto == nil {
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(vrDto)
		return
	}

	if err != nil{
		fmt.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(vrDto)
}

func (handler *VerificationRequestHandler) ApproveVerificationRequest(w http.ResponseWriter,r *http.Request){
	handler.CheckAuthorize(w,r)

	w.Header().Set("Content-Type", "application/json")
	vars := mux.Vars(r)
	user := vars["user"]
	if user == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	err := handler.Service.ApproveVerificationRequest(user)
	if err != nil{
		fmt.Println(err)
		w.WriteHeader(http.StatusExpectationFailed)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (handler *VerificationRequestHandler) DeleteVerificationRequest(w http.ResponseWriter,r *http.Request){
	handler.CheckAuthorize(w,r)

	w.Header().Set("Content-Type", "application/json")
	vars := mux.Vars(r)
	user := vars["user"]
	if user == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	err := handler.Service.DeleteVerificationRequest(user)
	if err != nil{
		fmt.Println(err)
		w.WriteHeader(http.StatusExpectationFailed)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (handler *VerificationRequestHandler) CreateAgentRegistrationRequest(w http.ResponseWriter,r *http.Request){
	handler.CheckAuthorize(w,r)

	w.Header().Set("Content-Type", "application/json")

	var aRR model.AgentRegistrationRequest
	err := json.NewDecoder(r.Body).Decode(&aRR)
	if err != nil{
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	err = handler.Service.CreateAgentRegistrationRequest(&aRR)
	if err != nil{
		fmt.Println(err)
		w.WriteHeader(http.StatusExpectationFailed)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

func (handler *VerificationRequestHandler) GetAllAgentRegistrationRequests(w http.ResponseWriter,r *http.Request){
	handler.CheckAuthorize(w,r)

	w.Header().Set("Content-Type", "application/json")

	verificationRequests,err := handler.Service.GetAllAgentRegistrationRequests()

	if err != nil{
		fmt.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(verificationRequests)
}

func (handler *VerificationRequestHandler) UpdateAgentRegistrationRequestToApproved(w http.ResponseWriter,r *http.Request){
	handler.CheckAuthorize(w,r)

	w.Header().Set("Content-Type", "application/json")
	vars := mux.Vars(r)
	username := vars["username"]

	if username == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	err := handler.Service.UpdateAgentRegistrationRequestToApproved(username)

	if err != nil{
		fmt.Println(err)
		w.WriteHeader(http.StatusExpectationFailed)
		return
	}
	w.WriteHeader(http.StatusOK)
}

func (handler *VerificationRequestHandler) DeleteAgentRegistrationRequestToApproved(w http.ResponseWriter,r *http.Request){
	handler.CheckAuthorize(w,r)

	w.Header().Set("Content-Type", "application/json")
	vars := mux.Vars(r)
	username := vars["username"]

	if username == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	err := handler.Service.DeleteAgentRegistrationRequestToApproved(username)

	if err != nil{
		fmt.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusOK)
}

func (handler *VerificationRequestHandler) UploadImage(w http.ResponseWriter,r *http.Request){
	vars := mux.Vars(r)
	imagePath := vars["id"]
	r.ParseMultipartForm(10 << 20)
	file, _, err := r.FormFile("myFile")
	if err != nil {
		fmt.Println("Error Retrieving the File")
		fmt.Println(err)
		return
	}

	defer file.Close()

	dst, err := os.Create("verification-docs/" + imagePath +".jpg")
	defer dst.Close()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if _, err := io.Copy(dst, file); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	fmt.Println("Successfully Uploaded File\n" + dst.Name())

	w.WriteHeader(http.StatusOK)


}

func (handler *VerificationRequestHandler) GetImage(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "image/jpeg")
	vars := mux.Vars(r)
	imagepath := vars["id"]
	if imagepath == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	file,err := ioutil.ReadFile("verification-docs/" + imagepath)
	if err != nil{
		w.WriteHeader(http.StatusBadRequest)
		fmt.Println(err)
		return
	}
	w.Write(file)


}
