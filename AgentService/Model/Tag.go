package Model

import "github.com/gocql/gocql"

type Tag struct {
	Tag			  string     	 `json:"Tag"`
	PostID        gocql.UUID     `json:"PostID"`
}
