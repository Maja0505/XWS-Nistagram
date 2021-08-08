package repository

import (
	"XWS-Nistagram/UserFollowersService/model"
	"errors"
	"fmt"
	"github.com/neo4j/neo4j-go-driver/v4/neo4j"
	"time"
)

type UserFollowersRepository struct{
	Driver neo4j.Driver
}


func (repository *UserFollowersRepository) FollowUser(fr *model.FollowRelationship) error{

	session := repository.Driver.NewSession(neo4j.SessionConfig{
		AccessMode:   neo4j.AccessModeWrite,
	})
	defer session.Close()

	fmt.Println("Pocetak kreiranja veze : ",time.Now())
	_,err := session.WriteTransaction(func(tx neo4j.Transaction) (interface{}, error){

		query :=  "MATCH (u1:User),(u2:User) " +
			"WHERE u1.userId = $userId1 and u2.userId = $userId2 " +
			"MERGE (u1)-[r:follow]->(u2) set r.close_friend=$close set r.mute=$mute "

		_,err := tx.Run(query,map[string]interface{}{
			"userId1" : fr.User ,
			"userId2" : fr.FollowedUser,
			"mute" : fr.Muted,
			"close" : fr.CloseFriend,
		})
		if err != nil{
			fmt.Println(err)
			return nil,err
		}
		return nil, nil
	})
	fmt.Println("Zavrseno kreiranje veze izmedju korisnika : ",time.Now())

	if err != nil{
		fmt.Println(err)
		return err
	}

	return nil
}


func (repository *UserFollowersRepository) UnfollowUser(fr *model.FollowRelationship) error{
	session := repository.Driver.NewSession(neo4j.SessionConfig{
		AccessMode:   neo4j.AccessModeWrite,
	})
	defer session.Close()

	fmt.Println("Pocetak brisanja veze : ",time.Now())
	_,err := session.WriteTransaction(func(tx neo4j.Transaction) (interface{}, error){

		query :=  "match (u1:User)-[r:follow]->( u2:User) " +
			      "where u1.userId = $userId1 and u2.userId = $userId2 delete r return r"

		_,err := tx.Run(query,map[string]interface{}{
			"userId1" : fr.User,
			"userId2" : fr.FollowedUser,
		})
		if err != nil{
			fmt.Println(err)
			return nil,err
		}
		return nil, nil
	})
	fmt.Println("Zavrsetak brisanja veze : ",time.Now())

	if err != nil{
		return err
	}

	return nil

}


func (repository *UserFollowersRepository) SendFollowRequest(fr *model.FollowRelationship) error {

	session := repository.Driver.NewSession(neo4j.SessionConfig{
		AccessMode:   neo4j.AccessModeWrite,
	})
	defer session.Close()

	fmt.Println("Pocetak slanja zahteva za vezu : ",time.Now())
	_,err := session.WriteTransaction(func(tx neo4j.Transaction) (interface{}, error){

		query :=  "MATCH (u1:User),(u2:User) " +
			 	  "WHERE u1.userId = $userId1 and u2.userId = $userId2 " +
				  "MERGE (u1)-[r:followRequest]->(u2) return r"

		_,err := tx.Run(query,map[string]interface{}{
			"userId1" : fr.User ,
			"userId2" : fr.FollowedUser,
		})
		if err != nil{
			fmt.Println(err)
			return nil,err
		}
		return nil, nil
	})
	fmt.Println("Zavrsetak slanja zahteva za vezu : ",time.Now())

	if err != nil {
		fmt.Println(err)
		return err
	}

	return nil

}


