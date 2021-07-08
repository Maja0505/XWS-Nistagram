package service

import (
	"errors"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"userService/dto"
	"userService/mapper"
	"userService/model"
	"userService/repository"
)

type VerificationRequestService struct {
	Repo *repository.VerificationRequestRepository
	UserService *UserService
}


func (service *VerificationRequestService) Create(verificationRequestDTO *dto.VerificationRequestDTO) error{
	userId,err := primitive.ObjectIDFromHex(verificationRequestDTO.User)
	if err != nil {
		return err
	}

	existVerificationRequest,_ := service.Repo.GetVerificationRequestByUser(userId)

	if existVerificationRequest != nil {
		return errors.New("Aleready exists verification request for user")
	}

	vq := mapper.ConvertVerificationRequestDTOToVerificationRequest(verificationRequestDTO)
	err = service.Repo.Create(vq)
	if err != nil{
		return err
	}
	return nil
}

func (service *VerificationRequestService) Update(userString string ,verificationRequestDTO *dto.VerificationRequestDTO) error{

	user,err := primitive.ObjectIDFromHex(userString)
	if err != nil{
		return err
	}

	vq,err := service.Repo.GetVerificationRequestByUser(user)

	if err != nil {
		return err
	}

	vq.Username = verificationRequestDTO.Username
	vq.FullName = verificationRequestDTO.FullName
	vq.KnownAs = verificationRequestDTO.KnowAs
	vq.Image = verificationRequestDTO.Image
	switch verificationRequestDTO.Category {
	case "Blogger/Influencer":
		vq.Category = 0
	case "Sports":
		vq.Category = 1
	case "News/Media":
		vq.Category = 2
	case "Business/Brand/Organization":
		vq.Category = 3
	case "Government/Politics":
		vq.Category = 4
	case "Music":
		vq.Category = 5
	case "Fashion":
		vq.Category = 6
	case "Entertainment":
		vq.Category = 7
	case "Other":
		vq.Category = 8
	default:
		vq.Category = 8
	}

	err = service.Repo.Update(user,vq)
	if err != nil{
		return err
	}
	return nil
}

func (service *VerificationRequestService) GetAllVerificationRequests() (*[]dto.VerificationRequestDTO,error){

	verificationRequests,err := service.Repo.GetAllVerificationRequests()
	if err != nil {
		return nil, err
	}

	verificationRequestsDTOList := mapper.ConvertVerificationRequestsListToVerificationRequestDTOList(verificationRequests)
	return verificationRequestsDTOList, nil
}

func (service *VerificationRequestService) ApproveVerificationRequest(userString string) error{
	user,err := primitive.ObjectIDFromHex(userString)
	if err != nil{
		return err
	}

	vr,err := service.Repo.ApproveVerificationRequest(user)
	if err != nil{
		return err
	}

	err = service.UserService.UpdateVerificationSettings(userString,vr.Category)
	if err != nil{
		return err
	}

	return nil
}

func (service *VerificationRequestService) GetVerificationRequestByUser(userString string) (*dto.VerificationRequestDTO,error){
	user,err := primitive.ObjectIDFromHex(userString)
	if err != nil {
		return nil, err
	}

	vr,err := service.Repo.GetVerificationRequestByUser(user)

	if err != nil {
		return nil, err
	}

	verificationRequestDTO := mapper.ConvertVerificationRequestsToVerificationRequestDTO(vr)

	return verificationRequestDTO,nil

}

func (service *VerificationRequestService) DeleteVerificationRequest(userString string) error{
	user,err := primitive.ObjectIDFromHex(userString)
	if err != nil{
		return err
	}

	err = service.Repo.DeleteVerificationRequestByUser(user)
	if err != nil{
		return err
	}

	return nil
}

func (service *VerificationRequestService) CreateAgentRegistrationRequest(agentRegistrationRequest *model.AgentRegistrationRequest) error{

	if agentRegistrationRequest.Username == "" || agentRegistrationRequest.WebSite == "" {
		return errors.New("Username or WebSite is empty")
	}

	existRegisteredUser,_ := service.UserService.FindUserByUsername(agentRegistrationRequest.Username)

	existAgentRequest,_ := service.Repo.GetAgentRegistrationRequestByUsername(agentRegistrationRequest.Username)

	if existAgentRequest != nil || existRegisteredUser != nil {
		return errors.New("Aleready exists verification request for user or exist user with same username")
	}

	err := service.Repo.CreateAgentRegistrationRequest(agentRegistrationRequest)
	if err != nil{
		return err
	}
	return nil
}

func (service *VerificationRequestService) GetAllAgentRegistrationRequests() (*[]model.AgentRegistrationRequest,error){

	agentRegistrationRequests,err := service.Repo.GetAllAgentRegistrationRequests()
	if err != nil {
		return nil, err
	}
	return agentRegistrationRequests, nil
}

func (service *VerificationRequestService) UpdateAgentRegistrationRequestToApproved(username string ) error{

	agentRequest,err := service.Repo.GetAgentRegistrationRequestByUsername(username)
	if err != nil{
		return err
	}

	var agentForRegistration dto.UserForRegistrationDTO
	agentForRegistration.Username = agentRequest.Username
	agentForRegistration.FirstName = agentRequest.FirstName
	agentForRegistration.LastName = agentRequest.LastName
	agentForRegistration.Email = agentRequest.Email
	agentForRegistration.IsAgent = true
	agentForRegistration.Password = agentRequest.Password
	agentForRegistration.ConfirmedPassword = agentRequest.Password
	agentForRegistration.DateOfBirth = agentRequest.DateOfBirth
	agentForRegistration.PhoneNumber = agentRequest.PhoneNumber
	agentForRegistration.Gender = agentRequest.Gender
	agentForRegistration.WebSite = agentRequest.WebSite


	err = service.UserService.CreateRegisteredUser(&agentForRegistration)
	if err != nil{
		return err
	}

	err = service.Repo.UpdateAgentRegistrationRequestToApproved(username)
	if err != nil{
		return err
	}

	return nil
}

func (service *VerificationRequestService) DeleteAgentRegistrationRequestToApproved(username string ) error{

	err := service.Repo.DeleteAgentRegistrationRequestToApproved(username)
	if err != nil{
		return err
	}
	return nil
}