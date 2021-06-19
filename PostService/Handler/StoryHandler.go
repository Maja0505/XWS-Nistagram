package Handler

import (
	"XWS-Nistagram/PostService/DTO"
	"XWS-Nistagram/PostService/Service"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"io"
	"io/ioutil"
	"net/http"
	"os"
)

type StoryHandler struct {
	Service *Service.StoryService
}

func (handler *StoryHandler) Create(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var storyDTO DTO.CreateStoryDTO
	err := json.NewDecoder(r.Body).Decode(&storyDTO)
	if err != nil{
		fmt.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	err = handler.Service.Create(&storyDTO)
	if err != nil{
		fmt.Println(err)
		w.WriteHeader(http.StatusExpectationFailed)
		return
	}

	w.WriteHeader(http.StatusCreated)

}

func (handler *StoryHandler) SetStoryForHighlights(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	vars := mux.Vars(r)
	id := vars["id"]
	if id == ""{
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	err := handler.Service.SetStoryForHighlights(id)
	if err != nil{
		fmt.Println(err)
		w.WriteHeader(http.StatusExpectationFailed)
		return
	}

	w.WriteHeader(http.StatusOK)

}

func (handler *StoryHandler) GetAllStoriesByUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	vars := mux.Vars(r)
	userId := vars["userId"]
	if userId == ""{
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	stories,err := handler.Service.GetAllStoriesByUser(userId)
	if err != nil{
		fmt.Println(err)
		w.WriteHeader(http.StatusExpectationFailed)
		return
	}

	json.NewEncoder(w).Encode(stories)
	w.WriteHeader(http.StatusOK)

}

func (handler *StoryHandler) GetAllNotExpiredStoriesByUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	vars := mux.Vars(r)
	userId := vars["userId"]
	if userId == ""{
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	stories,err := handler.Service.GetAllNotExpiredStoriesByUser(userId)
	if err != nil{
		fmt.Println(err)
		w.WriteHeader(http.StatusExpectationFailed)
		return
	}

	json.NewEncoder(w).Encode(stories)
	w.WriteHeader(http.StatusOK)

}

func (handler *StoryHandler) GetAllStoriesForCloseFriendsByUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	vars := mux.Vars(r)
	userId := vars["userId"]
	if userId == ""{
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	stories,err := handler.Service.GetAllStoriesForCloseFriendsByUser(userId)
	if err != nil{
		fmt.Println(err)
		w.WriteHeader(http.StatusExpectationFailed)
		return
	}

	json.NewEncoder(w).Encode(stories)
	w.WriteHeader(http.StatusOK)

}

func (handler *StoryHandler) GetAllHighlightsStoriesByUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	vars := mux.Vars(r)
	userId := vars["userId"]
	if userId == ""{
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	stories,err := handler.Service.GetAllHighlightsStoriesByUser(userId)
	if err != nil{
		fmt.Println(err)
		w.WriteHeader(http.StatusExpectationFailed)
		return
	}

	json.NewEncoder(w).Encode(stories)
	w.WriteHeader(http.StatusOK)

}

func (handler *StoryHandler) UploadVideo(w http.ResponseWriter,r *http.Request){
	vars := mux.Vars(r)
	imagePath := vars["videoId"]
	r.ParseMultipartForm(10 << 20)
	file, _, err := r.FormFile("myFile")
	if err != nil {
		fmt.Println("Error Retrieving the File")
		fmt.Println(err)
		return
	}

	defer file.Close()

	dst, err := os.Create("Videos/" + imagePath +".mp4")
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

func (handler *StoryHandler) GetVideo(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "video/mp4")
	vars := mux.Vars(r)
	imagepath := vars["videoId"]
	if imagepath == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	file,err := ioutil.ReadFile("Videos/" + imagepath)
	if err != nil{
		w.WriteHeader(http.StatusBadRequest)
		fmt.Println(err)
		return
	}
	w.Write(file)


}

func (handler *StoryHandler) UploadImage(w http.ResponseWriter,r *http.Request){
	vars := mux.Vars(r)
	imagePath := vars["imageId"]
	r.ParseMultipartForm(10 << 20)
	file, _, err := r.FormFile("myFile")
	if err != nil {
		fmt.Println("Error Retrieving the File")
		fmt.Println(err)
		return
	}

	defer file.Close()

	dst, err := os.Create("Images/" + imagePath +".jpg")
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

func (handler *StoryHandler) GetImage(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "image/jpeg")
	vars := mux.Vars(r)
	imagepath := vars["imageId"]
	if imagepath == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	file,err := ioutil.ReadFile("Images/" + imagepath)
	if err != nil{
		w.WriteHeader(http.StatusBadRequest)
		fmt.Println(err)
		return
	}
	w.Write(file)


}