func (repository *UserFollowersRepository) CreateUserNodeIfNotExist(userId string) error{

	session := repository.Driver.NewSession(neo4j.SessionConfig{
		AccessMode:   neo4j.AccessModeWrite,
	})
	defer session.Close()

	fmt.Println("Pocetak kreiranja korisnika : ",time.Now())
	_,err := session.WriteTransaction(func(tx neo4j.Transaction) (interface{}, error){

		query :=  "MERGE (u:User {userId:$userId})"

		_,err := tx.Run(query,map[string]interface{}{
			"userId" : userId,
		})
		if err != nil{
			fmt.Println(err)
			return nil,err
		}
		return nil, nil
	})
	fmt.Println("Zavrsetak kreiranja korisnika : ",time.Now())

	if err != nil{
		return err
	}

	return nil
}


func (repository *UserFollowersRepository) AcceptFollowRequest(user string,userWitchSendRequest string) error{

	session := repository.Driver.NewSession(neo4j.SessionConfig{
		AccessMode:   neo4j.AccessModeWrite,
	})
	defer session.Close()

	fmt.Println("Pocetak prihvatanja zahteva za vezu : ",time.Now())
	_,err := session.WriteTransaction(func(tx neo4j.Transaction) (interface{}, error){

		query :=  "match (u1:User)-[r1:followRequest]->(u2:User) " +
			"WHERE u1.userId = $userId1 and u2.userId = $userId2 " +
			"delete r1 " +
			"MERGE (u1)-[r2:follow{close_friend:$close,mute:$mute}]->(u2) return r1"

		result,err := tx.Run(query,map[string]interface{}{
			"userId1" : userWitchSendRequest,
			"userId2" : user,
			"mute" : false,
			"close" : false,
		})
		if err != nil{
			fmt.Println(err)
			return nil,err
		}

		if !result.Next(){
			return nil,errors.New("Already accepted follow request, or user1,user2 or relationship does't exist")
		}

		return nil, nil
	})
	fmt.Println("Zavrsetak prihvatanja zahteva za vezu : ",time.Now())


	if err != nil{
		return err
	}

	return nil

}


func (repository *UserFollowersRepository) CancelFollowRequest(user string,userWitchSendRequest string) error{

	session := repository.Driver.NewSession(neo4j.SessionConfig{
		AccessMode:   neo4j.AccessModeWrite,
	})
	defer session.Close()

	fmt.Println("Pocetak odbijanja zahteva za vezu : ",time.Now())
	_,err := session.WriteTransaction(func(tx neo4j.Transaction) (interface{}, error){

		query :=  "match (u1:User)-[r1:followRequest]->(u2:User) " +
			"WHERE u1.userId = $userWitchSendRequest and u2.userId = $user " +
			"delete r1 " +
			"return r1"

		result,err := tx.Run(query,map[string]interface{}{
			"userWitchSendRequest" : userWitchSendRequest,
			"user" : user,
		})
		if err != nil{
			fmt.Println(err)
			return nil,err
		}

		if !result.Next(){
			return nil,errors.New("Already cancel follow request, or user1,user2 or relationship does't exits!")
		}

		return nil, nil
	})
	fmt.Println("Zavrsetak odbijanja zahteva za vezu : ",time.Now())


	if err != nil{
		return err
	}


	return nil

}


func (repository *UserFollowersRepository) SetFriendForClose(userId string,friendId string, close bool) error{
	session := repository.Driver.NewSession(neo4j.SessionConfig{
		AccessMode:   neo4j.AccessModeWrite,
	})
	defer session.Close()

	fmt.Println("Pocetak postavljanja za bliskog prijatelja : ",time.Now())
	_,err := session.WriteTransaction(func(tx neo4j.Transaction) (interface{}, error){

		query :=  "match (u1:User)-[r:follow]->(u2:User) " +
			"where u1.userId = $userId and u2.userId = $friendId " +
			"set r.close_friend=$close return r.close_friend"

		result,err := tx.Run(query,map[string]interface{}{
			"userId" : userId,
			"friendId" : friendId,
			"close" : close,
		})
		if err != nil{
			fmt.Println(err)
			return nil,err
		}

		if !result.Next(){
			return nil,errors.New("User,friend or relationship doesn't exist")
		}

		return nil, nil
	})
	fmt.Println("Zavrsetak postavljanja za bliskog prijatelja : ",time.Now())

	if err != nil{
		return err
	}

	return nil

}


