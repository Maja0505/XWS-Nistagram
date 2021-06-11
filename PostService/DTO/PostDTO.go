package DTO

import (
	"github.com/gocql/gocql"
	"time"
)

type PostDTO struct{
	ID 				gocql.UUID 		`json:"ID"`
	CreatedAt 		time.Time 		`json:"CreatedAt"`
	Description  	string 			`json:"Description"`
	DislikesCount 	int64 			`json:"DislikesCount"`
	LikesCount		int64				`json:"LikesCount"`
	Image 			string 			`json:"Image"`
	UserID 			string		`json:"UserID"`
	CommentsCount 	int64  			`json:"CommentsCount"`
}
