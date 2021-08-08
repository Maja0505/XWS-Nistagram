package handler

import (
	"fmt"
	"github.com/gorilla/mux"
	"io"
	"io/ioutil"
	"net/http"
	"os"
)

type MediaStorageHandler struct {

}

func (handler *MediaStorageHandler) GetMediaImage(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "image/jpeg")
	vars := mux.Vars(r)
	imagepath := vars["id"]
	if imagepath == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	file,err := ioutil.ReadFile("media-docs/" + imagepath)
	if err != nil{
		w.WriteHeader(http.StatusBadRequest)
		fmt.Println(err)
		return
	}
	w.Write(file)


}

func (handler *MediaStorageHandler) GetProfilePicture(w http.ResponseWriter, r *http.Request) {
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

func (handler *MediaStorageHandler) GetVerificationDoc(w http.ResponseWriter, r *http.Request) {
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

func (handler *MediaStorageHandler) UploadMediaImage(w http.ResponseWriter,r *http.Request){
	vars := mux.Vars(r)
	imagePath := vars["id"]
	formKey := vars["formKey"]
	r.ParseMultipartForm(10 << 20)
	file, _, err := r.FormFile(formKey)
	if err != nil {
		fmt.Println("Error Retrieving the File")
		fmt.Println(err)
		return
	}

	defer file.Close()

	dst, err := os.Create("media-docs/" + imagePath +".jpg")
	defer dst.Close()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if _, err := io.Copy(dst, file); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	fmt.Fprintf(w, "Successfully Uploaded File\n")


}

func (handler *MediaStorageHandler) UploadProfilePicture(w http.ResponseWriter,r *http.Request){
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

func (handler *MediaStorageHandler) UploadVerificationDoc(w http.ResponseWriter,r *http.Request){
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

func (handler *MediaStorageHandler) UploadVideo(w http.ResponseWriter,r *http.Request){
	vars := mux.Vars(r)
	imagePath := vars["id"]
	formKey := vars["formKey"]
	r.ParseMultipartForm(10 << 20)
	file, _, err := r.FormFile(formKey)
	if err != nil {
		fmt.Println("Error Retrieving the File")
		fmt.Println(err)
		return
	}

	defer file.Close()

	dst, err := os.Create("media-docs/" + imagePath +".mp4")
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

func (handler *MediaStorageHandler) GetVideo(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "video/mp4")
	vars := mux.Vars(r)
	imagepath := vars["id"]
	if imagepath == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	file,err := ioutil.ReadFile("media-docs/" + imagepath)
	if err != nil{
		w.WriteHeader(http.StatusBadRequest)
		fmt.Println(err)
		return
	}
	w.Write(file)


}


