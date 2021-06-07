package service

import (
	"errors"
	"fmt"
	"userService/dto"
	"userService/mapper"
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

func (service *UserService) UpdateRegisteredUserProfile(username string, registeredUserDto *dto.RegisteredUserProfileInfoDTO) error {
	if username != registeredUserDto.Username{
		existedUser,_ := service.FindUserByUsername(registeredUserDto.Username)
		if existedUser != nil{
			return errors.New("Username already exist")
		}
	}
	registeredUser := mapper.ConvertRegisteredUserDtoToRegisteredUser(registeredUserDto)
	err := service.Repo.UpdateRegisteredUserProfile(username, registeredUser)
	if err != nil{
		fmt.Println(err)
		return err
	}
	return nil
}

func (service *UserService) FindUserByUsername(username string) (*model.RegisteredUser,error){
	user,err := service.Repo.FindUserByUsername(username)
	if err != nil{
		return nil, err
	}
	return user, nil
}