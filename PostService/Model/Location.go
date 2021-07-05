package Model

import "github.com/gocql/gocql"

type Location struct {
	Location	  string     	 `json:"Location"`
	PostID        gocql.UUID     `json:"PostID"`
}