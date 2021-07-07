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
		log.Println(err)
		log.Fatalln("Error connecting to Database")
	}

	log.Println("Starting to listen..")
	return session
}

func ConnectToDB() (neo4j.Session, neo4j.Driver, error) {
	var (
		driver  neo4j.Driver
		session neo4j.Session
		err     error
	)
	if driver, err = neo4j.NewDriver("neo4j://neo4j:7687", neo4j.BasicAuth("neo4j", "nistagram", "")); err != nil {
		return nil, nil, err
	}
	if session = driver.NewSession(neo4j.SessionConfig{AccessMode: neo4j.AccessModeWrite}); err != nil {
		return nil, nil, err
	}

	_,err  = session.WriteTransaction(func(tx neo4j.Transaction) (interface{}, error) {
		result,err := tx.Run("match (n) return n;", map[string]interface{}{})
		if err != nil{
			return nil, err
		}
		if result.Next(){
			return result.Record().Values[0], err
		}

		return nil, result.Err()
	})

	if err != nil{
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
	router.HandleFunc("/allNotMutedFollows/{userId}",handler.GetAllNotMutedFollowedUsers).Methods("GET")
	router.HandleFunc("/allAllFollowsWhomUserIsCloseFriend/{userId}",handler.GetAllFollowsWhomUserIsCloseFriend).Methods("GET")
	router.HandleFunc("/allAllFollowsWhomUserIsNotCloseFriend/{userId}",handler.GetAllFollowsWhomUserIsNotCloseFriend).Methods("GET")

	router.HandleFunc("/allFollowRequests/{userId}",handler.GetAllFollowRequests).Methods("GET")
	router.HandleFunc("/allCloseFriends/{userId}",handler.GetAllCloseFriends).Methods("GET")
	router.HandleFunc("/allMuteFriends/{userId}",handler.GetAllMuteFriends).Methods("GET")

	router.HandleFunc("/checkFollowing/{userId}/{followedUserId}",handler.CheckFollowing).Methods("GET")
	router.HandleFunc("/checkRequested/{userId}/{requestedUserId}",handler.CheckRequested).Methods("GET")
	router.HandleFunc("/checkMuted/{userId}/{mutedUserId}",handler.CheckMuted).Methods("GET")
	router.HandleFunc("/checkClosed/{userId}/{closedUserId}",handler.CheckClosed).Methods("GET")

	router.HandleFunc("/followSuggestions/{userId}",handler.GetFollowSuggestions).Methods("GET")



}

func handleUserBlockFunctions(handler *handler.BlockedUserHandler,router *mux.Router){

	router.HandleFunc("/blockUser",handler.BlockUser).Methods("POST")
	router.HandleFunc("/getAllBlockUsers/{userId}",handler.GetAllBlockedUsers).Methods("GET")
	router.HandleFunc("/checkBlock/{userId}/{blockedUserId}",handler.CheckBlock).Methods("GET")
	router.HandleFunc("/unblockUser",handler.UnblockUser).Methods("PUT")

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



