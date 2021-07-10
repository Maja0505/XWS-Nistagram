package DTO

import (
	"github.com/gocql/gocql"
	"time"
)

type AddInfluencerDTO struct {
	ID 				gocql.UUID 	`json:"ID"`
	UserID 			string		`json:"UserID"`
	Start 			time.Time 	`json:"Start"`
	InfluencerID 	string		`json:"InfluencerID"`
}
