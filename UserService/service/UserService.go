package service

import (
	"fmt"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"userService/model"
	"userService/repository"
)

type UserService struct {
	Repo *repository.UserRepository
}

func (service *UserService) FindAll() (*[]model.User, error) {
	users,err := service.Repo.FindAll()
	if err != nil {
		return nil,err
	}
	return users,nil
}

func (service *UserService) Create(user *model.User) error {
	err := service.Repo.Create(user)
	if err != nil{
		fmt.Println(err)
		return  err
	}
	return nil
}

func (service *UserService) Update(stringId string, user *model.User) error {
	id,err := primitive.ObjectIDFromHex(stringId)
	if err != nil{
		fmt.Println(err)
		return err
	}
	err = service.Repo.Update(id,user)
	if err != nil{
		fmt.Println(err)
		return err
	}
	return nil
}

