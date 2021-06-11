package service

import (
	"XWS-Nistagram/AuthenticationService/model/authentication"
	"XWS-Nistagram/AuthenticationService/repository"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/twinj/uuid"
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
	//Creating Access Token
	os.Setenv("ACCESS_SECRET", "mcmvmkmsdnfsdmfdsjf") //this should be in an env file
	atClaims := jwt.MapClaims{}
	atClaims["authorized"] = true
	atClaims["access_uuid"] = td.AccessUuid
	atClaims["user_id"] = userid
	atClaims["exp"] = td.AtExpires
	atClaims["role"] = role
	at := jwt.NewWithClaims(jwt.SigningMethodHS256, atClaims)
	td.AccessToken, err = at.SignedString([]byte(os.Getenv("ACCESS_SECRET")))
	if err != nil {
		return nil, err
	}
	fmt.Println("Uspesno kreiran access token!")
	//Creating Refresh Token
	os.Setenv("REFRESH_SECRET", "mcmvmkmsdnfsdmfdsjf") //this should be in an env file
	rtClaims := jwt.MapClaims{}
	rtClaims["refresh_uuid"] = td.RefreshUuid
	rtClaims["user_id"] = userid
	rtClaims["exp"] = td.RtExpires
	atClaims["role"] = role
	rt := jwt.NewWithClaims(jwt.SigningMethodHS256, rtClaims)
	td.RefreshToken, err = rt.SignedString([]byte(os.Getenv("REFRESH_SECRET")))
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
		return []byte(os.Getenv("ACCESS_SECRET")), nil
	})
	if err != nil {
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

func (service *AuthenticationService) CreateAuth(id uint64, ts *authentication.TokenDetails) error {
	return service.Repository.CreateAuth(id,ts)
}

func (service *AuthenticationService) DeleteAuth(accessUuid string) (int64,error) {
	return service.Repository.DeleteAuth(accessUuid)
}



