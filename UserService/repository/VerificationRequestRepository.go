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

func (repo *VerificationRequestRepository) ApproveVerificationRequest(user primitive.ObjectID) error{
	db := repo.Database.Database("user-service-database")
	coll := db.Collection("requests")
	var verificationRequest model.VerificationRequest

	err := coll.FindOne(context.TODO(),
		bson.M{"user": user}).Decode(&verificationRequest)

	verificationRequest.Approved = true

	_,err = coll.UpdateOne(context.TODO(),
		bson.M{"user": user},
		bson.D{
			{"$set", &verificationRequest},
		})
	if err != nil{
		return err
	}
	return nil

}
