package handler

import (
	"XWS-Nistagram/XWS-Nistagram/MessagesService/model"
	"XWS-Nistagram/XWS-Nistagram/MessagesService/service"
	"encoding/json"
	"fmt"
	"github.com/go-redis/redis/v7"
	"github.com/gorilla/mux"
	"github.com/gorilla/websocket"
	"log"
	"net/http"
)

var upgrader = websocket.Upgrader{}
var connectedUsers = make(map[string]*model.UserWS)


func ServeWS(w http.ResponseWriter, r *http.Request, rdb *redis.Client) {
	vars := mux.Vars(r)
	id := vars["id"]
	if id == "null"{
		return
	}
	upgrader.CheckOrigin = func(r *http.Request) bool { return true }
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		return
	}

	err = onConnect(r, conn, rdb)
	if err != nil {
		log.Println(err)
		return
	}

	closeCh := onDisconnect(r, conn, rdb)

	onChannelMessage(conn, r)

loop:
	for {
		select {
		case <-closeCh:
			break loop
		default:
			onUserMessage(conn, r, rdb)
		}
	}
}

func onConnect(r *http.Request, conn *websocket.Conn, rdb *redis.Client) error {
	vars := mux.Vars(r)
	id := vars["id"]
	fmt.Println("connected from:", conn.RemoteAddr(), "user:", id)

	u, err := service.Connect(rdb, id)
	if err != nil {
		return err
	}
	connectedUsers[id] = u
	return nil
}

func onDisconnect(r *http.Request, conn *websocket.Conn, rdb *redis.Client) chan struct{} {

	closeCh := make(chan struct{})

	vars := mux.Vars(r)
	id := vars["id"]

	conn.SetCloseHandler(func(code int, text string) error {
		fmt.Println("connection closed for user", id)

		u := connectedUsers[id]
		if err := service.Disconnect(u); err != nil {
			return err
		}
		delete(connectedUsers, id)
		close(closeCh)
		return nil
	})

	return closeCh
}

func onChannelMessage(conn *websocket.Conn, r *http.Request) {

	vars := mux.Vars(r)
	id := vars["id"]
	u := connectedUsers[id]

	go func() {
		for m := range u.MessageChan {

			msg := model.Message{
				Content: m.Payload,
				Channel: m.Channel,
			}

			if err := conn.WriteJSON(msg); err != nil {
				fmt.Println(err)
			}
		}

	}()
}

func onUserMessage(conn *websocket.Conn, r *http.Request, rdb *redis.Client) {
	fmt.Println("Primio poruku ili notifikaciju")

	var message model.Message

	if err := conn.ReadJSON(&message); err != nil {
			fmt.Println(err)
			return
	}

	vars := mux.Vars(r)
	id := vars["id"]
	u := connectedUsers[id]

	switch message.Command {
	case model.CommandSubscribe:
		if err := service.Subscribe(rdb, message.Channel,u); err != nil {
			fmt.Println(err)
		}
	case model.CommandUnsubscribe:
		if err := service.Unsubscribe(rdb, message.Channel,u); err != nil {
			fmt.Println(err)
		}
	case model.CommandSendMessage:
		if err := service.SendMessage(rdb,message,u); err != nil {
			fmt.Println(err)
		}
	case model.CommandSendNotification:
		fmt.Println("ajde vise")
		if err := service.SendNotification(rdb,message,u); err != nil {
			fmt.Println(err)
		}
	}
}

func GetAllMessageChatForUser(w http.ResponseWriter, r *http.Request, rdb *redis.Client) {

	vars := mux.Vars(r)
	id := vars["id"]
	fmt.Println("usao u GetAllMessageChatForUser")

	list, err := service.GetChannels(rdb, id)
	if err != nil {
		fmt.Println(err)
		return
	}
	err = json.NewEncoder(w).Encode(list)
	if err != nil {
		fmt.Println(err)
		return
	}

}
func ViewMessagesInChat(w http.ResponseWriter, r *http.Request, rdb *redis.Client) {


	w.Header().Set("Content-Type", "application/json")
	vars := mux.Vars(r)
	userid1 := vars["userid1"]
	userid2 := vars["userid2"]


	list, err := service.GetAllMessagesFromChat(rdb, userid1,userid2)
	if err != nil {
		fmt.Println(err)
		return
	}
	var notifications []model.Message
	for _, n := range list {
		var v model.Message
		json.Unmarshal([]byte(n),&v)
		notifications = append(notifications, v)
	}
	err = json.NewEncoder(w).Encode(notifications)
	if err != nil {
		fmt.Println(err)
		return
	}

}

func ViewAllNotificationsInChannel(w http.ResponseWriter, r *http.Request, rdb *redis.Client) {

	w.Header().Set("Content-Type", "application/json")
	vars := mux.Vars(r)
	channel := vars["channel"]

	list, err := service.GetAllNotificationsFromChannel(rdb, channel)
	if err != nil {
		fmt.Println(err)
		return
	}
	var notifications []model.Message
	for _, n := range list {
		var v model.Message
		json.Unmarshal([]byte(n),&v)
		notifications = append(notifications, v)
	}
	err = json.NewEncoder(w).Encode(notifications)
	if err != nil {
		fmt.Println(err)
		return
	}

}
func H(rdb *redis.Client, fn func(http.ResponseWriter, *http.Request, *redis.Client)) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		fn(w, r, rdb)
	}
}

