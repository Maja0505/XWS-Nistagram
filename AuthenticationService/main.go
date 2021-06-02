package main

import (
	"authenticationService/handler"
	"authenticationService/repository"
	"authenticationService/service"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis"
	"log"
	"os"
)

var  client *redis.Client


func initDB() *redis.Client {
	dsn := os.Getenv("REDIS_DSN")
	if len(dsn) == 0 {
		dsn = "localhost:6379"
	}
	client = redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
		Password: "",
		DB: 0,
	})
	return client
}

func initRepo(database *redis.Client) *repository.AuthenticationRepository {
	return &repository.AuthenticationRepository{Database: database}
}

func initService(repo *repository.AuthenticationRepository) *service.AuthenticationService {
	return &service.AuthenticationService{Repository: repo}
}

func initHandler(service *service.AuthenticationService) *handler.AuthenticationHandler{
	return &handler.AuthenticationHandler{Service: service}
}

func handleFunc(handler *handler.AuthenticationHandler) {
	router := gin.Default()
    router.POST("/todo", handler.CreateTodo)
	router.POST("/login",handler.Login)
	log.Fatal(router.Run(":8070"))

}

func main() {
	client := initDB()
	repo := initRepo(client)
	service := initService(repo)
	handler :=initHandler(service)
	handleFunc(handler)
}






