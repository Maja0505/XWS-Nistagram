package DTO

import (
	"github.com/gocql/gocql"
	"time"
)

type UpdateStoryAgentDTO struct {
	ID 				gocql.UUID 		`json:"ID"`
	UserID 			string			`json:"UserID"`
	CreatedAt 		time.Time		`json:"CreatedAt"`
}
