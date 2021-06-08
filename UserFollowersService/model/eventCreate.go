package model

import (
	"XWS-Nistagram/UserFollowersService/events"
	"log"
	"sync"
)
func CreateParticipant(e events.Event, label string, c chan error, mutex *sync.Mutex) {
	if e.GetField(label, "Email") == "" {
		c <- nil
		return
	}
	// beginning of the critical section
	mutex.Lock()
	result, err := events.Session.Run(`MATCH(a:EVENT) WHERE a.name=$EventName
// low level database functions
	CREATE (n:INCHARGE {name:$name, registrationNumber:$registrationNumber,
		email:$email, phoneNumber:$phoneNumber, gender: $gender})<-[:`+label+`]-(a) `, map[string]interface{}{
		"EventName":          e.Name,
		"name":               e.GetField(label, "Name"),
		"registrationNumber": e.GetField(label, "RegistrationNumber"),
		"email":              e.GetField(label, "Email"),
		"phoneNumber":        e.GetField(label, "PhoneNumber"),
		"gender":             e.GetField(label, "Gender"),
	})
	if err != nil {
		c <- err
		return
	}
	// critical section ends
	mutex.Unlock()
	if err = result.Err(); err != nil {
		c <- err
		return
	}
	log.Printf("Created %s node", label)
	c <- nil
	return
}

func CreateEvent(e events.Event, ce chan error) {
	c := make(chan error)

	// creating an event
	result, err := events.Session.Run(`CREATE (n:EVENT {name:$name, clubName:$clubName, toDate:$toDate, 
		fromDate: $fromDate, toTime:$toTime, fromTime:$fromTime, budget:$budget, 
		description:$description, category:$category, venue:$venue, attendance:$attendance, 
		expectedParticipants:$expectedParticipants, PROrequest:$PROrequest, 
		campusEngineerRequest:$campusEngineerRequest, duration:$duration}) 
		RETURN n.name`, map[string]interface{}{
		"name":                  e.Name,
		"clubName":              e.ClubName,
		"toDate":                e.ToDate,
		"fromDate":              e.FromDate,
		"toTime":                e.ToTime,
		"fromTime":              e.FromTime,
		"budget":                e.Budget,
		"description":           e.Description,
		"category":              e.Category,
		"venue":                 e.Venue,
		"PROrequest":            e.PROrequest,
		"campusEngineerRequest": e.CampusEngineerRequest,
		"duration":              e.Duration,
		"attendance":            e.Attendance,
		"expectedParticipants":  e.ExpectedParticipants,
	})
	if err != nil {
		ce <- err
		return
	}
	result.Next()
	log.Println(result.Record().GetByIndex(0).(string))
	if err = result.Err(); err != nil {
		ce <- err
		return
	}
	// CREATE STUDENT COORDINATOR, FACULTY COORDINATOR, AND SPONSOR NODES WHENEVER AN EVENT IS CREATED
	var mutex = &sync.Mutex{}
	go CreateParticipant(e, "StudentCoordinator", c, mutex)
	go CreateParticipant(e, "FacultyCoordinator", c, mutex)
	go CreateParticipant(e, "MainSponsor", c, mutex)
	err1, err2, err3 := <-c, <-c, <-c
	switch {
	case err1 != nil:
		ce <- err1
		return
	case err2 != nil:
		ce <- err2
		return
	case err3 != nil:
		ce <- err3
		return
	}
	log.Println("Created Event node")
	ce <- nil
	return
}

	func FollowUser(fr FollowRelationship, ce chan error) {
		_,err :=events.Session.Run(`MATCH (n) DETACH DELETE n`,nil)
		if err != nil {
			ce <- err
			return
	}

		result, err := events.Session.Run(`OPTIONAL MATCH (n:User{uuid:$uuid})
												 RETURN n IS NOT NULL AS Predicate`, map[string]interface{}{
			"uuid":                  fr.User,
		})

		if(result.Record() == nil){
			_, err = events.Session.Run(`CREATE (u:User{uuid:$uuid})`, map[string]interface{}{
				"uuid": fr.User,
			})
			if err != nil {
				ce <- err
				return
			}
		}
		result, err = events.Session.Run(`OPTIONAL MATCH (n:User{uuid:$uuid})
												 RETURN n IS NOT NULL AS Predicate`, map[string]interface{}{
			"uuid":                  fr.User,
		})

		if(result.Record() == nil){
			_, err = events.Session.Run(`CREATE (u1:User{uuid:$uuid})`, map[string]interface{}{
				"uuid": fr.FollowedUser,
			})
			if err != nil {
				ce <- err
				return
			}
		}

		result, err = events.Session.Run(`MATCH (u:User) WHERE u.uuid=$uuid 
				MATCH (u1:User) WHERE u.uuid=$uuid2
				CREATE (u)-[:Follows]->(u2)`, map[string]interface{}{
				"uuid":                  fr.User,
				"uuid2":              fr.FollowedUser,
			})

			ce <- nil
			return


	//	result,err=events.Session.Run(`CREATE (Danica:User{uuid: "Danica"})`,nil)
	//	result,err=events.Session.Run(`MATCH (u:User) WHERE u.uuid='Danica' return u.uid`,nil)
	//	result,err=events.Session.Run(`MATCH (u:User) WHERE u.uuid='Marko' return u.uid"})`,nil)
	//	fmt.Println(result)

	//	result, err = events.Session.Run(`MATCH (u:User) WHERE u.uuid='Danica'
	//	CREATE (u)-[:Follows]->(Marko:User{uuid:'Marko'})`, map[string]interface{}{
	//		"uuid":                  fr.User,
	//		"uuid2":              fr.FollowedUser,
	//	})
	//	if err != nil {
	//		ce <- err
	//		return
	//	}

	}