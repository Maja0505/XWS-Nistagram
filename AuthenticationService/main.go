package main

import (
	"XWS-Nistagram/AuthenticationService/handler"
	"XWS-Nistagram/AuthenticationService/model"
	"XWS-Nistagram/AuthenticationService/model/authentication"
	"XWS-Nistagram/AuthenticationService/repository"
	"XWS-Nistagram/AuthenticationService/service"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis"
	_ "github.com/lib/pq"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
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

func initRepo(tokenDatabase *redis.Client,userDatabase *gorm.DB) *repository.AuthenticationRepository {
	return &repository.AuthenticationRepository{TokenDatabase: tokenDatabase,UserDatabase:userDatabase}
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

	router.POST("/login", authenticationHandler.Login)
	authorized := router.Group("/")
	authorized.Use(gin.Logger())
	authorized.Use(gin.Recovery())
	authorized.Use(authenticationHandler.TokenAuthMiddleware())
	{
		authorized.POST("/authorize", authorizationHandler.Authorize())
		authorized.POST("/authorizeDemonstration", authorizationHandler.Authorize(), authenticationHandler.CreateTodo)
		authorized.POST("/logout", authenticationHandler.Logout)
		authorized.POST("/refreshToken", authenticationHandler.RefreshToken)
	}
	log.Fatal(router.Run(":8070"))

}



func initPostgreDB() *gorm.DB{
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Shanghai", "localhost", "postgres","root", "authdetailsdb", "5432")
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	db.Migrator().DropTable(&model.User{})
	db.AutoMigrate(&authentication.User{})
	user := authentication.User{ID: 1,Username: "Pera",Password: "pera",Role:"Admin"}
	user1 := authentication.User{ID: 2,Username: "Marko",Password: "marko",Role:"User"}
	user2:= authentication.User{ID: 3,Username: "Dana",Password: "dana",Role:"Agent"}
	db.Create(&user)
	db.Create(&user1)
	db.Create(&user2)

	fmt.Println("Successfully connected!")
	return db
}


func main(){
	redis:=initDB()
	postgres:=initPostgreDB()
	repository:=initRepo(redis,postgres)
	service:=initService(repository)
	authenticationHandler:=initAuthenticationHandler(service)
	authorizationHandler:=initAuthorizationHandler(service)
	handleFunc(authenticationHandler,authorizationHandler)
}









