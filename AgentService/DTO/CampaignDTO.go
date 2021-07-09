package DTO

import (
	"github.com/gocql/gocql"
	"time"
)

type CampaignDTO struct {
	ID 				gocql.UUID 	`json:"ID"`
	CreatedAt 		time.Time 	`json:"CreatedAt"`
	DislikesCount 	int64 		`json:"DislikesCount"`
	LikesCount		int64		`json:"LikesCount"`
	ViewsCount		int64		`json:"ViewsCount"`
	Media 			[]string 	`json:"Media"`
	Links 			[]string 	`json:"Links"`
	UserID 			string		`json:"UserID"`
	CommentsCount 	int64		`json:"CommentsCount"`
	Repeat			bool		`json:"Repeat"`
	IsPost 			bool		`json:"IsPost"`
	Start 			time.Time 	`json:"Start"`
	End 			time.Time 	`json:"End"`
	RepeatFactor	int			`json:"RepeatFactor"`
	Location 		string		`json:"Location"`
	Description  	string 		`json:"Description"`
	Tags 			[]string 	`json:"Tags"`
	Influencers 	[]string 	`json:"Influencers"`
}
