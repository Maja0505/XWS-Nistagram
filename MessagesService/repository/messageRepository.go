package repository

import (
	"XWS-Nistagram/XWS-Nistagram/MessagesService/model"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/go-redis/redis/v7"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
)



func Connect(rdb *redis.Client, id string) (*model.UserWS, error) {
	fmt.Println("usao u Connect ")

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
	fmt.Println("izasao u Connect ")

	return u, nil
}

func  Disconnect(u *model.UserWS) error {
	fmt.Println("usao u Disconnect ")

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
	fmt.Println("izasoa u Disconnect ")

	return nil
}

func connect(rdb *redis.Client,u *model.UserWS) error {
	fmt.Println("usao u connect ")

	var c []string

	Subscribe(rdb,u.Name,u)

	/*c1, err := rdb.SMembers(model.ChannelsKey).Result()
	if err != nil {
		return err
	}
	c = append(c, c1...)*/

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
	fmt.Println("izasao u connect ")


	return doConnect(rdb,u, c...)
}

func  doConnect(rdb *redis.Client,u *model.UserWS,channels ...string) error {
	// subscribe all channels in one request
	fmt.Println("usao u doConnect ")

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
	fmt.Println("izasao u doConnect ")

	return nil
}

func Subscribe(rdb *redis.Client, channel string,u *model.UserWS) error {
	fmt.Println("usao u Subscribe ")

	rdb = redis.NewClient(&redis.Options{Addr: "redis:6379"})
	userChannelsKey := fmt.Sprintf(model.UserChannelFmt, u.Name)

	if strings.Contains(channel,"-"){
		ids := strings.Split(channel,"-")
		if rdb.SIsMember(userChannelsKey,  ids[0] + "-" + ids[1]).Val() {
			fmt.Println("postoji")
			return nil
		}
		if rdb.SIsMember(userChannelsKey,  ids[1] + "-" + ids[0]).Val() {
			fmt.Println("postoji")
			return nil
		}
	}else{
		if rdb.SIsMember(userChannelsKey, channel).Val() {
			fmt.Println("postoji")
			return nil
		}
	}

	fmt.Println("ne postoji")

	if err := rdb.SAdd(userChannelsKey, channel).Err(); err != nil {
		return err
	}
	if err := rdb.SAdd(model.ChannelsKey, channel).Err(); err != nil {
		return err
	}

	fmt.Println("izasoa u Subscribe ")

	return connect(rdb,u)
}

func Unsubscribe(rdb *redis.Client, channel string,u *model.UserWS) error {
	fmt.Println("usao u Unsubscribe ")

	userChannelsKey := fmt.Sprintf(model.UserChannelFmt, u.Name)

	if !rdb.SIsMember(userChannelsKey, channel).Val() {
		return nil
	}
	if err := rdb.SRem(userChannelsKey, channel).Err(); err != nil {
		return err
	}
	fmt.Println("izasao u Unsubscribe ")

	return connect(rdb,u)
}

func addMessage(rdb *redis.Client, channel string, notification model.Message,u *model.UserWS) error {
	fmt.Println("usao u addMessage")
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
	fmt.Println("izasao iz addMessage")


	return connect(rdb,u)
}

func SendMessage(rdb *redis.Client,message model.Message,u *model.UserWS) error {
	fmt.Println("usao u SendMessage")

	//reqUrl := fmt.Sprintf("http://" + os.Getenv("USER_SERVICE_DOMAIN") + ":" + os.Getenv("USER_SERVICE_PORT") + "/userid/" + follow.Channel)

	//resp, err := http.Get(reqUrl)
	/*if err != nil || resp.StatusCode == 404 {
		return err
	}
	body, err := ioutil.ReadAll(resp.Body)
	var user RegisteredUser
	err = json.Unmarshal(body, &user)

	*/
	fmt.Println("usao da upise poruku")
	rdb = redis.NewClient(&redis.Options{Addr: "redis:6379"})
	if strings.Contains(message.Channel,"-"){
		ids := strings.Split(message.Channel,"-")
		u1 := &model.UserWS{
			Name: ids[0],
		}
		u2 := &model.UserWS{
			Name: ids[1],
		}
		Subscribe(rdb,message.Channel,u1)
		Subscribe(rdb,message.Channel,u2)

	}
	messageJson,err := json.Marshal(message)
	if err != nil {
		return err
	}
	err = addMessage(rdb,message.Channel,message,u)
	if err != nil {
		fmt.Println(err)
		return err
	}
	notification := model.Message{
		UserFrom: message.UserFrom,
		Channel: message.UserIdTo,
		Command: 3,
		Content: "sent you a message.",
		Media: "",
		PostId: "",

	}
	SendNotification(rdb, notification,u)
	fmt.Println("izasao iz SendMessage")

	return rdb.Publish(message.Channel, messageJson).Err()
}

