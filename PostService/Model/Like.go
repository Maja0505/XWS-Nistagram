package Model

import "github.com/gocql/gocql"

type Like struct {
	PostID 			gocql.UUID 		`json:"PostID"`
	UserID 			string 			`json:"UserID"`
	Username string
	PostUserID string
	MediaID string
}