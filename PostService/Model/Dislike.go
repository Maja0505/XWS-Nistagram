package Model

import (
	"github.com/gocql/gocql"
)

type Dislike struct {
	PostID 			gocql.UUID 		`json:"PostID"`
	UserID 			string 			`json:"UserID"`
}