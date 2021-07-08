package DTO

import "github.com/gocql/gocql"

type PostDTO struct{
	ID 				gocql.UUID		`json:"ID"`
	Description  	string 			`json:"Description"`
	MediaCount		int64			`json:"MediaCount"`
	Media 			[]string 		`json:"Media"`
	UserID 			string			`json:"UserID"`
	CommentsCount 	int64  			`json:"CommentsCount"`
	Album			bool			`json:"Album"`
	Location 		string			`json:"Location"`
	RepeatCampaign 	bool			`json:"RepeatCampaign"`
	Links 			[]string 		`json:"Links"`
}
