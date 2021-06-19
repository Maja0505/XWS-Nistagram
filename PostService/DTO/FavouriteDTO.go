package DTO

import (
	"github.com/gocql/gocql"
)

type FavouriteDTO struct{
	PostID 			gocql.UUID 		`json:"PostID"`
	UserID 			string			`json:"UserID"`
	Collection 		string  		`json:"Collection"`
}