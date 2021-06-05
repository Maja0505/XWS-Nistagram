package model

type RegisteredUser struct  {
	User `bson:",inline"`
	Following int64 `bson:"following,omitempty"`
	Followers int64 `bson:"followers,omitempty"`
	NumberOfPosts int64 `bson:"number_of_posts,omitempty"`
	WebSite string `bson:"web_site,omitempty"`
	Biography string `bson:"biography,omitempty"`
	ProfilePicture string `bson:"profile_picture,omitempty"`
	ProfileSettings ProfileSettings `bson:"profile_settings,omitempty"`
}

type ProfileSettings struct {
	Public bool `bson:"public,omitempty"`
	MessageRequest bool `bson:"message_request,omitempty"`
	AllowTags bool `bson:"allow_tags,omitempty"`
	ActiveProfile bool `bson:"active_profile,omitempty"`
}