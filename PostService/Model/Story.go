package Model

import (
	"github.com/gocql/gocql"
	"time"
)

type Story struct {
	ID 					gocql.UUID 	`json:"ID"`
	UserID 				string		`json:"UserID"`
	CreatedAt 			time.Time 	`json:"CreatedAt"`
	ExpiredAt 			time.Time 	`json:"ExpiredAt"`
	Image 				string 		`json:"Image"`
	Highlights			bool		`json:"Highlights"`
	ForCloseFriends		bool 		`json:"ForCloseFriends"`
}
