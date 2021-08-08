package repository

import (
	"XWS-Nistagram/AuthenticationService/model/authentication"
	"fmt"
	"github.com/go-redis/redis"
	"gorm.io/gorm"
	"strconv"
	"time"
)

type AuthenticationRepository struct{
	TokenDatabase *redis.Client
	UserDatabase *gorm.DB
}

func (repository *AuthenticationRepository) FetchAuth(authD *authentication.AccessDetails) (uint64, error) {
	userid, err := repository.TokenDatabase.Get(authD.AccessUuid).Result()
	if err != nil {
		return 0, err
	}
	userID, _ := strconv.ParseUint(userid, 10, 64)
	return userID, nil
}

func (repository *AuthenticationRepository) GetByUsername(username string) (authentication.User,bool){
	var user authentication.User
	if err:=repository.UserDatabase.First(&user, "username = ?", username).Error; err != gorm.ErrRecordNotFound {
		return user,true
	}
	return user,false
}

func (repository *AuthenticationRepository) CheckCredentials(username string,password string) (bool,bool){
	user,found:=repository.GetByUsername(username)

	if !found{
		return false,false
	}
	if user.Password == password{
		return true,true
	}
	return true,false
}

func (repository *AuthenticationRepository)  CreateAuth(userid uint64, td *authentication.TokenDetails) error {
	at := time.Unix(td.AtExpires, 0) //converting Unix to UTC(to Time object)
	rt := time.Unix(td.RtExpires, 0)
	now := time.Now()
	errAccess := repository.TokenDatabase.Set(td.AccessUuid, strconv.Itoa(int(userid)), at.Sub(now)).Err()
	if errAccess != nil {
		return errAccess
	}
	errRefresh := repository.TokenDatabase.Set(td.RefreshUuid, strconv.Itoa(int(userid)), rt.Sub(now)).Err()
	if errRefresh != nil {
		return errRefresh
	}
	return nil
}

func (repository *AuthenticationRepository)  DeleteAuth(givenUuid string) (int64,error) {
	deleted, err := repository.TokenDatabase.Del(givenUuid).Result()
	if err != nil {
		return 0, err
	}
	return deleted, nil
}

func (repository *AuthenticationRepository) CreateUser(user *authentication.User) (*authentication.User,error) {
	result := repository.UserDatabase.Create(&user)
	fmt.Println(result.RowsAffected)
	if result.RowsAffected == 0 {
		return nil,fmt.Errorf("ticket not created")
	}
	fmt.Println("Shopping cart Created")
	return user,nil
}

func (repository *AuthenticationRepository) FindUserById(username string) (*authentication.User,error){
	var user authentication.User
	if err:=repository.UserDatabase.First(&user, "username = ?", username).Error; err != nil {
		return nil,err
	}
	return &user,nil
}

func (repository *AuthenticationRepository) DeleteUser(username string) error {
	user,err:=repository.FindUserById(username)
	if err!=nil{
		fmt.Println(err)
		return err
	}
	repository.UserDatabase.Delete(&user)
	return nil
}
