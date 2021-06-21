package DTO

import "github.com/gocql/gocql"

type PostDTO struct{
	ID				gocql.UUID  	`json:"ID"`
	Description  	string 			`json:"Description"`
	Image 			string 			`json:"Image"`
	UserID 			string			`json:"UserID"`
}
