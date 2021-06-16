package Model

import (
	"github.com/gocql/gocql"
	"time"
)

type Comment struct {
	ID 				gocql.UUID	`json:"ID"`
	PostID 			gocql.UUID 	`json:"PostID"`
	UserID 			string 		`json:"UserID"`
	CreatedAt 		time.Time 	`json:"CreatedAt"`
	Content  		string 		`json:"Content"`
}