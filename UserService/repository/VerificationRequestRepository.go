package repository

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"userService/model"
)

type VerificationRequestRepository struct {
	Database *mongo.Client
}

func (repo *VerificationRequestRepository) Create(verificationRequest *model.VerificationRequest) error {
	db := repo.Database.Database("user-service-database")
	coll := db.Collection("requests")
	_,err := coll.InsertOne(context.TODO(),&verificationRequest)
	if err != nil{
		return err
	}
	return nil
}

func (repo *VerificationRequestRepository) Update(user primitive.ObjectID,verificationRequest *model.VerificationRequest) error{
	db := repo.Database.Database("user-service-database")
	coll := db.Collection("requests")
	_,err := coll.UpdateOne(context.TODO(),
		bson.M{"user": user},
		bson.D{
			{"$set", &verificationRequest},
		})
	if err != nil{
		return err
	}
	return nil
}

func (repo *VerificationRequestRepository) GetVerificationRequestByUser(user primitive.ObjectID) ( *model.VerificationRequest,error) {
	db := repo.Database.Database("user-service-database")
	coll := db.Collection("requests")
	var verificationRequest model.VerificationRequest
	err := coll.FindOne(context.TODO(),
		bson.M{"user": user}).Decode(&verificationRequest)
	if err != nil{
		return nil,err
	}
	return &verificationRequest,nil
}

func (repo *VerificationRequestRepository) GetAllVerificationRequests() ( *[]model.VerificationRequest,error) {
	db := repo.Database.Database("user-service-database")
	coll := db.Collection("requests")
	var verificationRequests []model.VerificationRequest
	cur,err := coll.Find(context.TODO(),bson.M{"approved": false})
	if err != nil{
		return nil,err
	}

	err = cur.All(context.TODO(),&verificationRequests)
	if err != nil{
		return nil, err
	}


	return &verificationRequests,nil
}

func (repo *VerificationRequestRepository) ApproveVerificationRequest(user primitive.ObjectID) (*model.VerificationRequest,error) {
	db := repo.Database.Database("user-service-database")
	coll := db.Collection("requests")
	var verificationRequest model.VerificationRequest

	err := coll.FindOne(context.TODO(),
		bson.M{"user": user}).Decode(&verificationRequest)

	if err != nil {
		return nil,err
	}

	verificationRequest.Approved = true

	_,err = coll.UpdateOne(context.TODO(),
		bson.M{"user": user},
		bson.D{
			{"$set", &verificationRequest},
		})
	if err != nil{
		return nil,err
	}
	return &verificationRequest,nil

}

func (repo *VerificationRequestRepository) DeleteVerificationRequestByUser(user primitive.ObjectID) error{
	db := repo.Database.Database("user-service-database")
	coll := db.Collection("requests")
	_,err := coll.DeleteOne(context.TODO(),
		bson.M{"user": user})
	if err != nil{
		return err
	}
	return nil
}

func (repo *VerificationRequestRepository) CreateAgentRegistrationRequest(agentRegistrationRequest *model.AgentRegistrationRequest) error {
	db := repo.Database.Database("user-service-database")
	coll := db.Collection("agent-registration-requests")
	_,err := coll.InsertOne(context.TODO(),&agentRegistrationRequest)
	if err != nil{
		return err
	}
	return nil
}

func (repo *VerificationRequestRepository) GetAllAgentRegistrationRequests() ( *[]model.AgentRegistrationRequest,error) {
	db := repo.Database.Database("user-service-database")
	coll := db.Collection("agent-registration-requests")
	var agentRegistrationRequests []model.AgentRegistrationRequest
	cur,err := coll.Find(context.TODO(),bson.M{"approved": false})
	if err != nil{
		return nil,err
	}

	err = cur.All(context.TODO(),&agentRegistrationRequests)
	if err != nil{
		return nil, err
	}


	return &agentRegistrationRequests,nil
}

func (repo *VerificationRequestRepository) GetAgentRegistrationRequestByUsername(username string) (*model.AgentRegistrationRequest,error) {
	db := repo.Database.Database("user-service-database")
	coll := db.Collection("agent-registration-requests")
	var agentRegistrationRequest model.AgentRegistrationRequest
	err := coll.FindOne(context.TODO(),
		bson.M{"username": username}).Decode(&agentRegistrationRequest)
	if err != nil{
		return nil,err
	}
	return &agentRegistrationRequest,nil
}

func (repo *VerificationRequestRepository) UpdateAgentRegistrationRequestToApproved(username string) error{
	db := repo.Database.Database("user-service-database")
	coll := db.Collection("agent-registration-requests")
	_,err := coll.UpdateOne(context.TODO(),
		bson.M{"username": username},
		bson.D{{"$set",bson.D{{"approved",true}}}})
	if err != nil{
		return err
	}
	return nil
}

func (repo *VerificationRequestRepository) DeleteAgentRegistrationRequestToApproved(username string) error{
	db := repo.Database.Database("user-service-database")
	coll := db.Collection("agent-registration-requests")
	_,err := coll.DeleteOne(context.TODO(),
		bson.M{"username": username})
	if err != nil{
		return err
	}
	return nil
}
