package DTO

import "github.com/gocql/gocql"

type UpdateLinksDTO struct{
	ID 				gocql.UUID		`json:"ID"`
	UserID 			string			`json:"UserID"`
	Links 			[]string 		`json:"Links"`
}
