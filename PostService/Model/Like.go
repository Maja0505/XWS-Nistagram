package Model

import "github.com/gocql/gocql"

type Like struct {
	PostID 			gocql.UUID 		`json:"PostID"`
	UserID 			string 			`json:"UserID"`
}