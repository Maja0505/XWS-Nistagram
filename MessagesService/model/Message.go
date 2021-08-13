package model

type MessageType int


const (
	TextMessage MessageType = iota
	PostShared
	StoryShared
	OneTimeMessage
)


type Message struct {
	ID bool `json:"id"`
	Opened bool `json:"opened"`
	Type MessageType `json:"type"`
	ContentId string `json:"content_id,omitempty"`
	UserForContentId string `json:"user_for_content_id,omitempty"`
	Text string `json:"text,omitempty"`
	//Media string `json:"media,omitempty"`
	UserFrom string `json:"user_from,omitempty"`
	UserTo string `json:"user_to,omitempty"`
	Channel string `json:"channel,omitempty"`
	Content string `json:"content,omitempty"`
	Command int    `json:"command,omitempty"`
	Err     string `json:"err,omitempty"`
}

const (
	CommandSubscribe = iota
	CommandUnsubscribe
	CommandSendMessage
)