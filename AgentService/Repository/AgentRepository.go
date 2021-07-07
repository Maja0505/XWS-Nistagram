package Repository

import (
	"XWS-Nistagram/AgentService/DTO"
	"XWS-Nistagram/AgentService/Model"
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/gocql/gocql"
	"io/ioutil"
	"net/http"
	"os"
	"time"
)

type AgentRepository struct {
	Session gocql.Session
}

func (repo *AgentRepository) CreateTables() error {

	if err := repo.Session.Query("DROP TABLE IF EXISTS agentkeyspace.campaigns;").Exec(); err != nil {
		fmt.Println("Error while dropping tables!")
		fmt.Println(err)
		return err
	}

	if err := repo.Session.Query("CREATE TABLE if not exists agentkeyspace.campaigns(id timeuuid, userid text, ispost boolean, repeat boolean, start timestamp, end timestamp, repeatfactor int, media list<text>, links list<text>, PRIMARY KEY((userid), id)) WITH CLUSTERING ORDER BY (id DESC);").Exec(); err != nil {
		fmt.Println("Error while creating tables!")
		fmt.Println(err)
		return err
	}

	return nil
}
func (repo *AgentRepository) CreateCampaign(campaign *Model.Campaign) error {
	ID := gocql.TimeUUID()
	if err := repo.Session.Query("INSERT INTO agentkeyspace.campaigns(id, userid, ispost, repeat, repeatfactor, start, end, media, links) VALUES(?, ?, ?, ?, ?, ?, ?, ?, ?)",
		ID, campaign.UserID, campaign.IsPost, campaign.Repeat, campaign.RepeatFactor, campaign.Start, campaign.End, campaign.Media, campaign.Links).Exec(); err != nil {
		fmt.Println("Error while creating campaign!")
		fmt.Println(err)
		return err
	}
	fmt.Println("Campaign created successfully!")
	if campaign.Repeat == false {
		repo.StartOneTimeTimer(campaign.Start, campaign)
	}else{

	}
	return nil
}

func (repo *AgentRepository) DeleteCampaign(campaign *Model.Campaign) error {
	if err := repo.Session.Query("DELETE FROM agentkeyspace.campaigns where id = ? AND userid = ? IF EXISTS;",
		campaign.ID, campaign.UserID).Exec(); err != nil {
		fmt.Println("Error while deleting campaign!")
		fmt.Println(err)
		return err
	}
	fmt.Println("Campaign deleted successfully!")
	return nil
}

func (repo *AgentRepository) StartTestTimeTimer(in time.Time, campaign *Model.Campaign) {
	timerdur, err := repo.CalculateTimerDuration(in)
	fmt.Println(in)
	fmt.Println(timerdur)
	if err != nil{
		fmt.Println(err)
	}
	time.AfterFunc(timerdur, func() {
		fmt.Println("Funkcijaaaaaa timer")
	})

}

func (repo *AgentRepository) StartOneTimeTimer(in time.Time, campaign *Model.Campaign) error {
	timerdur, err := repo.CalculateTimerDuration(in)
	if err != nil{
		fmt.Println(err)
		return err
	}
	time.AfterFunc(timerdur, func() {

		var isAlbum = false
		if len(campaign.Media) == 1{
			isAlbum = false
		}else{
			isAlbum = true
		}
		reqUrl := fmt.Sprintf("http://" + os.Getenv("POST_SERVICE_DOMAIN") + ":" + os.Getenv("POST_SERVICE_PORT") + "/create")
		Postdto := DTO.PostDTO{
			Description: "Agent",
			Media: campaign.Media,
			UserID: campaign.UserID,
			MediaCount: int64(len(campaign.Media)),
			Album: isAlbum,
		}
		jsonPost,_ := json.Marshal(Postdto)

		resp, err := http.Post(reqUrl,"appliation/json",bytes.NewBuffer(jsonPost))
		if err != nil || resp.StatusCode == 404 {
			fmt.Println(err)
		}
		body, err := ioutil.ReadAll(resp.Body)
		fmt.Println(body)
		// Printed after stated duration
		// by AfterFunc() method is over

		// loop stops at this point
	})
	return nil
}

func (repo *AgentRepository) CalculateTimerDuration(in time.Time) (time.Duration, error) {
	fmt.Println("A ",time.Now())
	in = in.Add(-2*time.Hour)
	fmt.Println("C ",in.UTC())
	now := time.Now()
	ret := in.Sub(now)

	fmt.Println("A ",ret)
	return ret, nil
}
func (repo *AgentRepository) CalculateRepeatTimerDuration(start time.Time, end time.Time, repeatfactor int) (time.Duration, error) {
	fmt.Println("A ",time.Now())
	start = start.Add(-2*time.Hour)
	end = end.Add(-2*time.Hour)
	time.

	fmt.Println("START: ",start)
	fmt.Println("END: ",end)
	now := time.Now()
	ret := in.Sub(now)

	fmt.Println("A ",ret)
	return ret, nil
}

func (repo *AgentRepository) CreatePost(campaign *Model.Campaign) error {
	if campaign == nil {
		return gocql.Error{Message: "Campaign is null!"}
	}
	var isAlbum = false
	if len(campaign.Media) == 1{
		isAlbum = false
	}else{
		isAlbum = true
	}
	reqUrl := fmt.Sprintf("http://" + os.Getenv("POST_SERVICE_DOMAIN") + ":" + os.Getenv("POST_SERVICE_PORT") + "/create")
	Postdto := DTO.PostDTO{
		Description: "Agent",
		Media: campaign.Media,
		UserID: campaign.UserID,
		MediaCount: int64(len(campaign.Media)),
		Album: isAlbum,
	}
	jsonPost,_ := json.Marshal(Postdto)

	resp, err := http.Post(reqUrl,"appliation/json",bytes.NewBuffer(jsonPost))
	if err != nil || resp.StatusCode == 404 {
		return err
	}
	body, err := ioutil.ReadAll(resp.Body)
	fmt.Println(body)
	return  nil

}
