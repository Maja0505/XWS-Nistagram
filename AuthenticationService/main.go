package main

import (
	"XWS-Nistagram/AuthenticationService/handler"
	"XWS-Nistagram/AuthenticationService/repository"
	"XWS-Nistagram/AuthenticationService/service"
	fileadapter "github.com/casbin/casbin/persist/file-adapter"
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

func initAuthenticationHandler(service *service.AuthenticationService) *handler.AuthenticationHandler{
	return &handler.AuthenticationHandler{Service: service}
}
func initAuthorizationHandler(authenticatonService *service.AuthenticationService) *handler.AuthorizationHandler{
	return &handler.AuthorizationHandler{AuthenticationService:authenticatonService}
}

func handleFunc(authenticationHandler *handler.AuthenticationHandler,authorizationHandler *handler.AuthorizationHandler) {
	router := gin.Default()
	fileAdapter := fileadapter.NewAdapter("C:\\Users\\danic\\GOprojects\\src\\XWS-Nistagram\\AuthenticationService\\model\\authorization\\policy.csv")
	router.POST("/login", authenticationHandler.Login)
	authorized := router.Group("/")
	authorized.Use(gin.Logger())
	authorized.Use(gin.Recovery())
	authorized.Use(authenticationHandler.TokenAuthMiddleware())
	{
		authorized.POST("/todo",  authorizationHandler.Authorize("resource", "write", fileAdapter), authenticationHandler.CreateTodo)
		authorized.POST("/logout", authenticationHandler.Logout)
		authorized.POST("/refreshToken", authenticationHandler.RefreshToken)
	}
	log.Fatal(router.Run(":8070"))

}

func main(){
	redis:=initDB()
	repository:=initRepo(redis)
	service:=initService(repository)
	authenticationHandler:=initAuthenticationHandler(service)
	authorizationHandler:=initAuthorizationHandler(service)
	handleFunc(authenticationHandler,authorizationHandler)
}









