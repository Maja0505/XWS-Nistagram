package repository

import (
	"context"
	"errors"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"userService/model"
)

type UserRepository struct {
	Database *mongo.Client
}

func (repo *UserRepository) FindAll() (*[]model.User, error) {
	ss := repo.Database.Database("user-service-database").Collection("users")
	var users []model.User
	cur, _ := ss.Find(context.TODO(),bson.M{})
	_ = cur.All(context.TODO(),&users)
	return &users,nil
}

func (repo *UserRepository) Create(user *model.User) error {
	db := repo.Database.Database("user-service-database").Collection("users")
	_,err := db.InsertOne(context.TODO(),&user)
	if err != nil{
		fmt.Println(err)
		return err
	}
	return nil
}

func (repo *UserRepository) Update(id primitive.ObjectID, user *model.User) error {
	db := repo.Database.Database("user-service-database").Collection("users")
	_, err := db.UpdateOne(
		context.TODO(),
		bson.M{"_id": id},
		bson.D{
			{"$set", &user},
		},
	)
	if err != nil{
		fmt.Println(err)
		return  errors.New("No exist")
	}
	return nil
}