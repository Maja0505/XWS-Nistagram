package repository

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"userService/model"
)

type VerificationRequestRepository struct {
	Database *mongo.Client
}

func (repo *VerificationRequestRepository) Create(verificationRequest *model.VerificationRequest) error {
	db := repo.Database.Database("user-service-database").Collection("requests")
	_,err := db.InsertOne(context.TODO(),&verificationRequest)
	if err != nil{
		fmt.Println(err)
		return err
	}
	return nil
}