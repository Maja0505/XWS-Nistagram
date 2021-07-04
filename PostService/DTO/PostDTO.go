package DTO


type PostDTO struct{
	Description  	string 			`json:"Description"`
	DislikesCount 	int64 			`json:"DislikesCount"`
	LikesCount		int64			`json:"LikesCount"`
	MediaCount		int64			`json:"MediaCount"`
	Media 			[]string 		`json:"Media"`
	UserID 			string			`json:"UserID"`
	CommentsCount 	int64  			`json:"CommentsCount"`
	Album			bool			`json:"Album"`
	Location 		string			`json:"Location"`
}
