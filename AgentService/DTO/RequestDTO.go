package DTO

import "github.com/gocql/gocql"

type RequestDTO struct {
	CampaignID    	gocql.UUID 	`json:"CampaignID"`
	UserID 			string     	`json:"UserID"`
}

