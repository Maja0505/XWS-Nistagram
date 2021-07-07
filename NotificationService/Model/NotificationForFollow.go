package Model

type NotificationForFollow struct {
	UserWhoFollow string `json:"user_who_follow,omitempty"`
	//CreatedAt redis.TimeCmd `json:"created_at,omitempty"`
	Opened bool `json:"opened"`
	PostId string `json:"post_id,omitempty"`
	Comment string `json:"comment,omitempty"`
	Content string `json:"content,omitempty"`
	Channel string `json:"channel,omitempty"`
	Command int    `json:"command,omitempty"`
	Err     string `json:"err,omitempty"`
}

const (
	CommandSubscribe = iota
	CommandUnsubscribe
	CommandSendNotification
)