package main

import (
	"XWS-Nistagram/XWS-Nistagram/MessageService/Handler"

	"fmt"
	"github.com/go-redis/redis/v7"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"log"
	"net/http"
	"os"
)

var rdb *redis.Client


func init(){
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatal("Error loading .env file")
	}

	rdb := redis.NewClient(&redis.Options{Addr: "redis:6379"})
	defer rdb.Close()
	//rdb.SAdd(Model.ChannelsKey, "general", "random")
}


func main() {
	fmt.Println("aaaaaaaaaa")

	rdb = redis.NewClient(&redis.Options{Addr: "redis:6379"})
	pong, err := rdb.Ping().Result()
	if err != nil {
		fmt.Println("Greskaaa")
	}
	fmt.Println(pong, err)
	// Output: PONG <nil>


	r := mux.NewRouter()
	fmt.Println(rdb)
	r.Path("/chat/{username}").Methods("GET").HandlerFunc(Handler.H(rdb, Handler.MessageWebSocketHandler))
	r.Path("/user/{user}/channels").Methods("GET").HandlerFunc(Handler.H(rdb, Handler.UserChannelsHandler))
	r.Path("/channels/{channel}").Methods("GET").HandlerFunc(Handler.H(rdb, Handler.UserChannelsNotificationsHandler))
	r.Path("/channels/{channel}/not-opened").Methods("GET").HandlerFunc(Handler.H(rdb, Handler.UserChannelsNotOpenedNotificationsHandler))
	r.Path("/user/{user}/channels/{channel}/update").Methods("PUT").HandlerFunc(Handler.H(rdb, Handler.UserChannelsNotificationsUpdateHandler))


	r.Path("/users").Methods("GET").HandlerFunc(Handler.H(rdb, Handler.UsersHandler))

	port := ":" + os.Getenv("PORT")
	if port == ":" {
		port = ":8080"
	}
	fmt.Println("notification service started on port", port)
	log.Fatal(http.ListenAndServe(port, r))
}