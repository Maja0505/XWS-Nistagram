package model

type Notification struct {
	ID bool `json:"id"`
	Opened bool `json:"opened"`
	UserFrom string `json:"user_from,omitempty"`
	Channel string `json:"channel,omitempty"`
	Content string `json:"content,omitempty"`
	PostId string `json:"post_id,omitempty"`
	Media string `json:"media,omitempty"`
	Comment string `json:"comment,omitempty"`
}
