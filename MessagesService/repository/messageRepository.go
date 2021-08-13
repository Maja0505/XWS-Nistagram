package repository

import (
	"XWS-Nistagram/XWS-Nistagram/MessagesService/model"
	"encoding/json"
	"fmt"
	"github.com/go-redis/redis/v7"
)



func Connect(rdb *redis.Client, id string) (*model.UserWS, error) {
	fmt.Println("1")
	rdb = redis.NewClient(&redis.Options{Addr: "redis:6379"})

	if _, err := rdb.SAdd(model.UsersKey, id).Result(); err != nil {
		return nil, err
	}
	u := &model.UserWS{
		Name:             id,
		StopListenerChan: make(chan struct{}),
		MessageChan:      make(chan redis.Message),
	}

	if err := connect(rdb,u); err != nil {
		return nil, err
	}

	return u, nil
}

func  Disconnect(u *model.UserWS) error {
	if u.ChannelsHandler != nil {
		if err := u.ChannelsHandler.Unsubscribe(); err != nil {
			return err
		}
		if err := u.ChannelsHandler.Close(); err != nil {
			return err
		}
	}
	if u.Listening {
		u.StopListenerChan <- struct{}{}
	}

	close(u.MessageChan)

	return nil
}

func connect(rdb *redis.Client,u *model.UserWS) error {
	fmt.Println("2")

	var c []string

	c1, err := rdb.SMembers(model.ChannelsKey).Result()
	if err != nil {
		return err
	}
	c = append(c, c1...)

	// get all user channels (from DB) and start subscribe
	c2, err := rdb.SMembers(fmt.Sprintf(model.UserChannelFmt, u.Name)).Result()
	if err != nil {
		return err
	}
	c = append(c, c2...)

	if len(c) == 0 {
		fmt.Println("no channels to connect to for user: ", u.Name)
		return nil
	}

	if u.ChannelsHandler != nil {
		if err := u.ChannelsHandler.Unsubscribe(); err != nil {
			return err
		}
		if err := u.ChannelsHandler.Close(); err != nil {
			return err
		}
	}
	if u.Listening {
		u.StopListenerChan <- struct{}{}
	}

	return doConnect(rdb,u, c...)
}

func  doConnect(rdb *redis.Client,u *model.UserWS,channels ...string) error {
	// subscribe all channels in one request
	fmt.Println("3")

	pubSub := rdb.Subscribe(channels...)
	// keep channel handler to be used in unsubscribe
	u.ChannelsHandler = pubSub

	// The Listener
	go func() {
		u.Listening = true
		fmt.Println("starting the listener for user:", u.Name, "on channels:", channels)
		for {
			select {
			case msg, ok := <-pubSub.Channel():
				if !ok {
					return
				}
				u.MessageChan <- *msg

			case <-u.StopListenerChan:
				fmt.Println("stopping the listener for user:", u.Name)
				return
			}
		}
	}()
	return nil
}

func Subscribe(rdb *redis.Client, channel string,u *model.UserWS) error {
	rdb = redis.NewClient(&redis.Options{Addr: "redis:6379"})
	userChannelsKey := fmt.Sprintf(model.UserChannelFmt, u.Name)

	if rdb.SIsMember(userChannelsKey, channel).Val() {
		fmt.Println("postoji")
		return nil
	}
	fmt.Println("ne postoji")

	if err := rdb.SAdd(userChannelsKey, channel).Err(); err != nil {
		return err
	}
	/*if err := rdb.SAdd(model.ChannelsKey, channel).Err(); err != nil {
		return err
	}*/


	return connect(rdb,u)
}

func Unsubscribe(rdb *redis.Client, channel string,u *model.UserWS) error {

	userChannelsKey := fmt.Sprintf(model.UserChannelFmt, u.Name)

	if !rdb.SIsMember(userChannelsKey, channel).Val() {
		return nil
	}
	if err := rdb.SRem(userChannelsKey, channel).Err(); err != nil {
		return err
	}

	return connect(rdb,u)
}

func addNotification(rdb *redis.Client, channel string, notification model.Message,u *model.UserWS) error {
	fmt.Println("usao da upise")
	userChannelNotifications := fmt.Sprintf(model.UserChannelMessage,channel)
	fmt.Println(userChannelNotifications)
	n,_ := json.Marshal(notification)
	if rdb.SIsMember(userChannelNotifications, n).Val() {
		return nil
	}
	if err := rdb.SAdd(userChannelNotifications, n).Err(); err != nil {
		return err
	}
	fmt.Println("dodao")


	return connect(rdb,u)
}

func SendMessage(rdb *redis.Client,message model.Message,u *model.UserWS) error {

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
	err = addNotification(rdb,message.Channel,message,u)
	if err != nil {
		fmt.Println(err)
		return err
	}
	return rdb.Publish(message.Channel, messageJson).Err()
}

