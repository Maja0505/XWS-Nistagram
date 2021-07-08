package Model

import "go.mongodb.org/mongo-driver/bson/primitive"

type Gender int

const (
	Female Gender = iota
	Male
)

type UserModel struct {
	ID	primitive.ObjectID `bson:"_id,omitempty"`
	IdString string `bson:"id_string,omitempty"`
	FirstName string  `bson:"first_name,omitempty"`
	LastName string  `bson:"last_name,omitempty"`
	Username string  `bson:"username,omitempty"`
	Password string  `bson:"password,omitempty"`
	Email string  `bson:"email,omitempty"`
	PhoneNumber string  `bson:"phone_number,omitempty"`
	DateOfBirth *primitive.DateTime `bson:"date_of_birth,omitempty"`
	Gender *Gender `bson:"gender,omitempty"`
}

type RegisteredUser struct  {
	User `bson:",inline"`
	Following int64 `bson:"following,omitempty"`
	Followers int64 `bson:"followers,omitempty"`
	NumberOfPosts int64 `bson:"number_of_posts,omitempty"`
	WebSite string `bson:"web_site,omitempty"`
	Biography string `bson:"biography,omitempty"`
	ProfilePicture string `bson:"profile_picture,omitempty"`
	ProfileSettings ProfileSettings `bson:"profile_settings,omitempty"`
	NotificationSettings NotificationSettings `bson:"notification_settings,omitempty"`
	VerificationSettings VerificationSettings `bson:"verification_settings,omitempty"`
}


type ProfileSettings struct {
	Public bool `bson:"public,omitempty"`
	MessageRequest bool `bson:"message_request,omitempty"`
	AllowTags bool `bson:"allow_tags,omitempty"`
	ActiveProfile bool `bson:"active_profile,omitempty"`
}

type NotificationSettings struct {
	LikeNotification bool `bson:"like_notification,omitempty"`
	CommentNotification bool `bson:"comment_notification,omitempty"`
	MessageRequestNotification bool `bson:"message_request_notification,omitempty"`
	MessageNotification bool `bson:"message_notification,omitempty"`
	FollowRequestNotification bool `bson:"follow_request_notification,omitempty"`
	FollowNotification bool `bson:"follow_notification,omitempty"`
}

type Category int

const (
	Blogger_Influencer Category = iota
	Sports
	News_Media
	Business_Brand_Organization
	Government_Politics
	Music
	Fashion
	Entertainment
	Other
)

type VerificationSettings struct {
	Verified bool `bson:"verified,omitempty"`
	Category *Category `bson:"category,omitempty"`
}
