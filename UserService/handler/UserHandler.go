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
	"userService/saga"
	"userService/service"
)

type UserHandler struct {
	Service *service.UserService
}

func (handler *UserHandler) CheckAuthorize(w http.ResponseWriter,r *http.Request){
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

}


func (handler *UserHandler) FindAll(w http.ResponseWriter,r *http.Request){
	handler.CheckAuthorize(w,r)
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

func (handler *UserHandler) CreateRegisteredUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var userForRegistrationDTO dto.UserForRegistrationDTO
	err := json.NewDecoder(r.Body).Decode(&userForRegistrationDTO)
	if err != nil{
		fmt.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(err)
		return
	}
	username,password,isAgent,idString,err := handler.Service.CreateRegisteredUser(&userForRegistrationDTO)
	if err != nil{
		fmt.Println(err)
		http.Error(w,err.Error(),417)
		return
	}
	var role string

	if isAgent {
		role = "AGENT"
	}else{
		role = "USER"
	}

	if username == "admin"{
		role = "ADMIN"
	}
	fmt.Println("Treba da posalje na auth kanal  za upis usera u auth!")
	m := saga.Message{Service: saga.ServiceAuthentication, SenderService: saga.ServiceUser, Action: saga.ActionStart, UserId: idString , Username: username,Password: password,Role: role}
	handler.Service.Orchestrator.Next(saga.AuthenticationChannel,saga.ServiceAuthentication,m)

	w.WriteHeader(http.StatusCreated)

}

