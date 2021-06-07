package repository

import (
	"context"
	"errors"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
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


func (repo *UserRepository) CreateRegisteredUser(userForRegistration *model.RegisteredUser) error {
	db := repo.Database.Database("user-service-database")
	collection := db.Collection("users")
	_, err := collection.InsertOne(context.TODO(), &userForRegistration)
	if err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}

func (repo *UserRepository) UpdateRegisteredUserProfile(username string, registeredUser *model.RegisteredUser) error {
	db := repo.Database.Database("user-service-database")
	collection := db.Collection("users")

	_, err := collection.UpdateOne(
		context.TODO(),
		bson.M{"username": username},
		bson.D{
			{"$set", &registeredUser},
		},
	)
	if err != nil{
		fmt.Println(err)
		return  errors.New("No exist")
	}
	return nil
}

func (repo *UserRepository) FindUserByUsername(username string) ( *model.RegisteredUser, error){
	db := repo.Database.Database("user-service-database")
	coll := db.Collection("users")
	var user model.RegisteredUser
	err := coll.FindOne(context.TODO(),bson.M{"username" : username}).Decode(&user)
	if err != nil {
		return nil, err
	}
	return &user,nil

}

func (repo *UserRepository) FindAllUsersBySearchingContent(searchContent string) (*[]model.RegisteredUser,error) {
	db := repo.Database.Database("user-service-database")
	coll := db.Collection("users")
	var users []model.RegisteredUser
	cursor,err := coll.Find(context.TODO(),bson.M{"username" : bson.D{{"$regex", searchContent + ".*"}}})
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	err = cursor.All(context.TODO(),&users)
	if err != nil{
		return nil, err
	}
	return &users,nil
}