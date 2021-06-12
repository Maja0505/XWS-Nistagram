package repository

import (
	"XWS-Nistagram/UserFollowersService/dto"
	"XWS-Nistagram/UserFollowersService/model"
	"fmt"
	"github.com/neo4j/neo4j-go-driver/v4/neo4j"
)

type UserFollowersRepository struct{
	Session neo4j.Session
}

func (repository *UserFollowersRepository) FollowUser(fr *model.FollowRelationship) error{

	err := repository.CreateUserNodeIfNotExist(fr.User)
	err = repository.CreateUserNodeIfNotExist(fr.FollowedUser)

	if err != nil {
		fmt.Println(err)
		return err
	}

	_,err = repository.Session.Run("MATCH (u1:User),(u2:User) WHERE u1.userId = $userId1 and u2.userId = $userId2 " +
		"MERGE (u1)-[r:follow]->(u2)",map[string]interface{}{
		"userId1" : fr.User ,
		"userId2" : fr.FollowedUser,
	})

	if err != nil{
		fmt.Println(err)
		return err
	}

	return nil
}


func (repository *UserFollowersRepository) UnfollowUser(fr *model.FollowRelationship) error{

	_,err := repository.Session.Run("match (u1:User{ userId:$userId1 } )-[r:follow]->( u2:User{ userId:$userId2 }) delete r" , map[string]interface{}{
		"userId1" : fr.User,
		"userId2" : fr.FollowedUser,
	})

	if err != nil{
		return err
	}

	return nil

}


func (repository *UserFollowersRepository) SendFollowRequest(fr *model.FollowRelationship) error {

	err := repository.CreateUserNodeIfNotExist(fr.User)
	err = repository.CreateUserNodeIfNotExist(fr.FollowedUser)

	if err != nil {
		fmt.Println(err)
		return err
	}

	_,err = repository.Session.Run("MATCH (u1:User),(u2:User) WHERE u1.userId = $userId1 and u2.userId = $userId2 " +
		"MERGE (u1)-[r:followRequest]->(u2) return r",map[string]interface{}{
		"userId1" : fr.User ,
		"userId2" : fr.FollowedUser,
	})

	if err != nil {
		fmt.Println(err)
		return err
	}

	return nil

}


func (repository *UserFollowersRepository) CreateUserNodeIfNotExist(userId string) error{
	_,err := repository.Session.Run("MERGE (u:User {userId:$userId})",map[string]interface{}{
		"userId" : userId,
	})
	return err
}


func (repository *UserFollowersRepository) DeleteAll(){
	_,err :=repository.Session.Run(`MATCH (n) DETACH DELETE n`,nil)
	if err != nil {
		fmt.Println(err)
		return
	}
}


func (repository *UserFollowersRepository) GetAllFollowedUsersByUser(userId string) (*[]interface{},error){
	var followedUsers []interface{}

	result,err := repository.Session.Run("MATCH (u1)-[r:follow]->(u2) WHERE u1.userId = $userId RETURN u2.userId",map[string]interface{}{
		"userId" : userId ,
	})

	if err != nil{
		return nil, err
	}

	for result.Next() {
		user := result.Record().Values[0]
		followedUsers = append(followedUsers, user)
	}

	return &followedUsers,nil
}


func (repository *UserFollowersRepository) GetAllFollowersByUser(userId string) (*[]interface{}, error) {
	var followedUsers []interface{}

	result,err := repository.Session.Run("MATCH (u1)-[r:follow]->(u2) WHERE u2.userId = $userId RETURN u1.userId",map[string]interface{}{
		"userId" : userId ,
	})

	if err != nil{
		return nil, err
	}

	for result.Next() {
		user := result.Record().Values[0]
		followedUsers = append(followedUsers, user)
	}

	return &followedUsers,nil
}


func (repository *UserFollowersRepository) AcceptFollowRequest(dto *dto.AcceptFollowRequestDTO) error{

	_,err := repository.Session.Run("match (u1:User{ userId:$userId1 } )-[r1:followRequest]->( u2:User{ userId:$userId2 }) " +
		"delete r1 " +
		"MERGE (u1)-[r2:follow]->(u2)" , map[string]interface{}{
		"userId1" : dto.UserWitchSendRequest,
		"userId2" : dto.User,
	})

	if err != nil{
		return err
	}

	return nil

}


func (repository *UserFollowersRepository) GetAllFollowRequests(userId string) (*[]interface{}, error) {
	var followedUsers []interface{}

	result,err := repository.Session.Run("MATCH (u1)-[r:followRequest]->(u2) WHERE u2.userId = $userId RETURN u1.userId",map[string]interface{}{
		"userId" : userId ,
	})

	if err != nil{
		return nil, err
	}

	for result.Next() {
		user := result.Record().Values[0]
		followedUsers = append(followedUsers, user)
	}

	return &followedUsers,nil
}


func (repository *UserFollowersRepository) CheckFollowing(userId string, followedUserId string) (*interface{}, error) {

	result,err := repository.Session.Run("return exists ( (:User{userId:$userId1})-[:follow]->(:User{userId:$userId2}))", map[string]interface{}{
		"userId1" : userId,
		"userId2" : followedUserId,
	})

	if err != nil{
		return nil, err
	}

	if result.Next(){
		return &result.Record().Values[0], nil
	}

	return nil, nil
}