func (repository *UserFollowersRepository) SetFriendForMute(userId string,friendId string, mute bool) error{
	session := repository.Driver.NewSession(neo4j.SessionConfig{
		AccessMode:   neo4j.AccessModeWrite,
	})
	defer session.Close()

	fmt.Println("Pocetak postavljanja za mutiranog prijatelja : ",time.Now())
	_,err := session.WriteTransaction(func(tx neo4j.Transaction) (interface{}, error){

		query :=  "match (u1:User)-[r:follow]->(u2:User) " +
			"where u1.userId = $userId and u2.userId = $friendId " +
			"set r.mute=$mute return r.mute"

		result,err := tx.Run(query,map[string]interface{}{
			"userId" : userId,
			"friendId" : friendId,
			"mute" : mute,
		})
		if err != nil{
			fmt.Println(err)
			return nil,err
		}

		if !result.Next(){
			return nil,errors.New("User,friend or relationship doesn't exist")
		}

		return nil, nil
	})
	fmt.Println("Zavrsetak postavljanja za mutiranog prijatelja : ",time.Now())

	if err != nil{
		return err
	}

	return nil

}

func (repository *UserFollowersRepository) GetAllUsers(userId string) (*[]interface{},error){
	var followedUsers []interface{}

	session := repository.Driver.NewSession(neo4j.SessionConfig{
		AccessMode:   neo4j.AccessModeRead,
	})
	defer session.Close()

	fmt.Println("Pocetak dobavljanja svih korisnika : ",time.Now())
	_,err := session.ReadTransaction(func(tx neo4j.Transaction) (interface{}, error){

		query :=  "match (u:User) where u.userId <> $userId return u.userId"

		result,err := tx.Run(query,map[string]interface{}{
			"userId" : userId ,
		})

		if err != nil{
			fmt.Println(err)
			return nil,err
		}

		for result.Next(){
			user := result.Record().Values[0]
			followedUsers = append(followedUsers, user)
		}

		return &followedUsers,nil
	})
	fmt.Println("Zavrsetak dobavljanja svih korisnika : ",time.Now())

	if err != nil{
		return nil, err
	}

	return &followedUsers,nil
}

func (repository *UserFollowersRepository) GetAllFollowedUsersByUser(userId string) (*[]interface{},error){
	var followedUsers []interface{}

	session := repository.Driver.NewSession(neo4j.SessionConfig{
		AccessMode:   neo4j.AccessModeRead,
	})
	defer session.Close()

	fmt.Println("Pocetak dobavljanja svih pratilaca : ",time.Now())
	_,err := session.ReadTransaction(func(tx neo4j.Transaction) (interface{}, error){

		query :=  "MATCH (u1:User)-[r:follow]->(u2:User) WHERE u1.userId = $userId RETURN u2.userId"

		result,err := tx.Run(query,map[string]interface{}{
			"userId" : userId ,
		})

		if err != nil{
			fmt.Println(err)
			return nil,err
		}

		for result.Next(){
			user := result.Record().Values[0]
			followedUsers = append(followedUsers, user)
		}

		return &followedUsers,nil
	})
	fmt.Println("Zavrsetak dobavljanja svih pratilaca : ",time.Now())

	if err != nil{
		return nil, err
	}

	return &followedUsers,nil
}

