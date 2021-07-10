package Model

import (
	"github.com/gocql/gocql"
	"time"
)

type Story struct {
	ID 					gocql.UUID 	`json:"ID"`
	UserID 				string		`json:"UserID"`
	CreatedAt 			time.Time 	`json:"CreatedAt"`
	Available 			bool 		`json:"Available"`
	Image 				string 		`json:"Image"`
	Highlights			bool		`json:"Highlights"`
	ForCloseFriends		bool 		`json:"ForCloseFriends"`
	Link 				string 		`json:"Link"`
}
