package Model

import (
	"github.com/gocql/gocql"
)

type ReportedContent struct {
	ID 				gocql.UUID 	`json:"ID"`
	Description  	string 		`json:"Description"`
	ContentID 			string 		`json:"Image"`
	UserID 			string		`json:"UserID"`
	AdminID 			string		`json:"AdminID"`
}