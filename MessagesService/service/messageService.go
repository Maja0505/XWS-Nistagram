package service

import (
	"XWS-Nistagram/XWS-Nistagram/MessagesService/model"
	"XWS-Nistagram/XWS-Nistagram/MessagesService/repository"
	"fmt"
	"github.com/go-redis/redis/v7"
)


func Connect(rdb *redis.Client, id string) (*model.UserWS, error) {
	fmt.Println("Id od usera koji treba da se upise : ", id)
	u, err := repository.Connect(rdb, id)
	return u,err
}

func  Disconnect(u *model.UserWS) error {
	err := repository.Disconnect(u)
	return err
}

func Subscribe(rdb *redis.Client, channel string,u *model.UserWS) error {
	err := repository.Subscribe(rdb,channel,u)
	return err
}

func Unsubscribe(rdb *redis.Client, channel string,u *model.UserWS) error {
	err := repository.Unsubscribe(rdb,channel,u)
	return err
}

func SendMessage(rdb *redis.Client,message model.Message,u *model.UserWS) error {
	err := repository.SendMessage(rdb,message,u)
	return err
}
