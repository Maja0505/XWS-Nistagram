package main

import (

	"context"
	"fmt"
	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"net/http"
	"userService/handler"
	"userService/repository"
	"userService/service"
)

func initDB() *mongo.Client {
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}
	err = client.Ping(context.TODO(), nil)
	if err != nil {

		log.Fatal(err)

	}
	return client
}

func initUserRepo(database *mongo.Client) *repository.UserRepository{
	return &repository.UserRepository{Database: database}
}

func initUserService(userRepo *repository.UserRepository) *service.UserService{
	return &service.UserService{Repo : userRepo}
}

func initHandler(service *service.UserService) *handler.UserHandler{
	return &handler.UserHandler{Service: service}
}


func handleFunc(handler *handler.UserHandler){
	router := mux.NewRouter().StrictSlash(true)

	router.HandleFunc("/api/findAllUsers",handler.FindAll).Methods("GET")
	router.HandleFunc("/api/create",handler.Create).Methods("POST")
	router.HandleFunc("/api/update/{id}",handler.Update).Methods("PUT")


	fmt.Println("server running ")
	log.Fatal(http.ListenAndServe(":8000", router))
}


func main() {

	database := initDB()
	userRepo := initUserRepo(database)
	userService := initUserService(userRepo)
	handler := initHandler(userService)

	handleFunc(handler)

}
