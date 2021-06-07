package service

import (
	"errors"
	"fmt"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"userService/dto"
	"userService/model"
	"userService/repository"
	"userService/mapper"
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

func (service *UserService) CreateRegisteredUser(userForRegistrationDTO *dto.UserForRegistrationDTO) error {

	if userForRegistrationDTO.Password != userForRegistrationDTO.ConfirmedPassword{
		return errors.New("Password and confirmed password are not same!")
	}
	
	existingUser,_ := service.Repo.FindUserByUsername(userForRegistrationDTO.Username)

	if existingUser != nil{
		return errors.New("User with same name aleready exist!")
	}

	userForRegistration := mapper.ConvertUserForRegistrationDTOToRegisteredUser(userForRegistrationDTO)
	err := service.Repo.CreateRegisteredUser(userForRegistration)
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

func (service *UserService) FindUserByUsername(username string) (*model.RegisteredUser,error){
	user,err := service.Repo.FindUserByUsername(username)
	if err != nil{
		return nil, err
	}
	return user, nil
}