package main

import (
	"XWS-Nistagram/AgentService/Handler"
	"XWS-Nistagram/AgentService/Repository"
	"XWS-Nistagram/AgentService/Service"
	"fmt"
	"github.com/gocql/gocql"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"log"
	"net/http"
	"os"
)

var Session *gocql.Session

func init() {
	fmt.Println("Initialization of cassandra...")
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatal("Error loading .env file")
	}
	cluster := gocql.NewCluster(os.Getenv("AGENT_SERVICE_HOST"))
	//cluster.ProtoVersion = 4
	//cluster.Keyspace = "postkeyspace"

	Session, err = cluster.CreateSession()

	fmt.Println("Cassandra well initialized!")
}


func initAgentRepo(session *gocql.Session) *Repository.AgentRepository{
	return &Repository.AgentRepository{Session: *session}
}

func initAgentService(AgentRepo *Repository.AgentRepository) *Service.AgentService{
	return &Service.AgentService{Repo : *AgentRepo}
}

func initHandler(service *Service.AgentService) *Handler.AgentHandler{
	return &Handler.AgentHandler{Service: service}
}


func handleFunc(handler *Handler.AgentHandler,router *mux.Router){
	router.HandleFunc("/create-campaign", handler.CreateCampaign).Methods("POST")
	router.HandleFunc("/delete-campaign", handler.DeleteCampaign).Methods("POST")

}


func main(){

	fmt.Println("\n----------------MAIN----------------\n")

	AgentRepo := initAgentRepo(Session)
	AgentService := initAgentService(AgentRepo)
	handler := initHandler(AgentService)

	//AgentRepo.CreateTables()


	router := mux.NewRouter().StrictSlash(true)
	handleFunc(handler,router)

	headers := handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization"})
	methods := handlers.AllowedMethods([]string{"GET", "POST", "PUT", "HEAD", "OPTIONS"})
	origins := handlers.AllowedOrigins([]string{"*"})

	fmt.Println("\nServer running...")
	log.Fatal(http.ListenAndServe(":" + os.Getenv("AGENT_SERVICE_PORT"), handlers.CORS(headers, methods, origins)(router)))

}