func (repository *UserFollowersRepository) GetAllFollowersByUser(userId string) (*[]interface{}, error) {
	var followedUsers []interface{}

	session := repository.Driver.NewSession(neo4j.SessionConfig{
		AccessMode:   neo4j.AccessModeRead,
	})
	defer session.Close()

	fmt.Println("Pocetak dobavljanja svih zapracenih : ",time.Now())
	_,err := session.ReadTransaction(func(tx neo4j.Transaction) (interface{}, error){

		query :=  "MATCH (u1:User)-[r:follow]->(u2:User) WHERE u2.userId = $userId RETURN u1.userId"

		result,err := tx.Run(query,map[string]interface{}{
			"userId" : userId ,
		})

		if err != nil{
			fmt.Println(err)
			return nil,err
		}

		for result.Next(){
			user := result.Record().Values[0]
			followedUsers = append(followedUsers, user)
		}

		return &followedUsers,nil
	})
	fmt.Println("Zavrsetak dobavljanja svih zapracenih : ",time.Now())


	if err != nil{
		return nil, err
	}


	return &followedUsers,nil
}

func (repository *UserFollowersRepository) GetAllNotMutedFollowedUsersByUser(userId string) (*[]interface{}, error){
	var followedUsers []interface{}

	session := repository.Driver.NewSession(neo4j.SessionConfig{
		AccessMode:   neo4j.AccessModeRead,
	})
	defer session.Close()

	fmt.Println("Pocetak dobavljanja svih ne mjutovanih : ",time.Now())
	_,err := session.ReadTransaction(func(tx neo4j.Transaction) (interface{}, error){

		query :=  "MATCH (u1:User)-[r:follow]->(u2:User) " +
			      "WHERE u1.userId = $userId and r.mute=FALSE " +
			      "RETURN u2.userId"

		result,err := tx.Run(query,map[string]interface{}{
			"userId" : userId ,
		})

		if err != nil{
			fmt.Println(err)
			return nil,err
		}

		for result.Next(){
			user := result.Record().Values[0]
			followedUsers = append(followedUsers, user)
		}

		return &followedUsers,nil
	})
	fmt.Println("Zavrsetak dobavljanja svih ne mjutovanih : ",time.Now())


	if err != nil{
		return nil, err
	}

	return &followedUsers,nil
}

func (repository *UserFollowersRepository) GetAllFollowsWhomUserIsCloseFriend(userId string) (*[]interface{}, error){
	var follows []interface{}

	session := repository.Driver.NewSession(neo4j.SessionConfig{
		AccessMode:   neo4j.AccessModeRead,
	})
	defer session.Close()

	fmt.Println("Pocetak dobavljanja svih kojima je blizak : ",time.Now())
	_,err := session.ReadTransaction(func(tx neo4j.Transaction) (interface{}, error){

		query :=  "MATCH (u1:User)-[r:follow]->(u2:User) " +
			"WHERE u2.userId = $userId and r.close_friend=TRUE " +
			"RETURN u1.userId"

		result,err := tx.Run(query,map[string]interface{}{
			"userId" : userId ,
		})

		if err != nil{
			fmt.Println(err)
			return nil,err
		}

		for result.Next(){
			user := result.Record().Values[0]
			follows = append(follows, user)
		}

		return &follows,nil
	})
	fmt.Println("Zavrsetak dobavljanja svih kojima je blizakh : ",time.Now())

	if err != nil{
		return nil, err
	}


	return &follows,nil
}

func (repository *UserFollowersRepository) GetAllFollowRequests(userId string) (*[]interface{}, error) {
	var followedUsers []interface{}

	session := repository.Driver.NewSession(neo4j.SessionConfig{
		AccessMode:   neo4j.AccessModeRead,
	})
	defer session.Close()

	fmt.Println("Pocetak dobavljanja svih zahteva za vezu : ",time.Now())
	_,err := session.ReadTransaction(func(tx neo4j.Transaction) (interface{}, error){

		query :=  "MATCH (u1:User)-[r:followRequest]->(u2:User) " +
			"WHERE u2.userId = $userId " +
			"RETURN u1.userId"

		result,err := tx.Run(query,map[string]interface{}{
			"userId" : userId ,
		})

		if err != nil{
			fmt.Println(err)
			return nil,err
		}

		for result.Next(){
			user := result.Record().Values[0]
			followedUsers = append(followedUsers, user)
		}

		return &followedUsers,nil
	})
	fmt.Println("Zavrsetak dobavljanja svih zahteva za vezu : ",time.Now())


	if err != nil{
		return nil, err
	}

	return &followedUsers,nil
}


