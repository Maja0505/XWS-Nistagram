package DTO

import (
	"github.com/gocql/gocql"
)

type ReportedContentDTO struct {
	ID 				gocql.UUID 	`json:"ID"`
	Description  	string 		`json:"Description"`
	ContentID 		string 		`json:"ContentID"`
	UserID 			string		`json:"UserID"`
	AdminID 		string		`json:"AdminID"`
}
