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

func initUserFollowRepo(session neo4j.Session) *repository.UserFollowersRepository {
	return &repository.UserFollowersRepository{Session: session}
}

func initUserFollowService(repo *repository.UserFollowersRepository) *service.UserFollowersService {
	return &service.UserFollowersService{Repository:repo}
}

func initUserFollowHandler(service *service.UserFollowersService) *handler.UserFollowersHandler {
	return &handler.UserFollowersHandler{Service: service}
}

func initUserBlockRepo(session neo4j.Session) *repository.BlockedUserRepository {
	return &repository.BlockedUserRepository{Session: session}
}

func initUserBlockService(repo *repository.BlockedUserRepository) *service.BlockedUserService {
	return &service.BlockedUserService{Repository:repo}
}

func initUserBlockHandler(service *service.BlockedUserService) *handler.BlockedUserHandler {
	return &handler.BlockedUserHandler{Service: service}
}

func handleUserFollowFunctions(handler *handler.UserFollowersHandler,router *mux.Router){

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



}

func handleUserBlockFunctions(handler *handler.BlockedUserHandler,router *mux.Router){

	router.HandleFunc("/blockUser",handler.BlockUser).Methods("POST")
	router.HandleFunc("/getAllBlockUsers/{userId}",handler.GetAllBlockedUsers).Methods("GET")
	router.HandleFunc("/checkBlock/{userId}/{blockedUserId}",handler.CheckBlock).Methods("GET")

}

func init() {

	err := godotenv.Load(".env")

	if err != nil {
		log.Fatal("Error loading .env file")
	}
}


func main() {
	db:= initDB()

	userFollowRepo :=initUserFollowRepo(db)
	userFollowService :=initUserFollowService(userFollowRepo)
	userFollowHandler :=initUserFollowHandler(userFollowService)

	userBlockRepo :=initUserBlockRepo(db)
	userBlockService :=initUserBlockService(userBlockRepo)
	userBlockHandler :=initUserBlockHandler(userBlockService)


	router := mux.NewRouter().StrictSlash(true)

	handleUserFollowFunctions(userFollowHandler,router)
	handleUserBlockFunctions(userBlockHandler,router)


	fmt.Println("Server running on port " + os.Getenv("USER_FOLLOWERS_SERVICE_PORT"))
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", os.Getenv("USER_FOLLOWERS_SERVICE_PORT")),router))

}



