package service

import (
	"XWS-Nistagram/AgentApplication/dto"
	"XWS-Nistagram/AgentApplication/mapper"
	"XWS-Nistagram/AgentApplication/model"
	"XWS-Nistagram/AgentApplication/repository"
	"errors"
	"fmt"
)

type UserService struct{
	Repository *repository.UserRepository
}

func (service *UserService) RegisterUser(dto dto.UserDTO) error {

	if dto.Password != dto.ConfirmedPassword{
		return errors.New("Passwords doesn't match!")
	}
	if service.Repository.UserExists(dto.Email){
		return errors.New("User with same name already exist!")
	}
	dto.Role="User"
	err := service.Repository.RegisterUser(mapper.MapDTOToUser(&dto))
	if err != nil{
		fmt.Println(err)
		return  err
	}
	return nil
}

func (service *UserService)  CheckCredentials(email string,password string) (bool,bool,*model.User) {
	return service.Repository.CheckCredentials(email,password)
}




