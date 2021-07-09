package service

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/go-redis/redis"
	"log"
	"userService/dto"
	"userService/mapper"
	"userService/model"
	"userService/repository"
	"userService/saga"
)

type UserService struct {
	Repo *repository.UserRepository
	Orchestrator *saga.Orchestrator
}

func (service *UserService) FindAll() (*[]model.User, error) {
	users,err := service.Repo.FindAll()
	if err != nil {
		return nil,err
	}
	return users,nil
}

func (service *UserService) CreateRegisteredUser(userForRegistrationDTO *dto.UserForRegistrationDTO) (string,error) {

	if userForRegistrationDTO.Password != userForRegistrationDTO.ConfirmedPassword{
		return "",errors.New("Password and confirmed password are not same!")
	}
	
	existingUser,_ := service.Repo.FindUserByUsername(userForRegistrationDTO.Username)

	if existingUser != nil{
		return "",errors.New("User with same name aleready exist!")
	}

	userForRegistration := mapper.ConvertUserForRegistrationDTOToRegisteredUser(userForRegistrationDTO)
	idString,err := service.Repo.CreateRegisteredUser(userForRegistration)
	if err != nil{
		fmt.Println(err)
		return  idString,err
	}
	return idString,nil
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

func (service *UserService) FindUserByUserId(userId string) (*model.RegisteredUser,error){
	user,err := service.Repo.FindUserByUserId(userId)
	if err != nil{
		return nil, err
	}
	return user, nil
}

func (service *UserService) FindUserByUserIdAndGetHisUsernameAndProfilePicture(userId string) (*dto.UsernameAndProfilePictureDTO,error){
	user,err := service.Repo.FindUserByUserId(userId)
	if err != nil{
	return nil, err
	}
	var usernameAndProfilePictureDTO dto.UsernameAndProfilePictureDTO
	usernameAndProfilePictureDTO.Username = user.Username
	usernameAndProfilePictureDTO.ProfilePicture = user.ProfilePicture
	return &usernameAndProfilePictureDTO, nil
}

func (service *UserService) SearchUser(username string,searchContent string) (*[]dto.UserFromSearchDTO,error){
	users,err := service.Repo.FindAllUsersBySearchingContent(username,searchContent)
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

func (service *UserService) ConvertUsernamesToUsers(usernames dto.UsernamesDTO) (*[]dto.UserByUsernameDTO, error) {
	users,err := service.Repo.FindUserIdByUsername(usernames)
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

func (service *UserService) UpdatePublicProfileSetting(username string,setting string) error {
	return service.Repo.UpdatePublicProfileSetting(username,setting  == "true")
}

func (service *UserService) UpdateMessageRequestSetting(username string, setting string) error {
	return service.Repo.UpdateMessageRequestSetting(username,setting  == "true")

}

func (service *UserService) UpdateAllowTagsSetting(username string, setting string) error {
	return service.Repo.UpdateAllowTagsSetting(username,setting  == "true")

}

func (service *UserService) UpdateLikeNotificationSetting(username string, setting string) error {
	return service.Repo.UpdateLikeNotificationSetting(username,setting  == "true")

}

func (service *UserService) UpdateCommentNotificationSetting(username string, setting string) error {
	return service.Repo.UpdateCommentNotificationSetting(username,setting  == "true")

}

func (service *UserService) UpdateMessageRequestNotificationSetting(username string, setting string) error{
	return service.Repo.UpdateMessageRequestNotificationSetting(username,setting  == "true")

}

func (service *UserService) UpdateMessageNotificationSetting(username string, setting string) error {
	return service.Repo.UpdateMessageNotificationSetting(username,setting  == "true")

}

func (service *UserService) UpdateFollowRequestNotificationSetting(username string, setting string) error {
	return service.Repo.UpdateFollowRequestNotificationSetting(username,setting  == "true")

}

func (service *UserService) UpdateFollowNotificationSetting(username string, setting string) error {
	return service.Repo.UpdateFollowNotificationSetting(username,setting  == "true")

}

func (service *UserService) UpdateVerificationSettings(userId string,category model.Category) error {
	return service.Repo.UpdateVerificationSettings(userId,category)
}

func (service *UserService) DeleteUserByUserId(userId string) error {
	err := service.Repo.DeleteUserByUserId(userId)
	if err != nil {
		return err
	}
	return nil
}


//saga deo

func (service *UserService) RedisConnection() {
	// create client and ping redis
	var err error
	client := redis.NewClient(&redis.Options{Addr: "localhost:6379",})
	if _, err = client.Ping().Result(); err != nil {
		log.Fatalf("error creating redis client %s", err)
	}

	// subscribe to the required channels
	pubsub := client.Subscribe(saga.UserChannel, saga.ReplyChannel)
	if _, err = pubsub.Receive(); err != nil {
		log.Fatalf("error subscribing %s", err)
	}
	defer func() { _ = pubsub.Close() }()
	ch := pubsub.Channel()

	log.Println("starting the user service")
	for {
		select {
		case msg := <-ch:
			m := saga.Message{}
			err := json.Unmarshal([]byte(msg.Payload), &m)
			if err != nil {
				log.Println(err)
				continue
			}

			switch msg.Channel {
			case saga.UserChannel:

				// Happy Flow
				fmt.Println("Orkesrator salje na user kanal sa akcijom ",m.Action)

				if m.Action == saga.ActionStart {
					fmt.Println("user-follower-service uspesno upisao usera sa id-em : " + m.UserId)
				}

				// Rollback flow
				if m.Action == saga.ActionRollback {
					fmt.Println("user-follower-service nije uspesno upisao usera sa id-em : " + m.UserId)
					fmt.Println("potrebno pozvati metodu za rollback tj da se obrise user koji se kreirau u bazu")
					err := service.Repo.DeleteUserByUserId(m.UserId)
					if err != nil {
						sendToReplyChannel(client, &m, saga.ActionError, saga.ServiceUserFollower, saga.ServiceUser)
					}
					fmt.Println("Uspesno odradio rollback")
				}

			}
		}
	}
}

func sendToReplyChannel(client *redis.Client, m *saga.Message, action string, service string, senderService string) {
	var err error
	m.Action = action
	m.Service = service
	m.SenderService = senderService
	if err = client.Publish(saga.ReplyChannel, m).Err(); err != nil {
		log.Printf("error publishing done-message to %s channel", saga.ReplyChannel)
	}
	log.Printf("done message published to channel :%s", saga.ReplyChannel)
}