func (repository *UserFollowersRepository) GetAllCloseFriends(userId string) (*[]interface{}, error) {
	var closeFriends []interface{}

	session := repository.Driver.NewSession(neo4j.SessionConfig{
		AccessMode:   neo4j.AccessModeRead,
	})
	defer session.Close()

	fmt.Println("Pocetak dobavljanja svih bliskih : ",time.Now())
	_,err := session.ReadTransaction(func(tx neo4j.Transaction) (interface{}, error){

		query :=  "MATCH (u1)-[r:follow]->(u2) " +
			"WHERE u1.userId = $userId AND r.close_friend = $true " +
			"RETURN u2.userId"

		result,err := tx.Run(query,map[string]interface{}{
			"userId" : userId ,
			"true" : true,
		})

		if err != nil{
			fmt.Println(err)
			return nil,err
		}

		for result.Next(){
			user := result.Record().Values[0]
			closeFriends = append(closeFriends, user)
		}

		return &closeFriends,nil
	})
	fmt.Println("Zavrsetak dobavljanja svih bliskih : ",time.Now())


	if err != nil{
		return nil, err
	}


	return &closeFriends,nil
}


func (repository *UserFollowersRepository) GetAllMuteFriends(userId string) (*[]interface{}, error) {
	var muteFriends []interface{}

	session := repository.Driver.NewSession(neo4j.SessionConfig{
		AccessMode:   neo4j.AccessModeRead,
	})
	defer session.Close()

	fmt.Println("Pocetak dobavljanja svih mutiranih : ",time.Now())
	_,err := session.ReadTransaction(func(tx neo4j.Transaction) (interface{}, error){

		query :=  "MATCH (u1)-[r:follow]->(u2) " +
			"WHERE u1.userId = $userId AND r.mute = $true " +
			"RETURN u2.userId"

		result,err := tx.Run(query,map[string]interface{}{
			"userId" : userId ,
			"true" : true,
		})

		if err != nil{
			fmt.Println(err)
			return nil,err
		}

		for result.Next(){
			user := result.Record().Values[0]
			muteFriends = append(muteFriends, user)
		}

		return &muteFriends,nil
	})
	fmt.Println("Zavrsetak dobavljanja svih mutiranih : ",time.Now())


	if err != nil{
		return nil, err
	}

	return &muteFriends,nil
}


func (repository *UserFollowersRepository) CheckFollowing(userId string, followedUserId string) (*interface{}, error) {

	session := repository.Driver.NewSession(neo4j.SessionConfig{
		AccessMode:   neo4j.AccessModeRead,
	})
	defer session.Close()

	fmt.Println("Pocetak provere pracenja : ",time.Now())
	result,err := session.ReadTransaction(func(tx neo4j.Transaction) (interface{}, error){

		query :=  "match (u1)-[r:follow]->(u2) " +
			"WHERE u1.userId = $userId1 AND u2.userId = $userId2 " +
			"RETURN r"

		result,err := tx.Run(query,map[string]interface{}{
			"userId1" : userId,
			"userId2" : followedUserId,
		})

		if err != nil{
			fmt.Println(err)
			return nil,err
		}

		if result.Next(){
			return true,nil
		}

		return false,nil
	})
	fmt.Println("Zavrsetak provere pracenja  : ",time.Now())

	if err != nil || result == nil {
		return nil, err
	}

	return &result,nil



	/*result,err := repository.Session.Run("return exists ( (:User{userId:$userId1})-[:follow]->(:User{userId:$userId2}))", map[string]interface{}{
		"userId1" : userId,
		"userId2" : followedUserId,
	})

	if err != nil{
		return nil, err
	}

	if result.Next(){
		return &result.Record().Values[0], nil
	}

	return nil, nil*/
}


