package DTO

import "github.com/gocql/gocql"

type StoryDTO struct {
	Media       	string 		`json:"Media"`
	Duration        int 		`json:"Duration"`
	Type      		string  	`json:"Type"`
	Subheading		string   	`json:"Subheading"`
	UserID			string 		`json:"UserID"`
	ID 				gocql.UUID 	`json:"ID"`
	ForCloseFriends	bool 		`json:"ForCloseFriends"`
	Highlights		bool		`json:"Highlights"`
	Link 			string       `json:"Link"`
}
