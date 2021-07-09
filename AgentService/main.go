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

	if err != nil {
		panic(err)
	}
	if err := Session.Query("create keyspace  if not exists agentkeyspace with replication = {'class':'SimpleStrategy','replication_factor':1};").Exec(); err != nil {
		fmt.Println("Error while inserting agentkeyspace")
		fmt.Println(err)
	}

	if err := Session.Query("CREATE TABLE if not exists agentkeyspace.campaigns(id timeuuid, userid text, ispost boolean, repeat boolean, start timestamp, end timestamp, repeatfactor int, media list<text>, links list<text>, influencers list<text>, PRIMARY KEY((userid), id)) WITH CLUSTERING ORDER BY (id DESC);").Exec(); err != nil {
		fmt.Println("Error while creating tables!")
		fmt.Println(err)
		fmt.Println(err)	}

	if err := Session.Query("CREATE TABLE if not exists agentkeyspace.campaignrequest(userid text, campaignid timeuuid, PRIMARY KEY((userid), campaignid)) WITH CLUSTERING ORDER BY (campaignid DESC);").Exec(); err != nil {
		fmt.Println("Error while creating tables!")
		fmt.Println(err)
	}

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
	router.HandleFunc("/create-campaign-request", handler.CreateCampaignRequest).Methods("POST")
	router.HandleFunc("/delete-campaign", handler.DeleteCampaign).Methods("POST")
	router.HandleFunc("/add-influencer", handler.AddCampaignInfluencer).Methods("POST")
	router.HandleFunc("/get-campaigns-for-user/{id}", handler.GetCampaignsForUser).Methods("GET")
	router.HandleFunc("/get-campaign-requests/{id}", handler.GetCampaignRequests).Methods("GET")
}


func main(){

	fmt.Println("\n----------------MAIN----------------\n")

	AgentRepo := initAgentRepo(Session)
	AgentService := initAgentService(AgentRepo)
	handler := initHandler(AgentService)

	router := mux.NewRouter().StrictSlash(true)
	handleFunc(handler,router)

	headers := handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization"})
	methods := handlers.AllowedMethods([]string{"GET", "POST", "PUT", "HEAD", "OPTIONS"})
	origins := handlers.AllowedOrigins([]string{"*"})

	fmt.Println("\nServer running...")
	log.Fatal(http.ListenAndServe(":" + os.Getenv("AGENT_SERVICE_PORT"), handlers.CORS(headers, methods, origins)(router)))

}
