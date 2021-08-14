package main

import (
	"XWS-Nistagram/XWS-Nistagram/MessagesService/handler"
	"XWS-Nistagram/XWS-Nistagram/MessagesService/model"
	"fmt"
	"github.com/go-redis/redis/v7"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"log"
	"net/http"
)

var rdb *redis.Client

func init(){
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatal("Error loading .env file")
	}

	rdb = redis.NewClient(&redis.Options{Addr: "redis:6379"})
	pong, err := rdb.Ping().Result()
	if err != nil {
		fmt.Println("Error while connecting to db")
	}
	fmt.Println(pong, err)
	fmt.Println(rdb)
	rdb.SAdd(model.ChannelsKey, "general")

	defer rdb.Close()
}

func setupRoutes() {
	fmt.Println("rdb:")

	fmt.Println(rdb)

	r := mux.NewRouter()
	r.Path("/connect-ws/{id}").Methods("GET").HandlerFunc(handler.H(rdb, handler.ServeWS))
	r.Path("/user/{id}/all-chats").Methods("GET").HandlerFunc(handler.H(rdb, handler.GetAllMessageChatForUser))
	r.Path("/channels/{userid1}/{userid2}/view-messages").Methods("GET").HandlerFunc(handler.H(rdb, handler.ViewMessagesInChat))
	r.Path("/channels/{channel}/view-notifications").Methods("GET").HandlerFunc(handler.H(rdb, handler.ViewAllNotificationsInChannel))


	port := ":8080"

	headers := handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization"})
	methods := handlers.AllowedMethods([]string{"GET", "POST", "PUT", "HEAD", "OPTIONS"})
	origins := handlers.AllowedOrigins([]string{"*"})
	fmt.Println("notification service started on port", port)
	log.Fatal(http.ListenAndServe(port,  handlers.CORS(headers, methods, origins)(r)))
}

func main() {
	fmt.Println("Distributed Chat App v0.01")
	setupRoutes()
	http.ListenAndServe(":8080", nil)

}
