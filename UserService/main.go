package main

import (
	"context"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"net/http"
	"os"
	"time"
	"userService/handler"
	"userService/repository"
	"userService/service"
)

func initDB() *mongo.Client {

	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb://mongo:27017"))
	if err != nil {
		log.Fatal(err)
	}
	ctx, _ := context.WithTimeout(context.Background(), 30*time.Second)
	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}
	err = client.Ping(ctx, nil)
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

	router.HandleFunc("/",handler.FindAll).Methods("GET")
	router.HandleFunc("/create",handler.Create).Methods("POST")
	router.HandleFunc("/update/{id}",handler.Update).Methods("PUT")


	fmt.Println("server running ")
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", os.Getenv("PORT")), router))



}

func init() {

	err := godotenv.Load(".env")

	if err != nil {
		log.Fatal("Error loading .env file")
	}
}


func main() {

	database := initDB()
	//fmt.Println(d.Collection("users").Find(context.TODO(),bson.M{}))

	userRepo := initUserRepo(database)
	userService := initUserService(userRepo)
	handler := initHandler(userService)
	handleFunc(handler)


}
