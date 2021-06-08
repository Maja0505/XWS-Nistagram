package events

import (
	"github.com/neo4j/neo4j-go-driver/v4/neo4j"
	"reflect"
)

var Session neo4j.Session
type Participant struct {
	Name               string `json:"name"`
	RegistrationNumber string `json:"registrationNumber"`
	Email              string `json:"email"`
	PhoneNumber        string `json:"phoneNumber"`
	Gender             string `json:"gender"`
}
type Event struct {
	ClubName              string      `json:"clubName"`
	Name                  string      `json:"name"`
	ToDate                string      `json:"toDate"`
	FromDate              string      `json:"fromDate"`
	ToTime                string      `json:"toTime"`
	FromTime              string      `json:"fromTime"`
	Budget                string      `json:"budget"`
	Description           string      `json:"description"`
	Category              string      `json:"category"`
	Venue                 string      `json:"venue"`
	Attendance            string      `json:"attendance"`
	ExpectedParticipants  string      `json:"expectedParticipants"`
	FacultyCoordinator    Participant `json:"facultyCoordinator"`
	StudentCoordinator    Participant `json:"studentCoordinator"`
	PROrequest            string      `json:"PROrequest"`
	CampusEngineerRequest string      `json:"campusEngineerRequest"`
	Duration              string      `json:"duration"`
	MainSponsor           Participant `json:"mainSponsor"`
}

// To get embedded JSON fields
func (v Event) GetField(field string, value string) string {
	r := reflect.ValueOf(v)
	f := reflect.Indirect(r).FieldByName(field)
	return f.FieldByName(value).String()
}

func SetDB(session neo4j.Session){
	Session =session
}


