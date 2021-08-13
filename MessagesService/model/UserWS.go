package model

import "github.com/go-redis/redis/v7"

const (
	UsersKey = "users"
	UserChannelFmt = "user:%s:channels"
	ChannelsKey = "channels"
	UserChannelMessage = "channel:%s:message"
)


type UserWS struct {
	Name             string
	ChannelsHandler  *redis.PubSub
	StopListenerChan chan struct{}
	Listening        bool
	MessageChan      chan redis.Message
}
