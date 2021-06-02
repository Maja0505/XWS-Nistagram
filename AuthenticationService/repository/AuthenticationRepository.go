package repository

import (
	"authenticationService/model"
	"github.com/go-redis/redis"
	"strconv"
	"time"
)

type AuthenticationRepository struct{
	Database *redis.Client
}

func (repository *AuthenticationRepository) FetchAuth(authD *model.AccessDetails) (uint64, error) {
	userid, err := repository.Database.Get(authD.AccessUuid).Result()
	if err != nil {
		return 0, err
	}
	userID, _ := strconv.ParseUint(userid, 10, 64)
	return userID, nil
}

func (repository *AuthenticationRepository)  CreateAuth(userid uint64, td *model.TokenDetails) error {
	at := time.Unix(td.AtExpires, 0) //converting Unix to UTC(to Time object)
	rt := time.Unix(td.RtExpires, 0)
	now := time.Now()
	errAccess := repository.Database.Set(td.AccessUuid, strconv.Itoa(int(userid)), at.Sub(now)).Err()
	if errAccess != nil {
		return errAccess
	}
	errRefresh := repository.Database.Set(td.RefreshUuid, strconv.Itoa(int(userid)), rt.Sub(now)).Err()
	if errRefresh != nil {
		return errRefresh
	}
	return nil
}
