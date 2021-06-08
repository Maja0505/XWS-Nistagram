package DTO

import (
	"XWS-Nistagram/PostService/Model"
	"github.com/gocql/gocql"
	"time"
)

type PostDTO struct{
	ID 				gocql.UUID 		`json:"ID"`
	CreatedAt 		time.Time 		`json:"CreatedAt"`
	Description  	string 			`json:"Description"`
	DislikesCount 	int 			`json:"DislikesCount"`
	LikesCount		int				`json:"LikesCount"`
	Image 			string 			`json:"Image"`
	UserID 			gocql.UUID		`json:"UserID"`
	Comments 		[]Model.Comment	`json:"Comments"`
}
