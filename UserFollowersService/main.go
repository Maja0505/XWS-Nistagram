package main

import (
	"XWS-Nistagram/UserFollowersService/handler"
	"XWS-Nistagram/UserFollowersService/repository"
	"XWS-Nistagram/UserFollowersService/service"
	"github.com/neo4j/neo4j-go-driver/v4/neo4j"
	"log"
	"net/http"
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
	http.HandleFunc("/followUser",handler.FollowUser)
	log.Fatal(http.ListenAndServe(":3000", nil))
}


func main() {
	db:= initDB()
	repo :=initRepo(db)
	service :=initService(repo)
	handler :=initHandler(service)
	handleFunctions(handler)
}