func (handler *UserHandler) UpdateRegisteredUserProfile(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	username := vars["username"]
	if username == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	var registeredUserDto dto.RegisteredUserProfileInfoDTO
	err := json.NewDecoder(r.Body).Decode(&registeredUserDto)
	if err != nil{
		fmt.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	err = handler.Service.UpdateRegisteredUserProfile(username,&registeredUserDto)
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
	user,_ := handler.Service.FindUserByUsername(username)

	if user == nil{
		w.WriteHeader(http.StatusNotFound)
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(user)
}

func (handler *UserHandler) FindUserByUserId(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")
	vars := mux.Vars(r)
	userId := vars["userId"]
	if userId == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	user,_ := handler.Service.FindUserByUserId(userId)

	if user == nil{
		w.WriteHeader(http.StatusNotFound)
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(user)
}

func (handler *UserHandler) FindUsernameAndProfilePicture(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")
	vars := mux.Vars(r)
	userId := vars["userId"]
	if userId == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	user,_ := handler.Service.FindUserByUserIdAndGetHisUsernameAndProfilePicture(userId)

	if user == nil{
		w.WriteHeader(http.StatusNotFound)
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(user)
}

func (handler *UserHandler) SearchUser(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")
	vars := mux.Vars(r)
	searchContent := vars["searchContent"]
	username := vars["username"]
	if searchContent == ""{
		err := errors.New("No results found")
		http.Error(w,err.Error(),400)
		return
	}
	users,err := handler.Service.SearchUser(username,searchContent)
	if err != nil{
		w.WriteHeader(http.StatusExpectationFailed)
		return
	}
	if   len(*users) == 0 {
		err := errors.New("No results found")
		http.Error(w,err.Error(),400)
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(users)
}

func (handler *UserHandler) ConvertUserIdsToUsers(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")
	var userIds dto.UserIdsDTO
	err := json.NewDecoder(r.Body).Decode(&userIds)
	if err != nil{
		w.WriteHeader(http.StatusBadRequest)
	}
	fmt.Println(userIds)
	users,err := handler.Service.ConvertUserIdsToUsers(userIds)
	if err != nil {
		w.WriteHeader(http.StatusExpectationFailed)
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(users)
}

func (handler *UserHandler) ConvertUsernamesToUsers(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")
	var usernames dto.UsernamesDTO
	err := json.NewDecoder(r.Body).Decode(&usernames)
	if err != nil{
		w.WriteHeader(http.StatusBadRequest)
		fmt.Println("aaaa")
	}
	fmt.Println("usernames  ",usernames)
	fmt.Println(usernames)
	users,err := handler.Service.ConvertUsernamesToUsers(usernames)
	if err != nil {
		w.WriteHeader(http.StatusExpectationFailed)
		fmt.Println("aaaaa")
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(users)
}

func (handler *UserHandler) ChangePassword(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type","application/json")
	vars := mux.Vars(r)
	username := vars["username"]
	var passwordDto dto.PasswordDTO
	err := json.NewDecoder(r.Body).Decode(&passwordDto)
	if err != nil{
		w.WriteHeader(http.StatusBadRequest)
	}
	valid,err := handler.Service.ChangePassword(username,passwordDto)
	if valid{
		if err != nil{
			http.Error(w,err.Error(),400)
			return
		}
		w.WriteHeader(http.StatusOK)
		return
	}
	if err != nil {
		w.WriteHeader(http.StatusExpectationFailed)
		return
	}


}



func (handler *UserHandler) UpdatePublicProfileSetting(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	setting := vars["setting"]
	username := vars["username"]

	err := handler.Service.UpdatePublicProfileSetting(username,setting)
	if err != nil {
		w.WriteHeader(http.StatusExpectationFailed)
	}
	w.WriteHeader(http.StatusOK)
}

func (handler *UserHandler) UpdateMessageRequestSetting(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	setting := vars["setting"]
	username := vars["username"]
	err := handler.Service.UpdateMessageRequestSetting(username,setting)

	if err != nil {
		w.WriteHeader(http.StatusExpectationFailed)
	}
	w.WriteHeader(http.StatusOK)
}

func (handler *UserHandler) UpdateAllowTagsSetting(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	setting := vars["setting"]
	username := vars["username"]
	err := handler.Service.UpdateAllowTagsSetting(username,setting)

	if err != nil {
		w.WriteHeader(http.StatusExpectationFailed)
	}
	w.WriteHeader(http.StatusOK)
}

func (handler *UserHandler) UpdateLikeNotificationSetting(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	setting := vars["setting"]
	username := vars["username"]
	err := handler.Service.UpdateLikeNotificationSetting(username,setting)

	if err != nil {
		w.WriteHeader(http.StatusExpectationFailed)
	}
	w.WriteHeader(http.StatusOK)
}

func (handler *UserHandler) UpdateCommentNotificationSetting(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	setting := vars["setting"]
	username := vars["username"]
	err := handler.Service.UpdateCommentNotificationSetting(username,setting)

	if err != nil {
		w.WriteHeader(http.StatusExpectationFailed)
	}
	w.WriteHeader(http.StatusOK)
}

func (handler *UserHandler) UpdateMessageRequestNotificationSetting(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	setting := vars["setting"]
	username := vars["username"]
	err := handler.Service.UpdateMessageRequestNotificationSetting(username,setting)

	if err != nil {
		w.WriteHeader(http.StatusExpectationFailed)
	}
	w.WriteHeader(http.StatusOK)
}

func (handler *UserHandler) UpdateMessageNotificationSetting(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	setting := vars["setting"]
	username := vars["username"]
	err := handler.Service.UpdateMessageNotificationSetting(username,setting)

	if err != nil {
		w.WriteHeader(http.StatusExpectationFailed)
	}
	w.WriteHeader(http.StatusOK)

}

func (handler *UserHandler) UpdateFollowRequestNotificationSetting(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	setting := vars["setting"]
	username := vars["username"]
	err := handler.Service.UpdateFollowRequestNotificationSetting(username,setting)

	if err != nil {
		w.WriteHeader(http.StatusExpectationFailed)
	}
	w.WriteHeader(http.StatusOK)
}

func (handler *UserHandler) UpdateFollowNotificationSetting(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	setting := vars["setting"]
	username := vars["username"]
	err := handler.Service.UpdateFollowNotificationSetting(username,setting)

	if err != nil {
		w.WriteHeader(http.StatusExpectationFailed)
	}
	w.WriteHeader(http.StatusOK)
}

func (handler *UserHandler) DeleteUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userId := vars["userId"]

	if userId == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	err := handler.Service.DeleteUserByUserId(userId)

	if err != nil {
		w.WriteHeader(http.StatusExpectationFailed)
	}
	w.WriteHeader(http.StatusOK)
}

func (handler *UserHandler) UploadImage(w http.ResponseWriter,r *http.Request){
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

	dst, err := os.Create("profile-docs/" + imagePath +".jpg")
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

func (handler *UserHandler) GetImage(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "image/jpeg")
	vars := mux.Vars(r)
	imagepath := vars["id"]
	if imagepath == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	file,err := ioutil.ReadFile("profile-docs/" + imagepath)
	if err != nil{
		w.WriteHeader(http.StatusBadRequest)
		fmt.Println(err)
		return
	}
	w.Write(file)


}

