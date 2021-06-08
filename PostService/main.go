package main

import (
	"XWS-Nistagram/PostService/Handler"
	"XWS-Nistagram/PostService/Repository"
	"XWS-Nistagram/PostService/Service"
	"fmt"
	"github.com/gocql/gocql"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

var Session *gocql.Session

func init() {
	var err error
	cluster := gocql.NewCluster("127.0.0.1")
	cluster.ProtoVersion = 4
	cluster.Keyspace = "postkeyspace"

	Session, err = cluster.CreateSession()
	if err != nil {
		panic(err)
	}
	fmt.Println("Cassandra well initialized!")
}

func initPostRepo(session *gocql.Session) *Repository.PostRepository{
	return &Repository.PostRepository{Session: *session}
}

func initPostService(postRepo *Repository.PostRepository) *Service.PostService{
	return &Service.PostService{Repo : *postRepo}
}

func initHandler(service *Service.PostService) *Handler.PostHandler{
	return &Handler.PostHandler{Service: service}
}

func handleFunc(handler *Handler.PostHandler){

	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/create", handler.Create).Methods("POST")
	router.HandleFunc("/add-comment", handler.AddComment).Methods("POST")
	router.HandleFunc("/delete-comment", handler.DeleteComment).Methods("POST")
	router.HandleFunc("/like-post", handler.LikePost).Methods("POST")
	router.HandleFunc("/dislike-post", handler.DislikePost).Methods("POST")
	router.HandleFunc("/get-one-post/{id}", handler.FindPostById).Methods("GET")
	router.HandleFunc("/get-all-by-userid/{id}", handler.FindPostsByUserId).Methods("GET")
	router.HandleFunc("/get-comments-for-post/{id}", handler.GetCommentsForPost).Methods("GET")
	router.HandleFunc("/get-users-who-liked-post/{id}", handler.GetUsersWhoLikedPost).Methods("GET")
	router.HandleFunc("/get-users-who-disliked-post/{id}", handler.GetUsersWhoDislikedPost).Methods("GET")
	router.HandleFunc("/get-image/{id}", handler.GetImage).Methods("GET")
	
	headers := handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization"})
	methods := handlers.AllowedMethods([]string{"GET", "POST", "PUT", "HEAD", "OPTIONS"})
	origins := handlers.AllowedOrigins([]string{"*"})

	fmt.Println("server running ")
	log.Fatal(http.ListenAndServe(":8000", handlers.CORS(headers, methods, origins)(router)))
	//log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", os.Getenv("POST_SERVICE_PORT")),router))
	//log.Fatal(http.ListenAndServe(":8000", router))
}

func main(){
	fmt.Println("Main")
	postRepo := initPostRepo(Session)
	postService := initPostService(postRepo)
	handler := initHandler(postService)

	postRepo.CreateTables()
	handleFunc(handler)
}