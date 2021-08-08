package Handler

import (
	"XWS-Nistagram/NotificationService/Model"
	"encoding/json"
	"fmt"
	"github.com/go-redis/redis/v7"
	"github.com/gorilla/mux"
	"github.com/gorilla/websocket"
	"net/http"
)

var upgrader = websocket.Upgrader{}


var connectedUsers = make(map[string]*Model.User)


func UserChannelsHandler(w http.ResponseWriter, r *http.Request, rdb *redis.Client) {

	vars := mux.Vars(r)
	username := vars["user"]

	list, err := Model.GetChannels(rdb, username)
	if err != nil {
		handleError(err, w)
		return
	}
	err = json.NewEncoder(w).Encode(list)
	if err != nil {
		handleError(err, w)
		return
	}

}

func UserChannelsNotificationsHandler(w http.ResponseWriter, r *http.Request, rdb *redis.Client) {

	w.Header().Set("Content-Type", "application/json")
	vars := mux.Vars(r)
	channel := vars["channel"]

	list, err := Model.GetChannelsNotifications(rdb, channel)
	if err != nil {
		handleError(err, w)
		return
	}
	var notifications []Model.NotificationForFollow
	for _, n := range list {
		var v Model.NotificationForFollow
		json.Unmarshal([]byte(n),&v)
		notifications = append(notifications, v)
	}
	err = json.NewEncoder(w).Encode(notifications)
	if err != nil {
		handleError(err, w)
		return
	}

}

func UserChannelsNotOpenedNotificationsHandler(w http.ResponseWriter, r *http.Request, rdb *redis.Client) {

	w.Header().Set("Content-Type", "application/json")
	vars := mux.Vars(r)
	channel := vars["channel"]

	list, err := Model.GetChannelsNotifications(rdb, channel)
	if err != nil {
		handleError(err, w)
		return
	}
	var notifications []Model.NotificationForFollow
	for _, n := range list {
		var v Model.NotificationForFollow
		json.Unmarshal([]byte(n),&v)
		if v.Opened == false {
			notifications = append(notifications, v)
		}
	}
	err = json.NewEncoder(w).Encode(notifications)
	if err != nil {
		handleError(err, w)
		return
	}

}

func UserChannelsNotificationsUpdateHandler(w http.ResponseWriter, r *http.Request, rdb *redis.Client) {

	w.Header().Set("Content-Type", "application/json")
	vars := mux.Vars(r)
	channel := vars["channel"]
	username := vars["user"]
	u := connectedUsers[username]


	var notification Model.NotificationForFollow
	err := json.NewDecoder(r.Body).Decode(&notification)
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	err = Model.UpdateNotification(rdb,channel,notification,u)
	if err != nil {
		handleError(err, w)
		return
	}
	if err != nil {
		handleError(err, w)
		return
	}

}

func UsersHandler(w http.ResponseWriter, r *http.Request, rdb *redis.Client) {

	list, err := Model.List(rdb)
	if err != nil {
		handleError(err, w)
		return
	}
	err = json.NewEncoder(w).Encode(list)
	if err != nil {
		handleError(err, w)
		return
	}
}

func handleError(err error, w http.ResponseWriter) {
	w.WriteHeader(http.StatusBadRequest)
	w.Write([]byte(fmt.Sprintf(`{"err": "%s"}`, err.Error())))
}

func H(rdb *redis.Client, fn func(http.ResponseWriter, *http.Request, *redis.Client)) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		fn(w, r, rdb)
	}
}



func ChatWebSocketHandler(w http.ResponseWriter, r *http.Request, rdb *redis.Client) {
	upgrader.CheckOrigin = func(r *http.Request) bool { return true }
	fmt.Println(upgrader)
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		handleWSError(err, conn)
		return
	}
	fmt.Println("upgrade")


	err = onConnect(r, conn, rdb)
	if err != nil {
		handleWSError(err, conn)
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
	username := vars["username"]
	fmt.Println("connected from:", conn.RemoteAddr(), "user:", username)

	u, err := Model.Connect(rdb, username)
	if err != nil {
		return err
	}
	connectedUsers[username] = u
	return nil
}

func onDisconnect(r *http.Request, conn *websocket.Conn, rdb *redis.Client) chan struct{} {

	closeCh := make(chan struct{})

	vars := mux.Vars(r)
	username := vars["username"]

	conn.SetCloseHandler(func(code int, text string) error {
		fmt.Println("connection closed for user", username)

		u := connectedUsers[username]
		if err := u.Disconnect(); err != nil {
			return err
		}
		delete(connectedUsers, username)
		close(closeCh)
		return nil
	})

	return closeCh
}

func onUserMessage(conn *websocket.Conn, r *http.Request, rdb *redis.Client) {

	var notificationForFollow Model.NotificationForFollow

	if err := conn.ReadJSON(&notificationForFollow); err != nil {
		handleWSError(err, conn)
		return
	}

	vars := mux.Vars(r)
	username := vars["username"]
	u := connectedUsers[username]

	switch notificationForFollow.Command {
	case Model.CommandSubscribe:
		if err := u.Subscribe(rdb, notificationForFollow.Channel); err != nil {
			handleWSError(err, conn)
		}
	case Model.CommandUnsubscribe:
		if err := u.Unsubscribe(rdb, notificationForFollow.Channel); err != nil {
			handleWSError(err, conn)
		}
	case Model.CommandSendNotification:
		if err := Model.SendNotification(rdb,notificationForFollow,u); err != nil {
			handleWSError(err, conn)
		}
	}
}

func onChannelMessage(conn *websocket.Conn, r *http.Request) {

	vars := mux.Vars(r)
	username := vars["username"]
	u := connectedUsers[username]

	go func() {
		for m := range u.MessageChan {

			msg := Model.NotificationForFollow{
				Content: m.Payload,
				Channel: m.Channel,
			}

			if err := conn.WriteJSON(msg); err != nil {
				fmt.Println(err)
			}
		}

	}()
}

func handleWSError(err error, conn *websocket.Conn) {
	_ = conn.WriteJSON(Model.NotificationForFollow{Err: err.Error()})
}