func (repository *UserFollowersRepository) CheckRequested(userId string, followedUserId string) (*interface{}, error) {

	session := repository.Driver.NewSession(neo4j.SessionConfig{
		AccessMode:   neo4j.AccessModeRead,
	})
	defer session.Close()

	fmt.Println("Pocetak provere zahteva za pracenja : ",time.Now())
	result,err := session.ReadTransaction(func(tx neo4j.Transaction) (interface{}, error){

		query :=  "match (u1)-[r:followRequest]->(u2) " +
			"WHERE u1.userId = $userId1 AND u2.userId = $userId2 " +
			"RETURN r"

		result,err := tx.Run(query,map[string]interface{}{
			"userId1" : userId,
			"userId2" : followedUserId,
		})

		if err != nil{
			fmt.Println(err)
			return nil,err
		}

		if result.Next(){
			return true,nil
		}

		return false,nil
	})
	fmt.Println("Zavrsetak provere zahteva za pracenja  : ",time.Now())


	if err != nil || result == nil {
		return nil, err
	}


	return &result, nil
}


func (repository *UserFollowersRepository) DeleteAll(){

	session := repository.Driver.NewSession(neo4j.SessionConfig{
		AccessMode:   neo4j.AccessModeWrite,
	})
	defer session.Close()

	fmt.Println("Pocetak brisanja svega : ",time.Now())
	_,err := session.ReadTransaction(func(tx neo4j.Transaction) (interface{}, error){

		query :=  "MATCH (n) DETACH DELETE n"

		_,err := tx.Run(query,map[string]interface{}{})

		if err != nil{
			fmt.Println(err)
			return nil,err
		}

		return nil,nil
	})
	fmt.Println("Zavrsetak brisanja svega : ",time.Now())

	if err != nil {
		fmt.Println(err)
		return
	}
}

func (repository *UserFollowersRepository) CheckMuted(userId string, mutedUserId string) (*interface{}, error ){

	session := repository.Driver.NewSession(neo4j.SessionConfig{
		AccessMode:   neo4j.AccessModeRead,
	})
	defer session.Close()

	fmt.Println("Pocetak provere dal je mjutovan : ",time.Now())
	r,err := session.ReadTransaction(func(tx neo4j.Transaction) (interface{}, error){

		query :=  "match (u1)-[r:follow]->(u2) " +
			"WHERE u1.userId = $userId1 AND u2.userId = $userId2 " +
			"RETURN r.mute"

		result,err := tx.Run(query,map[string]interface{}{
			"userId1" : userId,
			"userId2" : mutedUserId,
		})

		if err != nil{
			fmt.Println(err)
			return nil,err
		}

		if result.Next(){
			return &result.Record().Values[0],nil
		}

		return nil,nil
	})
	fmt.Println("Zavrsetak provere dal je mjutovan  : ",time.Now())

	if err != nil || r == nil{
		return nil, err
	}


	return &r, nil
}

func (repository *UserFollowersRepository) CheckClosed(userId string, closedUserId string) (*interface{}, error) {

	session := repository.Driver.NewSession(neo4j.SessionConfig{
		AccessMode:   neo4j.AccessModeRead,
	})
	defer session.Close()

	fmt.Println("Pocetak provere dal je blizak : ",time.Now())
	r,err := session.ReadTransaction(func(tx neo4j.Transaction) (interface{}, error){

		query :=  "match (u1)-[r:follow]->(u2) " +
			"WHERE u1.userId = $userId1 AND u2.userId = $userId2 " +
			"RETURN r.close_friend"

		result,err := tx.Run(query,map[string]interface{}{
			"userId1" : userId,
			"userId2" : closedUserId,
		})

		if err != nil{
			fmt.Println(err)
			return nil,err
		}

		if result.Next(){
			return &result.Record().Values[0],nil
		}

		return nil,nil
	})
	fmt.Println("Zavrsetak provere dal je blizak  : ",time.Now())

	if err != nil || r == nil{
		return nil, err
	}

	return &r, nil
}