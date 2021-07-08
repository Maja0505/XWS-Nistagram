package saga


import "encoding/json"

const (
	UserChannel    			string = "UserChannel"
	UserFollowerChannel    	string = "UserFollowerChannel"
	ReplyChannel    		string = "ReplyChannel"
	ServiceUser    			string = "User"
	ServiceUserFollower    	string = "UserFollower"
	ActionStart     		string = "Start"
	ActionDone      		string = "DoneMsg"
	ActionError     		string = "ErrorMsg"
	ActionRollback  		string = "RollbackMsg"
)

type Message struct {
	Service       string         `json:"service"`
	SenderService string         `json:"sender_service"`
	Action        string         `json:"action"`
	UserId		  string		  `json:"test"`

}

func (m Message) MarshalBinary() ([]byte, error) {
	return json.Marshal(m)
}