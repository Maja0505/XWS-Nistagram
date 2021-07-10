package saga

import "encoding/json"

type Message struct {
	Service       string         `json:"service"`
	SenderService string         `json:"sender_service"`
	Action        string         `json:"action"`
	UserId		  string		  `json:"test"`

}


func (m Message) MarshalBinary() ([]byte, error) {
	return json.Marshal(m)
}