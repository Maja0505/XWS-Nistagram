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
	router.HandleFunc("/add-influencer", handler.AddCampaignInfluencer).Methods("POST")

}


func main(){

	fmt.Println("\n----------------MAIN----------------\n")

	AgentRepo := initAgentRepo(Session)
	AgentService := initAgentService(AgentRepo)
	handler := initHandler(AgentService)

	//AgentRepo.CreateTables()
	//uuid, err := ParseUUID("df397943-e018-11eb-80d7-d43d7e26656f")
	//fmt.Println(err)
	//a, err := AgentRepo.GetCampaignInfluencers(uuid)
	//fmt.Println(a)


	router := mux.NewRouter().StrictSlash(true)
	handleFunc(handler,router)

	headers := handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization"})
	methods := handlers.AllowedMethods([]string{"GET", "POST", "PUT", "HEAD", "OPTIONS"})
	origins := handlers.AllowedOrigins([]string{"*"})

	fmt.Println("\nServer running...")
	log.Fatal(http.ListenAndServe(":" + os.Getenv("AGENT_SERVICE_PORT"), handlers.CORS(headers, methods, origins)(router)))

}

func ParseUUID(input string) (gocql.UUID, error) {
	var u gocql.UUID
	j := 0
	for _, r := range input {
		switch {
		case r == '-' && j&1 == 0:
			continue
		case r >= '0' && r <= '9' && j < 32:
			u[j/2] |= byte(r-'0') << uint(4-j&1*4)
		case r >= 'a' && r <= 'f' && j < 32:
			u[j/2] |= byte(r-'a'+10) << uint(4-j&1*4)
		case r >= 'A' && r <= 'F' && j < 32:
			u[j/2] |= byte(r-'A'+10) << uint(4-j&1*4)
		default:
			return gocql.UUID{}, fmt.Errorf("invalid UUID %q", input)
		}
		j += 1
	}
	if j != 32 {
		return gocql.UUID{}, fmt.Errorf("invalid UUID %q", input)
	}
	return u, nil
}
