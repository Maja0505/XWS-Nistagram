package Model

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/go-redis/redis/v7"
	"strings"
)

const (
	usersKey = "users"
	userChannelFmt = "user:%s:channels"
	ChannelsKey = "channels"
	userChannelNotifications = "channel:%s:notification"
)


type User struct {
	name string
	channelsHandler *redis.PubSub
	stopListenerChan chan struct{}
	listening bool
	MessageChan chan redis.Message
}



func Connect(rdb *redis.Client, name string) (*User, error) {
	if _, err := rdb.SAdd(usersKey, name).Result(); err != nil {
		return nil, err
	}

	u := &User{
		name:             name,
		stopListenerChan: make(chan struct{}),
		MessageChan:      make(chan redis.Message),
	}

	if err := u.connect(rdb); err != nil {
		return nil, err
	}

	return u, nil
}

func (u *User) Subscribe(rdb *redis.Client, channel string) error {

	userChannelsKey := fmt.Sprintf(userChannelFmt, u.name)

	if rdb.SIsMember(userChannelsKey, channel).Val() {
		fmt.Println("postoji")
		return nil
	}
	fmt.Println("ne postoji")

	if err := rdb.SAdd(userChannelsKey, channel).Err(); err != nil {
		return err
	}
	if err := rdb.SAdd(ChannelsKey, channel).Err(); err != nil {
		return err
	}


	return u.connect(rdb)
}

func AddNotification(rdb *redis.Client, channel string, notification Message,u *User) error {
	fmt.Println("usao da upise")
	userChannelNotifications := fmt.Sprintf(userChannelNotifications,channel)
	fmt.Println(userChannelNotifications)
	n,_ := json.Marshal(notification)
	if rdb.SIsMember(userChannelNotifications, n).Val() {
		return nil
	}
	if err := rdb.SAdd(userChannelNotifications, n).Err(); err != nil {
		return err
	}
	fmt.Println("dodao")


	return u.connect(rdb)
}

func UpdateNotification(rdb *redis.Client, channel string, notification Message,u *User) error {
	fmt.Println("usao da izmeni")
	userChannelNotifications := fmt.Sprintf(userChannelNotifications,channel)
	fmt.Println(userChannelNotifications)
	n,_ := json.Marshal(notification)
	if rdb.SIsMember(userChannelNotifications, n).Val() {
		if err := rdb.SRem(userChannelNotifications, n).Err(); err != nil{
			return err
		}
		notification.Opened = true
		n,_ = json.Marshal(notification)
		if err := rdb.SAdd(userChannelNotifications, n).Err(); err != nil {
			return err
		}
	}

	fmt.Println("izmenio")


	return u.connect(rdb)
}

func (u *User) Unsubscribe(rdb *redis.Client, channel string) error {

	userChannelsKey := fmt.Sprintf(userChannelFmt, u.name)

	if !rdb.SIsMember(userChannelsKey, channel).Val() {
		return nil
	}
	if err := rdb.SRem(userChannelsKey, channel).Err(); err != nil {
		return err
	}

	return u.connect(rdb)
}


func (u *User) connect(rdb *redis.Client) error {

	var c []string

	c1, err := rdb.SMembers(ChannelsKey).Result()
	if err != nil {
		return err
	}
	c = append(c, c1...)

	// get all user channels (from DB) and start subscribe
	c2, err := rdb.SMembers(fmt.Sprintf(userChannelFmt, u.name)).Result()
	if err != nil {
		return err
	}
	c = append(c, c2...)

	if len(c) == 0 {
		fmt.Println("no channels to connect to for user: ", u.name)
		return nil
	}

	if u.channelsHandler != nil {
		if err := u.channelsHandler.Unsubscribe(); err != nil {
			return err
		}
		if err := u.channelsHandler.Close(); err != nil {
			return err
		}
	}
	if u.listening {
		u.stopListenerChan <- struct{}{}
	}

	return u.doConnect(rdb, c...)
}



func (u *User) doConnect(rdb *redis.Client, channels ...string) error {
	// subscribe all channels in one request
	pubSub := rdb.Subscribe(channels...)
	// keep channel handler to be used in unsubscribe
	u.channelsHandler = pubSub

	// The Listener
	go func() {
		u.listening = true
		fmt.Println("starting the listener for user:", u.name, "on channels:", channels)
		for {
			select {
			case msg, ok := <-pubSub.Channel():
				if !ok {
					return
				}
				u.MessageChan <- *msg

			case <-u.stopListenerChan:
				fmt.Println("stopping the listener for user:", u.name)
				return
			}
		}
	}()
	return nil
}

func GetChannels(rdb *redis.Client, userid string) ([]string, error) {

	if !rdb.SIsMember(usersKey, userid).Val() {
		return nil, errors.New("user not exists")
	}

	var c []string

	// get all user channels (from DB) and start subscribe
	c2, err := rdb.SMembers(fmt.Sprintf(userChannelFmt, userid)).Result()
	if err != nil {
		return nil, err
	}
	for _, s := range c2 {
		if strings.Contains(s,"-") {
			var ids = strings.Split(s,"-")
			for _, id := range ids {
				if id != userid {
					c = append(c, id)
				}
			}
		}
	}

	return c, nil
}

func GetChannelsNotifications(rdb *redis.Client, userid1 string,userid2 string) ([]string, error) {

	channel := userid1 + "-" + userid2
	if !rdb.SIsMember(ChannelsKey, channel).Val() {
		channel = userid2 + "-" + userid1
	}
	if !rdb.SIsMember(ChannelsKey, channel).Val() {
		return nil, errors.New("channel not exists")
	}

	var c []string

	// get all user channels (from DB) and start subscribe
	c1, err := rdb.SMembers(fmt.Sprintf(userChannelNotifications, channel)).Result()
	if err != nil {
		return nil, err
	}
	c = append(c, c1...)

	return c, nil
}

func List(rdb *redis.Client) ([]string, error) {
	return rdb.SMembers(usersKey).Result()
}

func (u *User) Disconnect() error {
	if u.channelsHandler != nil {
		if err := u.channelsHandler.Unsubscribe(); err != nil {
			return err
		}
		if err := u.channelsHandler.Close(); err != nil {
			return err
		}
	}
	if u.listening {
		u.stopListenerChan <- struct{}{}
	}

	close(u.MessageChan)

	return nil
}

func SendNotification(rdb *redis.Client,message Message,u *User) error {

	//reqUrl := fmt.Sprintf("http://" + os.Getenv("USER_SERVICE_DOMAIN") + ":" + os.Getenv("USER_SERVICE_PORT") + "/userid/" + follow.Channel)

	//resp, err := http.Get(reqUrl)
	/*if err != nil || resp.StatusCode == 404 {
		return err
	}
	body, err := ioutil.ReadAll(resp.Body)
	var user RegisteredUser
	err = json.Unmarshal(body, &user)

	*/
	messageJson,err := json.Marshal(message)
	if err != nil {
		return err
	}
	err = AddNotification(rdb,message.Channel,message,u)
	if err != nil {
		fmt.Println(err)
		return err
	}
	return rdb.Publish(message.Channel, messageJson).Err()
}


