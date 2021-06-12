package main

import (
	"XWS-Nistagram/UserFollowersService/handler"
	"XWS-Nistagram/UserFollowersService/repository"
	"XWS-Nistagram/UserFollowersService/service"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"github.com/neo4j/neo4j-go-driver/v4/neo4j"
	"log"
	"net/http"
	"os"
)

func initDB()  neo4j.Session{
	session, _, err := ConnectToDB()
	if err != nil {
		log.Fatalln("Error connecting to Database")
		log.Fatalln(err)
	}
//	defer driver.Close()
//	defer session.Close()
	log.Println("Starting to listen..")
	return session
}

func ConnectToDB() (neo4j.Session, neo4j.Driver, error) {
	// define driver, session and result vars
	var (
		driver  neo4j.Driver
		session neo4j.Session
		err     error
	)
	// initialize driver to connect to localhost with ID and password
	if driver, err = neo4j.NewDriver("neo4j://localhost:7687", neo4j.BasicAuth("neo4j", "danica", "")); err != nil {
		return nil, nil, err
	}
	// Open a new session with write access
	if session, err = driver.Session(neo4j.AccessModeWrite); err != nil {
		return nil, nil, err
	}
	return session, driver, nil
}

func initRepo(session neo4j.Session) *repository.UserFollowersRepository {
	return &repository.UserFollowersRepository{Session: session}
}

func initService(repo *repository.UserFollowersRepository) *service.UserFollowersService {
	return &service.UserFollowersService{Repository:repo}
}

func initHandler(service *service.UserFollowersService) *handler.UserFollowersHandler {
	return &handler.UserFollowersHandler{Service: service}
}

func handleFunctions(handler *handler.UserFollowersHandler){
	router := mux.NewRouter().StrictSlash(true)

	router.HandleFunc("/followUser",handler.FollowUser).Methods("POST")
	router.HandleFunc("/unfollowUser",handler.UnfollowUser).Methods("PUT")
	router.HandleFunc("/acceptFollowRequest",handler.AcceptFollowRequest).Methods("PUT")
	router.HandleFunc("/cancelFollowRequest",handler.CancelFollowRequest).Methods("PUT")
	router.HandleFunc("/setCloseFriend",handler.SetCloseFriend).Methods("PUT")
	router.HandleFunc("/setMuteFriend",handler.SetMuteFriend).Methods("PUT")

	router.HandleFunc("/allFollows/{userId}",handler.GetAllFollowedUsers).Methods("GET")
	router.HandleFunc("/allFollowers/{userId}",handler.GetAllFollowersByUser).Methods("GET")
	router.HandleFunc("/allFollowRequests/{userId}",handler.GetAllFollowRequests).Methods("GET")
	router.HandleFunc("/allCloseFriends/{userId}",handler.GetAllCloseFriends).Methods("GET")
	router.HandleFunc("/allMuteFriends/{userId}",handler.GetAllMuteFriends).Methods("GET")

	router.HandleFunc("/checkFollowing/{userId}/{followedUserId}",handler.CheckFollowing).Methods("GET")


	fmt.Println("Server running on port " + os.Getenv("USER_FOLLOWERS_SERVICE_PORT"))
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", os.Getenv("USER_FOLLOWERS_SERVICE_PORT")),router))
}

func init() {

	err := godotenv.Load(".env")

	if err != nil {
		log.Fatal("Error loading .env file")
	}
}


func main() {
	db:= initDB()
	repo :=initRepo(db)
	service :=initService(repo)
	handler :=initHandler(service)
	handleFunctions(handler)



}

