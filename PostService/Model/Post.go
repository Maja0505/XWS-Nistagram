package Model

import (
	"github.com/gocql/gocql"
	"time"
)

type Post struct {
	ID 				gocql.UUID 	`json:"ID"`
	CreatedAt 		time.Time 	`json:"CreatedAt"`
	Description  	string 		`json:"Description"`
	DislikesCount 	int 		`json:"DislikesCount"`
	LikesCount		int			`json:"LikesCount"`
	Image 			string 		`json:"Image"`
	UserID 			gocql.UUID	`json:"UserID"`
	Comments 		[]Comment	`json:"Comments"`
}