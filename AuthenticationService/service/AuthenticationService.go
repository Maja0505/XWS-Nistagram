package service

import (
	"XWS-Nistagram/AuthenticationService/model/authentication"
	"XWS-Nistagram/AuthenticationService/repository"
	"XWS-Nistagram/AuthenticationService/saga"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/go-redis/redis"
	"github.com/twinj/uuid"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"
)

type AuthenticationService struct{
	Repository *repository.AuthenticationRepository
}

func (service *AuthenticationService)  CreateToken(userid uint64,role string) (*authentication.TokenDetails, error) {
	td := &authentication.TokenDetails{}
	td.AtExpires = time.Now().Add(time.Minute * 15).Unix()
	td.AccessUuid = uuid.NewV4().String()

	td.RtExpires = time.Now().Add(time.Hour * 24 * 7).Unix()
	td.RefreshUuid = uuid.NewV4().String()

	var err error

	os.Setenv("ACCESS_SECRET", "") //this should be in an env file
	atClaims := jwt.MapClaims{}
	atClaims["authorized"] = true
	atClaims["access_uuid"] = td.AccessUuid
	atClaims["user_id"] = userid
	atClaims["exp"] = td.AtExpires
	atClaims["role"] = role
	at := jwt.NewWithClaims(jwt.SigningMethodHS256, atClaims)
	fmt.Println(os.Getenv("ACCESS_SECRET"))
	accessSecret,_:=base64.URLEncoding.DecodeString(os.Getenv("ACCESS_SECRET"))

	td.AccessToken, err = at.SignedString(accessSecret)
	if err != nil {
		return nil, err
	}
	fmt.Println("Uspesno kreiran access token!")
	//Creating Refresh Token
	os.Setenv("REFRESH_SECRET", "") //this should be in an env file
	rtClaims := jwt.MapClaims{}
	rtClaims["refresh_uuid"] = td.RefreshUuid
	rtClaims["user_id"] = userid
	rtClaims["exp"] = td.RtExpires
	atClaims["role"] = role
	rt := jwt.NewWithClaims(jwt.SigningMethodHS256, rtClaims)
	refreshSecret,_:=base64.URLEncoding.DecodeString(os.Getenv("REFRESH_SECRET"))
	td.RefreshToken, err = rt.SignedString(refreshSecret)
	if err != nil {
		return nil, err
	}
	fmt.Println("Uspesno kreiran refresh token!")
	return td, nil
}


func (service *AuthenticationService)  ExtractToken(r *http.Request) string {
	bearToken := r.Header.Get("Authorization")
	//normally Authorization the_token_xxx
	strArr := strings.Split(bearToken, " ")
	if len(strArr) == 2 {
		return strArr[1]
	}
	return ""
}

func (service *AuthenticationService) VerifyToken(r *http.Request) (*jwt.Token, error) {
	tokenString := service.ExtractToken(r)
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		//Make sure that the token method conform to "SigningMethodHMAC"
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		secret,_:=base64.URLEncoding.DecodeString(os.Getenv("ACCESS_SECRET"))
		return secret, nil
	})
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	return token, nil
}

func (service *AuthenticationService) TokenValid(r *http.Request) error {
	token, err := service.VerifyToken(r)
	if err != nil {
		return err
	}
	if _, ok := token.Claims.(jwt.Claims); !ok && !token.Valid {
		return err
	}
	return nil
}

func (service *AuthenticationService) ExtractTokenMetadata(r *http.Request) (*authentication.AccessDetails, error) {
	token, err := service.VerifyToken(r)
	if err != nil {
		return nil, err
	}
	claims, ok := token.Claims.(jwt.MapClaims)
	if ok && token.Valid {
		accessUuid, ok := claims["access_uuid"].(string)
		if !ok {
			return nil, err
		}
		userId, err := strconv.ParseUint(fmt.Sprintf("%.f", claims["user_id"]), 10, 64)
		if err != nil {
			return nil, err
		}
		role, ok := claims["role"].(string)
		if !ok {
			return nil, err
		}
		return &authentication.AccessDetails{
			AccessUuid: accessUuid,
			UserId:   userId,
			Role : role,
		}, nil
	}

	return nil, err
}

func (service *AuthenticationService) CreateAuth(id uint64, tokenDetails *authentication.TokenDetails) error {
	return service.Repository.CreateAuth(id,tokenDetails)
}

func (service *AuthenticationService) DeleteAuth(accessUuid string) (int64,error) {
	return service.Repository.DeleteAuth(accessUuid)
}
func (service *AuthenticationService)  CheckCredentials(username string,password string) (bool,bool) {
	return service.Repository.CheckCredentials(username,password)
}

func (service *AuthenticationService)  GetByUsername(username string) (authentication.User) {
	user,_:=service.Repository.GetByUsername(username)
	return user
}


//saga deo


func (service *AuthenticationService) RedisConnection() {
	// create client and ping redis
	var err error
	client := redis.NewClient(&redis.Options{Addr: "redis:6379"})
	if _, err = client.Ping().Result(); err != nil {
		log.Fatalf("error creating redis client %s", err)
	}

	// subscribe to the required channels
	pubsub := client.Subscribe(saga.AuthenticationChannel, saga.ReplyChannel)
	if _, err = pubsub.Receive(); err != nil {
		log.Fatalf("error subscribing %s", err)
	}
	defer func() { _ = pubsub.Close() }()
	ch := pubsub.Channel()

	log.Println("starting the authentication-service service")
	for {
		select {
		case msg := <-ch:
			m := saga.Message{}
			err := json.Unmarshal([]byte(msg.Payload), &m)
			if err != nil {
				log.Println(err)
				continue
			}

			switch msg.Channel {
			case saga.AuthenticationChannel:

				// Happy Flow
				if m.Action == saga.ActionStart {

					fmt.Println("Stigao zahtev sa user service  !")
					var user authentication.User
					user.Username = m.Username
					user.Password = m.Password
					user.Role = m.Role
					fmt.Println("Upisuje se user ",user.Username, "  ",user.Password,  " ",user.Role)
					_,err := service.Repository.CreateUser(&user)
					if err != nil{
						fmt.Println("Neuspesno upisao u bazu authentication")
						sendToReplyChannel(client, &m, saga.ActionRollback, saga.ServiceUser, saga.ServiceAuthentication)
					}else{
						fmt.Println("Uspesno upisao u bazu authentication")
						sendToReplyChannel(client, &m, saga.ActionDone, saga.ServiceUserFollower, saga.ServiceAuthentication)
					}

				}

				// Rollback flow
				if m.Action == saga.ActionRollback {
					err := service.Repository.DeleteUser(m.Username)
					if err != nil{
						fmt.Println("Neuspesan rolback err : ", err)
					}
					sendToReplyChannel(client, &m, saga.ActionRollback, saga.ServiceUser, saga.ServiceAuthentication)
				}

			}
		}
	}
}

func sendToReplyChannel(client *redis.Client, m *saga.Message, action string, service string, senderService string) {
	var err error
	m.Action = action
	m.Service = service
	m.SenderService = senderService
	if err = client.Publish(saga.ReplyChannel, m).Err(); err != nil {
		log.Printf("error publishing done-message to %s channel", saga.ReplyChannel)
	}
	log.Printf("done message published to channel :%s", saga.ReplyChannel)
}



