package handler

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"io"
	"io/ioutil"
	"net/http"
	"os"
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
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	err = handler.Service.Create(&vrDTO)
	if err != nil{
		w.WriteHeader(http.StatusExpectationFailed)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

func (handler *VerificationRequestHandler) Update(w http.ResponseWriter,r *http.Request){
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

func (handler *VerificationRequestHandler) UploadImage(w http.ResponseWriter,r *http.Request){
	r.ParseMultipartForm(10 << 20)

	file, h, err := r.FormFile("myFile")
	if err != nil {
		fmt.Println("Error Retrieving the File")
		fmt.Println(err)
		return
	}

	defer file.Close()
	fmt.Printf("Uploaded File: %+v\n", h.Filename)
	fmt.Printf("File Size: %+v\n", h.Size)
	fmt.Printf("MIME Header: %+v\n", h.Header)

	dst, err := os.Create("verification-docs/" + h.Filename)
	defer dst.Close()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Copy the uploaded file to the created file on the filesystem
	if _, err := io.Copy(dst, file); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	fmt.Fprintf(w, "Successfully Uploaded File\n")


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
