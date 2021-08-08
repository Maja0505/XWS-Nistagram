package model

type RegisteredUser struct  {
	User `bson:",inline"`
	IsAgent bool `bson:"is_agent,omitempty"`
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

type VerificationSettings struct {
	Verified bool `bson:"verified,omitempty"`
	Category *Category `bson:"category,omitempty"`
}