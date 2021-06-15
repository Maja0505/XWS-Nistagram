package service

import (
	"errors"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"userService/dto"
	"userService/mapper"
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

	err = service.Repo.ApproveVerificationRequest(user)
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

