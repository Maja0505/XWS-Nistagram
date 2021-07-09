package main

import (
	"XWS-Nistagram/MediaStorageService/handler"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"log"
	"net/http"
	"os"
)

func handleFunc(handler *handler.MediaStorageHandler,router *mux.Router){
	router.HandleFunc("/get-media-image/{id}", handler.GetMediaImage).Methods("GET")
	router.HandleFunc("/get-profile-picture/{id}", handler.GetProfilePicture).Methods("GET")
	router.HandleFunc("/get-verification-doc/{id}", handler.GetVerificationDoc).Methods("GET")
	router.HandleFunc("/get-video/{id}", handler.GetVideo).Methods("GET")
	router.HandleFunc("/upload-media-image/{id}/{formKey}",handler.UploadMediaImage).Methods("POST")
	router.HandleFunc("/upload-profile-picture/{id}",handler.UploadProfilePicture).Methods("POST")
	router.HandleFunc("/upload-verification-doc/{id}",handler.UploadVerificationDoc).Methods("POST")
	router.HandleFunc("/upload-video/{id}/{formKey}",handler.UploadVideo).Methods("POST")

}

func init() {
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatal("Error loading .env file")
	}  

}

func main() {
	router := mux.NewRouter().StrictSlash(true)
	storageHandler := &handler.MediaStorageHandler{}
	handleFunc(storageHandler,router)
	fmt.Println("Media storage service running ")
	log.Fatal(http.ListenAndServe(":" + os.Getenv("MEDIA_STORAGE_SERVICE_PORT"), router))

}
