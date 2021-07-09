package repository

import (
	"XWS-Nistagram/UserFollowersService/model"
	"errors"
	"github.com/neo4j/neo4j-go-driver/v4/neo4j"
)

type BlockedUserRepository struct{
	Session neo4j.Session
}


func (repository *BlockedUserRepository) BlockUser(br *model.BlockRelationship) error{
	err := repository.CreateUserNodeIfNotExist(br.User)
	err = repository.CreateUserNodeIfNotExist(br.BlockedUser)

	if err != nil{
		return err
	}

	_,err = repository.Session.Run("MATCH (u1:User),(u2:User) WHERE u1.userId = $userId1 and u2.userId = $userId2 " +
		"MERGE (u1)-[r:block]->(u2)",map[string]interface{}{
		"userId1" : br.User ,
		"userId2" : br.BlockedUser,
	})

	if err != nil{
		return err
	}

	_,err = repository.Session.Run("match (:User{userId:$user})-[r:follow]->(:User{userId:$blockedUser}) delete r",map[string]interface{}{
		"user" : br.User ,
		"blockedUser" : br.BlockedUser,
	})
	_,err = repository.Session.Run("match (:User{userId:$user})-[r:follow]->(:User{userId:$blockedUser}) delete r",map[string]interface{}{
		"user" : br.BlockedUser ,
		"blockedUser" : br.User,
	})
	_,err = repository.Session.Run("match (:User{userId:$user})-[r:followRequest]->(:User{userId:$blockedUser}) delete r",map[string]interface{}{
		"user" : br.User ,
		"blockedUser" : br.BlockedUser,
	})
	_,err = repository.Session.Run("match (:User{userId:$user})-[r:followRequest]->(:User{userId:$blockedUser}) delete r",map[string]interface{}{
		"user" : br.BlockedUser ,
		"blockedUser" : br.User,
	})

	if err != nil {
		return err
	}

	return nil

}

func (repository *BlockedUserRepository) GetAllBlockedUsers(user string) (*[]interface{},error){

	var blockedUsers []interface{}

	result,err := repository.Session.Run("match (u1:User{userId:$user})-[r:block]->(u2) return u2.userId", map[string]interface{}{
		"user" : user,
	},
	)

	if err != nil {
		return nil, err
	}

	for result.Next(){
		blockedUsers = append(blockedUsers, result.Record().Values[0])
	}

	return &blockedUsers, err

}

func (repository *BlockedUserRepository) CreateUserNodeIfNotExist(userId string) error{
	_,err := repository.Session.Run("MERGE (u:User {userId:$userId})",map[string]interface{}{
		"userId" : userId,
	})
	return err
}

func (repository *BlockedUserRepository) CheckBlock(userId string, blockedUserId string) (*interface{}, error) {

	result,err := repository.Session.Run("return exists ( (:User{userId:$userId1})-[:block]->(:User{userId:$userId2}))", map[string]interface{}{
		"userId1" : userId,
		"userId2" : blockedUserId,
	})

	if err != nil{
		return nil, err
	}

	if result.Next(){
		return &result.Record().Values[0], nil
	}

	return nil, nil
}

func (repository *BlockedUserRepository) UnblockUser(m *model.BlockRelationship) error {
	result,err := repository.Session.Run("match (u1:User{ userId:$userId1 } )-[r:block]->( u2:User{ userId:$userId2 }) delete r return r" , map[string]interface{}{
		"userId1" : m.User,
		"userId2" : m.BlockedUser,
	})

	if err != nil{
		return err
	}

	if !result.Next(){
		return errors.New("User are already unfollowed, or user1,user2 or relationship doesn't exist")
	}

	return nil
}