package repository

import (
	"XWS-Nistagram/UserFollowersService/model"
	"fmt"
	"github.com/neo4j/neo4j-go-driver/v4/neo4j"
)

type UserFollowersRepository struct{
	Session neo4j.Session
}

func (repository *UserFollowersRepository) FollowUser(fr model.FollowRelationship, ce chan error) {
	//repository.DeleteAll()
	if(!repository.UserExists(fr.User)){
			fmt.Println("nije nasao prvog ")
			go repository.CreateUser(fr.User)}
	if(!repository.UserExists(fr.FollowedUser)){
			fmt.Println("nije nasao drugog ")
		go repository.CreateUser(fr.FollowedUser)
		}

	if(!repository.RelationshipExists(fr.User,fr.FollowedUser)){
		fmt.Println("kreirao vezu ")
		go repository.CreateRelationship(fr.User,fr.FollowedUser)
	}


	ce <- nil
	return
}

func (repository *UserFollowersRepository) UserExists(uuid string) bool{
	result, err := repository.Session.Run(`OPTIONAL MATCH (n:User) WHERE n.uuid = $uuid
													RETURN n IS NOT NULL AS Predicate`,
		map[string]interface{}{"uuid":uuid,})
	if(err!=nil){
		fmt.Println(err)
		return false
	}
	if(result.Next()){
		return result.Record().GetByIndex(0).(bool)
	}
	return false
}

func (repository *UserFollowersRepository) CreateUser(uuid string) {
	_, err := repository.Session.Run(`CREATE (u:User{uuid:$uuid})`, map[string]interface{}{
		"uuid": uuid,
	})
	if err != nil {
		fmt.Println(err)
		return
	}
}

func (repository *UserFollowersRepository) RelationshipExists(uuid1 string,uuid2 string) bool{
	result, err := repository.Session.Run(`MATCH  (u:User {uuid:$uuid}), (u1:User {uuid:$uuid2})
		RETURN EXISTS( (u)-[:Follows]-(u1) )`, map[string]interface{}{"uuid":  uuid1, "uuid2": uuid2,})
	if err != nil {
		fmt.Println(err)
		return false
	}
	if(result.Next()){
		return result.Record().GetByIndex(0).(bool)
	}
	fmt.Println("izasao iz ifa")
	return false
}

func (repository *UserFollowersRepository) CreateRelationship(uuid1 string,uuid2 string){
	_, err := repository.Session.Run(`MATCH (u:User),(u1:User) WHERE u.uuid = $uuid and u1.uuid = $uuid2
					CREATE (u)-[:Follows]->(u1)`, map[string]interface{}{
		"uuid":  uuid1,
		"uuid2": uuid2,})

	if err != nil {
		fmt.Println(err)
		return
	}

}

func (repository *UserFollowersRepository) DeleteAll(){
	_,err :=repository.Session.Run(`MATCH (n) DETACH DELETE n`,nil)
	if err != nil {
		fmt.Println(err)
		return
	}
}