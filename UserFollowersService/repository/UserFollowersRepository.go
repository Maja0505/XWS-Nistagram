package repository

import (
	"XWS-Nistagram/UserFollowersService/model"
	"errors"
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
		"MERGE (u1)-[r:follow]->(u2) set r.close_friend=$close set r.mute=$mute",map[string]interface{}{
		"userId1" : fr.User ,
		"userId2" : fr.FollowedUser,
		"mute" : fr.Muted,
		"close" : fr.CloseFriend,
	})

	if err != nil{
		fmt.Println(err)
		return err
	}

	return nil
}


func (repository *UserFollowersRepository) UnfollowUser(fr *model.FollowRelationship) error{

	result,err := repository.Session.Run("match (u1:User{ userId:$userId1 } )-[r:follow]->( u2:User{ userId:$userId2 }) delete r return r" , map[string]interface{}{
		"userId1" : fr.User,
		"userId2" : fr.FollowedUser,
	})

	if err != nil{
		return err
	}

	if !result.Next(){
		return errors.New("User are already unfollowed, or user1,user2 or relationship doesn't exist")
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


func (repository *UserFollowersRepository) AcceptFollowRequest(user string,userWitchSendRequest string) error{

	result,err := repository.Session.Run("match (u1:User{ userId:$userId1 } )-[r1:followRequest]->( u2:User{ userId:$userId2 }) " +
		"delete r1 " +
		"MERGE (u1)-[r2:follow{close_friend:$close,mute:$mute}]->(u2) return r1" , map[string]interface{}{
		"userId1" : userWitchSendRequest,
		"userId2" : user,
		"mute" : false,
		"close" : false,

	})

	if err != nil{
		return err
	}

	if !result.Next(){
		return errors.New("Already accepted follow request, or user1,user2 or relationship does't exist")
	}

	return nil

}


func (repository *UserFollowersRepository) CancelFollowRequest(user string,userWitchSendRequest string) error{
	result,err := repository.Session.Run("match (u1:User{ userId:$userWitchSendRequest } )-[r1:followRequest]->( u2:User{ userId:$user }) " +
		"delete r1 return r1", map[string]interface{}{
		"userWitchSendRequest" : userWitchSendRequest,
		"user" : user,
	})

	if err != nil{
		return err
	}

	if !result.Next(){
		return errors.New("Already cancel follow request, or user1,user2 or relationship does't exits!")
	}

	return nil

}


func (repository *UserFollowersRepository) SetFriendForClose(userId string,friendId string, close bool) error{
	result,err := repository.Session.Run("match (:User{userId:$userId})-[r:follow]->(:User{userId:$friendId}) set r.close_friend=$close return r.close_friend;", map[string]interface{}{
		"userId" : userId,
		"friendId" : friendId,
		"close" : close,
	})

	if err != nil{
		return err
	}

	if !result.Next(){
		return errors.New("User,friend or relationship doesn't exist")
	}

	return nil

}


func (repository *UserFollowersRepository) SetFriendForMute(userId string,friendId string, mute bool) error{
	result,err := repository.Session.Run("match (:User{userId:$userId})-[r:follow]->(:User{userId:$friendId}) set r.mute=$mute return r.mute;", map[string]interface{}{
		"userId" : userId,
		"friendId" : friendId,
		"mute" : mute,
	})

	if err != nil{
		return err
	}

	if !result.Next(){
		return errors.New("User,friend or relationship doesn't exist")
	}

	return nil

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


func (repository *UserFollowersRepository) GetAllCloseFriends(userId string) (*[]interface{}, error) {
	var closeFriends []interface{}

	result,err := repository.Session.Run("MATCH (u1)-[r:follow{close_friend:TRUE}]->(u2) WHERE u1.userId = $userId RETURN u2.userId",map[string]interface{}{
		"userId" : userId ,
	})

	if err != nil{
		return nil, err
	}

	for result.Next() {
		user := result.Record().Values[0]
		closeFriends = append(closeFriends, user)
	}

	return &closeFriends,nil
}


func (repository *UserFollowersRepository) GetAllMuteFriends(userId string) (*[]interface{}, error) {
	var muteFriends []interface{}

	result,err := repository.Session.Run("MATCH (u1)-[r:follow{mute:TRUE}]->(u2) WHERE u1.userId = $userId RETURN u2.userId",map[string]interface{}{
		"userId" : userId ,
	})

	if err != nil{
		return nil, err
	}

	for result.Next() {
		user := result.Record().Values[0]
		muteFriends = append(muteFriends, user)
	}

	return &muteFriends,nil
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


func (repository *UserFollowersRepository) CheckRequested(userId string, followedUserId string) (*interface{}, error) {

	result,err := repository.Session.Run("return exists ( (:User{userId:$userId1})-[:followRequest]->(:User{userId:$userId2}))", map[string]interface{}{
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


func (repository *UserFollowersRepository) DeleteAll(){
	_,err :=repository.Session.Run(`MATCH (n) DETACH DELETE n`,nil)
	if err != nil {
		fmt.Println(err)
		return
	}
}