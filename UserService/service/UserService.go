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

func (service *UserService) SearchUser(searchContent string) (*[]dto.UserFromSearchDTO,error){
	users,err := service.Repo.FindAllUsersBySearchingContent(searchContent)
	if err != nil{
		return nil, err
	}
	usersListDTO := mapper.ConvertUsersListTOUserFromSearchDTOList(users)
	return usersListDTO, err
}

func (service *UserService) ConvertUserIdsToUsers(userIds dto.UserIdsDTO) (*[]dto.UserByUsernameDTO, error) {
	users,err := service.Repo.FindUsernameByUserId(userIds)
	if err != nil{
		return nil, err
	}
	return users,nil
}

func (service *UserService) ChangePassword(username string,passwordDto dto.PasswordDTO) (bool,error) {
	if passwordDto.NewPassword != passwordDto.ConfirmNewPassword{
		return true,errors.New("Please make sure both passwords match.")
	}
	if service.Repo.CheckOldPassword(username,passwordDto.OldPassword) == false{
		return true,errors.New("Your old password was entered incorrectly. Please enter it again.")
	}
	err := service.Repo.ChangePassword(username,passwordDto.NewPassword)
	if err != nil{
		return false,err
	}
	return true,nil

}