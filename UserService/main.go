package main

import (
	"context"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"image"
	"image/jpeg"
	"image/png"
	"log"
	"net/http"
	"os"
	"time"
	"userService/handler"
	"userService/repository"
	"userService/service"
)

func initDB() *mongo.Client {

	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb://" + os.Getenv("USER_SERVICE_HOST") + ":27017"))
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
	router.HandleFunc("/update/{username}",handler.UpdateRegisteredUserProfile).Methods("PUT")
	router.HandleFunc("/create",handler.CreateRegisteredUser).Methods("POST")
	router.HandleFunc("/{username}",handler.FindUserByUsername).Methods("GET")
	router.HandleFunc("/search/{searchContent}",handler.SearchUser).Methods("GET")
	router.HandleFunc("/convert-user-ids",handler.ConvertUserIdsToUsers).Methods("POST")
	router.HandleFunc("/change-password/{username}",handler.ChangePassword).Methods("PUT")
	router.HandleFunc("/{username}/public-profile/{setting}",handler.UpdatePublicProfileSetting).Methods("PUT")
	router.HandleFunc("/{username}/message-request/{setting}",handler.UpdateMessageRequestSetting).Methods("PUT")
	router.HandleFunc("/{username}/allow-tags/{setting}",handler.UpdateAllowTagsSetting).Methods("PUT")
	router.HandleFunc("/{username}/like-notification/{setting}",handler.UpdateLikeNotificationSetting).Methods("PUT")
	router.HandleFunc("/{username}/comment-notification/{setting}",handler.UpdateCommentNotificationSetting).Methods("PUT")
	router.HandleFunc("/{username}/message-request-notification/{setting}",handler.UpdateMessageRequestNotificationSetting).Methods("PUT")
	router.HandleFunc("/{username}/message-notification/{setting}",handler.UpdateMessageNotificationSetting).Methods("PUT")
	router.HandleFunc("/{username}/follow-request-notification/{setting}",handler.UpdateFollowRequestNotificationSetting).Methods("PUT")
	router.HandleFunc("/{username}/follow-notification/{setting}",handler.UpdateFollowNotificationSetting).Methods("PUT")

}

func handleVerificationRequestFunc(handler *handler.VerificationRequestHandler,router *mux.Router){

	router.HandleFunc("/verification-request/create",handler.Create).Methods("POST")
	router.HandleFunc("/verification-request/update/{user}",handler.Update).Methods("PUT")
	router.HandleFunc("/verification-request/all",handler.GetAllVerificationRequest).Methods("GET")
	router.HandleFunc("/verification-request/{user}",handler.GetVerificationRequestByUser).Methods("GET")
	router.HandleFunc("/verification-request/approve/{user}",handler.ApproveVerificationRequest).Methods("PUT")
	router.HandleFunc("/verification-request/upload-verification-doc/{id}",handler.UploadImage).Methods("POST")
	router.HandleFunc("/verification-request/get-image/{id}", handler.GetImage).Methods("GET")

}


func init() {

	err := godotenv.Load(".env")

	if err != nil {
		log.Fatal("Error loading .env file")
	}

	image.RegisterFormat("jpeg", "jpeg", jpeg.Decode, jpeg.DecodeConfig)
	image.RegisterFormat("png", "png", png.Decode, png.DecodeConfig)

}

func main() {
	database := initDB()

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

