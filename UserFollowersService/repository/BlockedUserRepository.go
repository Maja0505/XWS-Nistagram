package repository

import (
	"XWS-Nistagram/UserFollowersService/model"
	"errors"
	"fmt"
	"github.com/neo4j/neo4j-go-driver/v4/neo4j"
	"time"
)

type BlockedUserRepository struct{
	Driver neo4j.Driver
}


func (repository *BlockedUserRepository) BlockUser(br *model.BlockRelationship) error{

	session := repository.Driver.NewSession(neo4j.SessionConfig{
		AccessMode:   neo4j.AccessModeWrite,
	})
	defer session.Close()

	fmt.Println("Pocetak kreiranja blokiranja : ",time.Now())
	_,err := session.WriteTransaction(func(tx neo4j.Transaction) (interface{}, error){

		query :=  "MATCH (u1:User),(u2:User) " +
			"WHERE u1.userId = $userId1 and u2.userId = $userId2 " +
			"MERGE (u1)-[r:block]->(u2) "

		_,err := tx.Run(query,map[string]interface{}{
			"userId1" : br.User ,
			"userId2" : br.BlockedUser,
		})
		if err != nil{
			fmt.Println(err)
			return nil,err
		}
		return nil, nil
	})
	fmt.Println("Zavrseno kreiranje blokiranja : ",time.Now())

	fmt.Println("Pocetak brisanja follow veze ako postoji : ",time.Now())
	_,err = session.WriteTransaction(func(tx neo4j.Transaction) (interface{}, error){

		query :=  "match (u1)-[r:follow]->(u2) " +
			"WHERE u1.userId = $user and u2.userId = $blockedUser " +
			"delete r"

		_,err := tx.Run(query,map[string]interface{}{
			"user" : br.User ,
			"blockedUser" : br.BlockedUser,
		})
		if err != nil{
			fmt.Println(err)
			return nil,err
		}
		return nil, nil
	})
	fmt.Println("Zavrseno brisanja follow veze ako postoji : ",time.Now())

	fmt.Println("Pocetak brisanja follow veze ako postoji : ",time.Now())
	_,err = session.WriteTransaction(func(tx neo4j.Transaction) (interface{}, error){

		query :=  "match (u1)-[r:follow]->(u2) " +
			"WHERE u1.userId = $user and u2.userId = $blockedUser " +
			"delete r"

		_,err := tx.Run(query,map[string]interface{}{
			"user" : br.BlockedUser ,
			"blockedUser" : br.User,
		})
		if err != nil{
			fmt.Println(err)
			return nil,err
		}
		return nil, nil
	})
	fmt.Println("Zavrseno brisanja follow veze ako postoji : ",time.Now())

	fmt.Println("Pocetak brisanja followRequest veze ako postoji : ",time.Now())
	_,err = session.WriteTransaction(func(tx neo4j.Transaction) (interface{}, error){

		query :=  "match (u1)-[r:followRequest]->(u2) " +
			"WHERE u1.userId = $user and u2.userId = $blockedUser " +
			"delete r"

		_,err := tx.Run(query,map[string]interface{}{
			"user" : br.User ,
			"blockedUser" : br.BlockedUser,
		})
		if err != nil{
			fmt.Println(err)
			return nil,err
		}
		return nil, nil
	})
	fmt.Println("Zavrseno brisanja follow veze ako postoji : ",time.Now())

	fmt.Println("Pocetak brisanja followRequest veze ako postoji : ",time.Now())
	_,err = session.WriteTransaction(func(tx neo4j.Transaction) (interface{}, error){

		query :=  "match (u1)-[r:followRequest]->(u2) " +
			"WHERE u1.userId = $user and u2.userId = $blockedUser " +
			"delete r"

		_,err := tx.Run(query,map[string]interface{}{
			"user" : br.BlockedUser ,
			"blockedUser" : br.User,
		})
		if err != nil{
			fmt.Println(err)
			return nil,err
		}
		return nil, nil
	})
	fmt.Println("Zavrseno brisanja follow veze ako postoji : ",time.Now())

	if err != nil{
		return err
	}

	return nil

}

func (repository *BlockedUserRepository) GetAllBlockedUsers(user string) (*[]interface{},error){

	var blockedUsers []interface{}

	session := repository.Driver.NewSession(neo4j.SessionConfig{
		AccessMode:   neo4j.AccessModeRead,
	})
	defer session.Close()

	fmt.Println("Pocetak dobavljanja svih blokiranih : ",time.Now())
	_,err := session.ReadTransaction(func(tx neo4j.Transaction) (interface{}, error){

		query :=  "match (u1:User)-[r:block]->(u2) " +
			"WHERE u1.userId = $user " +
			"return u2.userId "

		result,err := tx.Run(query,map[string]interface{}{
			"user" : user,
		})
		if err != nil{
			fmt.Println(err)
			return nil,err
		}
		for result.Next(){
			blockedUsers = append(blockedUsers, result.Record().Values[0])
		}

		return &blockedUsers,nil
	})
	fmt.Println("Zavrseno dobavljanja svih blokiranih : ",time.Now())


	if err != nil {
		return nil, err
	}

	return &blockedUsers, err

}

func (repository *BlockedUserRepository) CheckBlock(userId string, blockedUserId string) (*interface{}, error) {

	session := repository.Driver.NewSession(neo4j.SessionConfig{
		AccessMode:   neo4j.AccessModeRead,
	})
	defer session.Close()

	fmt.Println("Pocetak provere da li je blokiran : ",time.Now())
	result,err := session.ReadTransaction(func(tx neo4j.Transaction) (interface{}, error){

		query :=  "match (u1)-[r:block]->(u2:User) " +
			"WHERE u1.userId = $userId1 and u2.userId = $userId2 " +
			"return r "

		result,err := tx.Run(query,map[string]interface{}{
			"userId1" : userId,
			"userId2" : blockedUserId,
		})
		if err != nil{
			fmt.Println(err)
			return nil,err
		}
		for result.Next(){
			return true,nil
		}

		return false,nil
	})
	fmt.Println("Zavrseno provere da li je blokiran : ",time.Now())


	if err != nil{
		return nil, err
	}

	return &result, nil
}

func (repository *BlockedUserRepository) UnblockUser(m *model.BlockRelationship) error {

	session := repository.Driver.NewSession(neo4j.SessionConfig{
		AccessMode:   neo4j.AccessModeWrite,
	})
	defer session.Close()

	fmt.Println("Pocetak odblokiranja korisika : ",time.Now())
	_,err := session.WriteTransaction(func(tx neo4j.Transaction) (interface{}, error){

		query :=  "match (u1)-[r:block]->(u2:User) " +
			"WHERE u1.userId = $userId1 and u2.userId = $userId2 " +
			"delete r return r "

		result,err := tx.Run(query,map[string]interface{}{
			"userId1" : m.User,
			"userId2" : m.BlockedUser,
		})
		if err != nil{
			fmt.Println(err)
			return nil,err
		}
		if !result.Next(){
			return nil,errors.New("User are already unblocked, or user1,user2 or relationship doesn't exist")
		}
		return nil, nil
	})
	fmt.Println("Zavrseno odblokiranja korisika : ",time.Now())

	if err != nil{
		return err
	}

	return nil
}