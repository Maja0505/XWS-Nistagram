package main

import (
	"XWS-Nistagram/UserFollowersService/handler"
	"XWS-Nistagram/UserFollowersService/model"
	"XWS-Nistagram/UserFollowersService/repository"
	"XWS-Nistagram/UserFollowersService/service"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"github.com/neo4j/neo4j-go-driver/v4/neo4j"
	"github.com/rs/cors"
	"log"
	"net/http"
	"os"
)

func initDatabase() (neo4j.Session, neo4j.Driver, error) {
	var (
		driver  neo4j.Driver
		session neo4j.Session
		err     error
	)

	

	if driver, err = neo4j.NewDriver("bolt://neo4j", neo4j.BasicAuth("neo4j", "nistagram", "")); err != nil {
		return nil, nil, err
	}
	if session = driver.NewSession(neo4j.SessionConfig{AccessMode: neo4j.AccessModeWrite}); err != nil {
		return nil, nil, err
	}

	defer session.Close()

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

func initUserFollowRepo(driver *neo4j.Driver) *repository.UserFollowersRepository {
	return &repository.UserFollowersRepository{Driver: driver}
}

func initUserFollowService(repo *repository.UserFollowersRepository) *service.UserFollowersService {
	return &service.UserFollowersService{Repository:repo}
}

func initUserFollowHandler(service *service.UserFollowersService,service2 *service.BlockedUserService) *handler.UserFollowersHandler {
	return &handler.UserFollowersHandler{Service: service,BlockedUserService: service2}
}

func initUserBlockRepo(driver neo4j.Driver) *repository.BlockedUserRepository {
	return &repository.BlockedUserRepository{Driver: driver}
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

	router.HandleFunc("/relationship/{userId1}/{userId2}",handler.GetRelationshipBetweenUsers).Methods("GET")


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

func SetupCors() *cors.Cors {
	return cors.New(cors.Options{
		AllowedOrigins: []string{"*"}, // All origins, for now
		AllowedMethods: []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders: []string{"*"},
		AllowCredentials: true,
	})
}


func main() {

	_,driver ,err:= initDatabase()
	if err != nil {
		fmt.Println(err)
		return
	}

	userFollowRepo :=initUserFollowRepo(&driver)
	userFollowService :=initUserFollowService(userFollowRepo)


	userBlockRepo :=initUserBlockRepo(driver)
	userBlockService :=initUserBlockService(userBlockRepo)
	userBlockHandler :=initUserBlockHandler(userBlockService)
	userFollowHandler :=initUserFollowHandler(userFollowService,userBlockService)

	go userFollowService.RedisConnection()


	router := mux.NewRouter().StrictSlash(true)

	handleUserFollowFunctions(userFollowHandler,router)
	handleUserBlockFunctions(userBlockHandler,router)


	//userFollowHandler.Test()

	//cacheMemoryOfNeo4j(userFollowHandler,userBlockHandler)

	c := SetupCors()
	http.Handle("/", c.Handler(router))
	fmt.Println("Server running on port " + os.Getenv("USER_FOLLOWERS_SERVICE_PORT"))

	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", os.Getenv("USER_FOLLOWERS_SERVICE_PORT")),c.Handler(router)))

}

func cacheMemoryOfNeo4j(userFollowHandler *handler.UserFollowersHandler,userBlockHandler *handler.BlockedUserHandler) {

	for i := 0; i < 15; i++ {
		userFollowHandler.Service.GetAllFollowersByUser("test")
		userFollowHandler.Service.GetAllFollowsWhomUserIsNotCloseFriend("test")
		userFollowHandler.Service.GetAllFollowsWhomUserIsCloseFriend("test")
		userFollowHandler.Service.FollowSuggestions("test")
		userFollowHandler.Service.GetAllNotMutedFollowedUsersByUser("test")
		userFollowHandler.Service.Repository.CancelFollowRequest("test","test")
		userFollowHandler.Service.GetAllMuteFriends("test")
		userFollowHandler.Service.GetAllCloseFriends("test")
		userFollowHandler.Service.Repository.SetFriendForClose("test","test",true)
		userFollowHandler.Service.Repository.SetFriendForMute("test","test",true)
		userFollowHandler.Service.CheckClosed("test","test")
		userFollowHandler.Service.CheckMuted("test","test")
		userFollowHandler.Service.CheckFollowing("test","test")
		userFollowHandler.Service.CheckRequested("test","test")
		userFollowHandler.Service.GetAllFollowedUsers("test")
		userFollowHandler.Service.GetAllFollowRequests("test")
		userFollowHandler.Service.Repository.AcceptFollowRequest("test","test")
		userFollowHandler.Service.GetAllFollowsWithoutFollowsWhomUserIsCloseFriend("test")
		userFollowHandler.Service.Repository.FollowUser(&model.FollowRelationship{User: "test",FollowedUser: "test",CloseFriend: false,Muted: false})
		userFollowHandler.Service.Repository.UnfollowUser(&model.FollowRelationship{User: "test",FollowedUser: "test",CloseFriend: false,Muted: false})

		userBlockHandler.Service.CheckBlock("test","test")
		userBlockHandler.Service.BlockUser(&model.BlockRelationship{User: "test",BlockedUser: "test"})
		userBlockHandler.Service.GetAllBlockedUser("test")
		userBlockHandler.Service.UnblockUser(&model.BlockRelationship{User: "test",BlockedUser: "test"})
	}

}



