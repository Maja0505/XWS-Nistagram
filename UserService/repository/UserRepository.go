package repository

import (
	"context"
	"errors"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
	"time"
	"userService/model"
)

type UserRepository struct {
	Database *mongo.Client
}

func (repo *UserRepository) FindAll() (*[]model.User, error) {
	db := repo.Database.Database("user-service-database")
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	collection := db.Collection("users")
	var users []model.User
	cur, err := collection.Find(ctx, bson.M{})
	if err != nil {
		log.Fatal(err)
		return nil,err

	}
	err = cur.All(ctx,&users)
	if err != nil {
		log.Fatal(err)
		return nil,err

	}
	return &users,nil
}


func (repo *UserRepository) Create(user *model.User) error {
	db := repo.Database.Database("user-service-database")
	collection := db.Collection("users")
	_, err := collection.InsertOne(context.TODO(), &user)
	if err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}

func (repo *UserRepository) Update(id primitive.ObjectID, user *model.User) error {
	db := repo.Database.Database("user-service-database")
	collection := db.Collection("users")
	_, err := collection.UpdateOne(
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

func (repo *UserRepository) FindUserByUsername(username string) ( *model.RegisteredUser, error){
	db := repo.Database.Database("user-service-database").Collection("users")
	var user model.RegisteredUser
	err := db.FindOne(context.TODO(),bson.M{"username" : username}).Decode(&user)
	if err != nil {
		return nil, err
	}
	return &user,nil

}