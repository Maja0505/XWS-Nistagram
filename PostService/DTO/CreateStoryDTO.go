package DTO

type CreateStoryDTO struct {
	UserID				string	`json:"UserID"`
	Image				string 	`json:"Image"`
	Highlights			bool 	`json:"Highlights"`
	ForCloseFriends		bool 	`json:"ForCloseFriends"`
}


