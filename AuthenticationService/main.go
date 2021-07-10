package main

import (
	"XWS-Nistagram/AuthenticationService/handler"
	"XWS-Nistagram/AuthenticationService/model/authentication"
	"XWS-Nistagram/AuthenticationService/repository"
	"XWS-Nistagram/AuthenticationService/service"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis"
	"github.com/joho/godotenv"
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
		dsn = "redis:6379"
	}
	client = redis.NewClient(&redis.Options{
		Addr: "redis:6379",
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
		authorized.POST("/authorize",  authorizationHandler.Authorize())
		authorized.POST("/authorizeDemonstration", authorizationHandler.Authorize(), authenticationHandler.CreateTodo)
		authorized.POST("/logout", authenticationHandler.Logout)
		authorized.POST("/refreshToken", authenticationHandler.RefreshToken)
	}
	log.Fatal(router.Run(":" + os.Getenv("AUTHENTICATION_SERVICE_PORT")))

}

func init() {

	err := godotenv.Load(".env")

	if err != nil {
		log.Fatal("Error loading .env file")
	}

}



func initPostgreDB() *gorm.DB{
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s", "postgres", "postgres","root", "authdetailsdb", "5432")
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	//db.Migrator().DropTable(&model.User{})
	db.AutoMigrate(&authentication.User{})
	/*user := authentication.User{Username: "Pera",Password: "pera",Role:"Admin"}
	user1 := authentication.User{Username: "Marko",Password: "marko",Role:"User"}
	user2:= authentication.User{Username: "Dana",Password: "dana",Role:"Agent"}
	db.Create(&user)
	db.Create(&user1)
	db.Create(&user2)*/

	fmt.Println("Successfully connected!")
	return db
}


func main(){
	redis:=initDB()
	postgres:=initPostgreDB()
	repository:=initRepo(redis,postgres)
	service:=initService(repository)

	go service.RedisConnection()

	authenticationHandler:=initAuthenticationHandler(service)
	authorizationHandler:=initAuthorizationHandler(service)
	handleFunc(authenticationHandler,authorizationHandler)
}









