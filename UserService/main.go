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

	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb://" + os.Getenv("HOST") + ":27017"))
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

func initVerificationRequestRepo(database *mongo.Client) *repository.VerificationRequestRepository{
	return &repository.VerificationRequestRepository{Database: database}
}

func initUserService(userRepo *repository.UserRepository) *service.UserService{
	return &service.UserService{Repo : userRepo}
}

func initVerificationRequestService(verificationRequestRepo *repository.VerificationRequestRepository,userService *service.UserService) *service.VerificationRequestService{
	return &service.VerificationRequestService{Repo : verificationRequestRepo,UserService:userService }
}


func initUserHandler(service *service.UserService) *handler.UserHandler{
	return &handler.UserHandler{Service: service}
}

func initVerificationRequestHandler(service *service.VerificationRequestService) *handler.VerificationRequestHandler{
	return &handler.VerificationRequestHandler{Service: service}
}

func handleUserFunc(handler *handler.UserHandler,router *mux.Router){

	router.HandleFunc("/",handler.FindAll).Methods("GET")
	//router.HandleFunc("/create",handler.Create).Methods("POST")
	router.HandleFunc("/update/{username}",handler.UpdateRegisteredUserProfile).Methods("PUT")
	router.HandleFunc("/user/create",handler.CreateRegisteredUser).Methods("POST")
	//router.HandleFunc("/update/{id}",handler.Update).Methods("PUT")
	router.HandleFunc("/user/{username}",handler.FindUserByUsername).Methods("GET")
	router.HandleFunc("/user/search/{searchContent}",handler.SearchUser).Methods("GET")
	router.HandleFunc("/convert-user-ids",handler.ConvertUserIdsToUsers).Methods("POST")

}

func init() {

	err := godotenv.Load(".env")

	if err != nil {
		log.Fatal("Error loading .env file")
	}
}
func handleVerificationRequestFunc(handler *handler.VerificationRequestHandler,router *mux.Router){

	router.HandleFunc("/verificationRequest/create",handler.Create).Methods("POST")

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
	userHandler := initUserHandler(userService)

	verificationRequestRepo := initVerificationRequestRepo(database)
	verificationRequestService := initVerificationRequestService(verificationRequestRepo,userService)
	verificationRequestHandler := initVerificationRequestHandler(verificationRequestService)


	router := mux.NewRouter().StrictSlash(true)
	handleUserFunc(userHandler,router)
	handleVerificationRequestFunc(verificationRequestHandler,router)

	fmt.Println("Server running on port " + os.Getenv("USER_SERVICE_PORT"))
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", os.Getenv("USER_SERVICE_PORT")),router))
}