func GetChannels(rdb *redis.Client, userid string) ([]string, error) {
	rdb = redis.NewClient(&redis.Options{Addr: "redis:6379"})

	fmt.Println("usao u GetChannels")
 	fmt.Println(userid)
	if !rdb.SIsMember(model.UsersKey, userid).Val() {
		return nil, errors.New("user not exists")
	}

	var c []string

	// get all user channels (from DB) and start subscribe
	c2, err := rdb.SMembers(fmt.Sprintf(model.UserChannelFmt, userid)).Result()
	if err != nil {
		return nil, err
	}
	fmt.Println("dobavio sve kanale koje user prati")

	for _, s := range c2 {
		if strings.Contains(s,"-") {
			var ids = strings.Split(s,"-")
			for _, id := range ids {
				fmt.Println("pronasao -")
				if id != userid {
					c = append(c, id)
				}
			}
		}
	}
	fmt.Println("izasao u GetChannels")

	return c, nil
}

func GetAllMessagesFromChat(rdb *redis.Client, userid1 string,userid2 string) ([]string, error) {
	fmt.Println("usao u GetAllMessagesFromChat")

	rdb = redis.NewClient(&redis.Options{Addr: "redis:6379"})
	userChannelsKey := fmt.Sprintf(model.UserChannelFmt, userid1)

	channel := userid1 + "-" + userid2
	fmt.Println(channel)

	if !rdb.SIsMember(userChannelsKey, channel).Val() {
		channel = userid2 + "-" + userid1
	}
	fmt.Println(channel)
	if !rdb.SIsMember(userChannelsKey, channel).Val() {
		return nil, errors.New("channel not exists")
	}

	var c []string

	// get all user channels (from DB) and start subscribe
	c1, err := rdb.SMembers(fmt.Sprintf(model.UserChannelMessage, channel)).Result()
	if err != nil {
		return nil, err
	}
	c = append(c, c1...)
	fmt.Println("izasao u GetAllMessagesFromChat")

	return c, nil
}

func AddNotification(rdb *redis.Client, channel string, notification model.Message,u *model.UserWS) error {
	fmt.Println("usao u AddNotification")

	fmt.Println("usao da upise notifikaciju")
	userChannelNotifications := fmt.Sprintf(model.UserChannelNotification,channel)
	fmt.Println(userChannelNotifications)
	n,_ := json.Marshal(notification)
	if rdb.SIsMember(userChannelNotifications, n).Val() {
		return nil
	}
	if err := rdb.SAdd(userChannelNotifications, n).Err(); err != nil {
		return err
	}
	fmt.Println("dodao")
	fmt.Println("izasao u AddNotification")

	return connect(rdb,u)
}

func SendNotification(rdb *redis.Client,notification model.Message,u *model.UserWS) error {
	fmt.Println("usao u SendNotification")

	rdb = redis.NewClient(&redis.Options{Addr: "redis:6379"})
	fmt.Println(notification.Channel)
	reqUrl := fmt.Sprintf("http://" + os.Getenv("USER_SERVICE_DOMAIN") + ":" + os.Getenv("USER_SERVICE_PORT") + "/userid/" + notification.Channel)

	resp, err := http.Get(reqUrl)
	if err != nil || resp.StatusCode == 404 {
		fmt.Println("ovde je greska")
		fmt.Println(err)
		fmt.Println(resp.StatusCode)
		return err
	}
	body, err := ioutil.ReadAll(resp.Body)
	var user model.RegisteredUser
	err = json.Unmarshal(body, &user)

	fmt.Println("dobavio registrovanog usera")

	fmt.Println(user.ProfileSettings.Public)

	fmt.Println(user.NotificationSettings.FollowNotification)
	fmt.Println(user.NotificationSettings.FollowRequestNotification)

	if notification.Content == "liked your photo." || notification.Content == "disliked your photo." {
		if !user.NotificationSettings.LikeNotification {
			return errors.New("false")
		}
	}
	if notification.Content == "commented your post:" {
		if !user.NotificationSettings.CommentNotification {

			return errors.New("false")
		}
	}

	if notification.Content == "requested to following you." {
		if !user.NotificationSettings.FollowRequestNotification {

			return errors.New("false")
		}
	}

	if notification.Content == "started following you." {
		if !user.NotificationSettings.FollowNotification {

			return errors.New("false")
		}
	}
	fmt.Println(notification.Content)
	ntf,err := json.Marshal(notification)
	if err != nil {
		return err
	}
	err = AddNotification(rdb,notification.Channel,notification,u)
	if err != nil {
		fmt.Println(err)
		return err
	}
	fmt.Println("izasao iz SendNotification")

	return rdb.Publish(notification.Channel, ntf).Err()
}

func GetAllNotificationsFromChannel(rdb *redis.Client, channel string) ([]string, error) {
	fmt.Println("usao u GetAllNotificationsFromChannel")

	rdb = redis.NewClient(&redis.Options{Addr: "redis:6379"})

	if !rdb.SIsMember(model.ChannelsKey, channel).Val() {
		return nil, errors.New("channel not exists")
	}

	var c []string

	// get all user channels (from DB) and start subscribe
	c1, err := rdb.SMembers(fmt.Sprintf(model.UserChannelNotification, channel)).Result()
	if err != nil {
		return nil, err
	}
	c = append(c, c1...)
	fmt.Println("izasao iz GetAllNotificationsFromChannel")

	return c, nil
}

