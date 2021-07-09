package repository

import (
	"XWS-Nistagram/AgentApplication/model"
	"fmt"
	"gorm.io/gorm"
)

type UserRepository struct{
	Database *gorm.DB
}

func (repository *UserRepository) RegisterUser(user *model.User) error {
	err:= repository.Database.Create(&user)
	if err != nil {
		fmt.Println(err)
		return err.Error
	}
	return nil
}

func (repository *UserRepository) UserExists(email string) bool{
	var user model.User
	if err:=repository.Database.First(&user, "email = ?", email).Error; err != gorm.ErrRecordNotFound {
		return true
	}
	return false
}

func (repository *UserRepository) FindById(userID string) *model.User {
	user := &model.User{}
	repository.Database.First(&user, "id = ?", userID)
	return user
}

func (repository *UserRepository) GetByEmail(email string) (model.User,bool){
	var user model.User
	if err:=repository.Database.First(&user, "email = ?", email).Error; err != gorm.ErrRecordNotFound {
		return user,true
	}
	return user,false
}

func (repository *UserRepository) CheckCredentials(email string,password string) (bool,bool,*model.User){
	user,found:=repository.GetByEmail(email)

	if !found{
		return false,false,nil
	}
	if user.Password == password{
		return true,true,&user
	}
	return true,false,nil
}


