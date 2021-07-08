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

	if err := repo.Session.Query("CREATE TABLE if not exists agentkeyspace.campaigns(id timeuuid, userid text, ispost boolean, repeat boolean, start timestamp, end timestamp, repeatfactor int, media list<text>, links list<text>, influencers list<text>, PRIMARY KEY((userid), id)) WITH CLUSTERING ORDER BY (id DESC);").Exec(); err != nil {
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
	if campaign.IsPost == true{
		if campaign.Repeat == false {
			repo.StartOneTimeTimer(campaign.Start, campaign)
		}else{
			tick, err := repo.CalculateRepeatTimerTick(campaign.RepeatFactor)
			if err != nil{
				fmt.Println(err)
				return err
			}
			repo.StartRepeatTimer(campaign.Start, campaign, tick)
		}
	}else{
		if campaign.Repeat == false {
			repo.StartOneTimeTimerStory(campaign.Start, campaign)
		}else{
			tick, err := repo.CalculateRepeatTimerTick(campaign.RepeatFactor)
			if err != nil{
				fmt.Println(err)
				return err
			}
			repo.StartRepeatTimerStory(campaign.Start, campaign, tick)
		}
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

func (repo *AgentRepository) StartOneTimeTimer(in time.Time, campaign *Model.Campaign) error {
	timerdur, err := repo.CalculateTimerDuration(in)
	if err != nil{
		fmt.Println(err)
		return err
	}
	go time.AfterFunc(timerdur, func() {

		var isAlbum = false
		if len(campaign.Media) == 1{
			isAlbum = false
		}else{
			isAlbum = true
		}
		reqUrl := fmt.Sprintf("http://" + os.Getenv("POST_SERVICE_DOMAIN") + ":" + os.Getenv("POST_SERVICE_PORT") + "/create")
		Postdto := DTO.PostDTO{
			Description: campaign.Description,
			Media: campaign.Media,
			UserID: campaign.UserID,
			MediaCount: int64(len(campaign.Media)),
			Album: isAlbum,
			Location: campaign.Location,
		}
		jsonPost,_ := json.Marshal(Postdto)

		resp, err := http.Post(reqUrl,"appliation/json",bytes.NewBuffer(jsonPost))
		if err != nil || resp.StatusCode == 404 {
			fmt.Println(err)
		}
		body, err := ioutil.ReadAll(resp.Body)
		var a string
		json.Unmarshal(body, &a)
		fmt.Println(a)
		uuid, err := ParseUUID(a)
		if err != nil{
			fmt.Println(err)
		}

		reqUrl = fmt.Sprintf("http://" + os.Getenv("POST_SERVICE_DOMAIN") + ":" + os.Getenv("POST_SERVICE_PORT") + "/add-links")
		addLinksDTO := DTO.UpdateLinksDTO{
			ID: uuid,
			UserID: campaign.UserID,
			Links: campaign.Links,
		}

		jsonaddLinks,_ := json.Marshal(addLinksDTO)

		resp, err = http.Post(reqUrl,"appliation/json",bytes.NewBuffer(jsonaddLinks))
		if err != nil || resp.StatusCode == 404 {
			fmt.Println(err)
		}
		body, err = ioutil.ReadAll(resp.Body)

		for i := range campaign.Tags {
			reqUrl = fmt.Sprintf("http://" + os.Getenv("POST_SERVICE_DOMAIN") + ":" + os.Getenv("POST_SERVICE_PORT") + "/add-tag")
			tag := Model.Tag{
				Tag: campaign.Tags[i],
				PostID: uuid,
			}
			jsonTag, _ := json.Marshal(tag)

			resp, err = http.Post(reqUrl, "appliation/json", bytes.NewBuffer(jsonTag))
			if err != nil || resp.StatusCode == 404 {
				fmt.Println(err)
			}
			body, err = ioutil.ReadAll(resp.Body)
		}
	})
	return nil
}

func (repo *AgentRepository) StartOneTimeTimerStory(in time.Time, campaign *Model.Campaign) error {
	timerdur, err := repo.CalculateTimerDuration(in)
	if err != nil{
		fmt.Println(err)
		return err
	}
	go time.AfterFunc(timerdur, func() {

		for _,s := range campaign.Media {
			reqUrl := fmt.Sprintf("http://" + os.Getenv("POST_SERVICE_DOMAIN") + ":" + os.Getenv("POST_SERVICE_PORT") + "/story/create")

			storyDTO := DTO.StoryDTO{
				Media:      s,
				UserID:     campaign.UserID,
				ForCloseFriends: false,
				Highlights: false,
			}
			jsonPost, _ := json.Marshal(storyDTO)

			resp, err := http.Post(reqUrl, "appliation/json", bytes.NewBuffer(jsonPost))
			if err != nil || resp.StatusCode == 404 {
				fmt.Println(err)
			}
			body, err := ioutil.ReadAll(resp.Body)
			fmt.Println(body)
		}
	})
	return nil
}

func (repo *AgentRepository) StartRepeatTimer(in time.Time, campaign *Model.Campaign, tick time.Duration) error {
	timerdur, err := repo.CalculateTimerDuration(in)
	if err != nil{
		fmt.Println(err)
		return err
	}
	go time.AfterFunc(timerdur, func() {

		var isAlbum = false
		if len(campaign.Media) == 1{
			isAlbum = false
		}else{
			isAlbum = true
		}
		reqUrl := fmt.Sprintf("http://" + os.Getenv("POST_SERVICE_DOMAIN") + ":" + os.Getenv("POST_SERVICE_PORT") + "/create")
		Postdto := DTO.PostDTO{
			Description: campaign.Description,
			Media: campaign.Media,
			UserID: campaign.UserID,
			MediaCount: int64(len(campaign.Media)),
			Album: isAlbum,
			Location: campaign.Location,
		}
		jsonPost,_ := json.Marshal(Postdto)

		resp, err := http.Post(reqUrl,"appliation/json",bytes.NewBuffer(jsonPost))
		if err != nil || resp.StatusCode == 404 {
			fmt.Println(err)
		}
		body, err := ioutil.ReadAll(resp.Body)


		var a string
		json.Unmarshal(body, &a)
		fmt.Println(a)
		uuid, err := ParseUUID(a)
		if err != nil{
			fmt.Println(err)
		}
		reqUrl = fmt.Sprintf("http://" + os.Getenv("POST_SERVICE_DOMAIN") + ":" + os.Getenv("POST_SERVICE_PORT") + "/add-links")
		addLinksDTO := DTO.UpdateLinksDTO{
			ID: uuid,
			UserID: campaign.UserID,
			Links: campaign.Links,
		}

		jsonaddLinks,_ := json.Marshal(addLinksDTO)

		resp, err = http.Post(reqUrl,"appliation/json",bytes.NewBuffer(jsonaddLinks))
		if err != nil || resp.StatusCode == 404 {
			fmt.Println(err)
		}
		body, err = ioutil.ReadAll(resp.Body)

		updto := DTO.UpdateCreatedAtDTO{
			ID: uuid,
			UserID: campaign.UserID,
			CreatedAt: time.Now(),
		}

		for i := range campaign.Tags {
			reqUrl = fmt.Sprintf("http://" + os.Getenv("POST_SERVICE_DOMAIN") + ":" + os.Getenv("POST_SERVICE_PORT") + "/add-tag")
			tag := Model.Tag{
				Tag: campaign.Tags[i],
				PostID: uuid,
			}
			jsonTag, _ := json.Marshal(tag)

			resp, err = http.Post(reqUrl, "appliation/json", bytes.NewBuffer(jsonTag))
			if err != nil || resp.StatusCode == 404 {
				fmt.Println(err)
			}
			body, err = ioutil.ReadAll(resp.Body)
		}

		go doEvery(tick, campaign.End, &updto, updateCreatedAt)
	})
	return nil
}

func (repo *AgentRepository) StartRepeatTimerStory(in time.Time, campaign *Model.Campaign, tick time.Duration) error {
	timerdur, err := repo.CalculateTimerDuration(in)
	if err != nil{
		fmt.Println(err)
		return err
	}
	time.AfterFunc(timerdur, func() {

		for _,s := range campaign.Media {
			reqUrl := fmt.Sprintf("http://" + os.Getenv("POST_SERVICE_DOMAIN") + ":" + os.Getenv("POST_SERVICE_PORT") + "/story/create")

			storyDTO := DTO.StoryDTO{
				Media:           s,
				UserID:          campaign.UserID,
				ForCloseFriends: false,
				Highlights:      false,
			}
			jsonPost, _ := json.Marshal(storyDTO)

			resp, err := http.Post(reqUrl, "appliation/json", bytes.NewBuffer(jsonPost))
			if err != nil || resp.StatusCode == 404 {
				fmt.Println(err)
			}
			body, err := ioutil.ReadAll(resp.Body)
			var a string
			json.Unmarshal(body, &a)
			fmt.Println(a)
			uuid, err := ParseUUID(a)
			if err != nil {
				fmt.Println(err)
			}

			updto := DTO.UpdateCreatedAtDTO{
				ID:     uuid,
				UserID: campaign.UserID,
				CreatedAt: time.Now(),
			}

			go doEvery(tick, campaign.End, &updto, updateStoryCreatedAt)
		}
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
func (repo *AgentRepository) CalculateRepeatTimerTick(repeatfactor int) (time.Duration, error) {
	day := time.Duration(time.Hour*24)

	fmt.Println("Day: ",day)
	sec := day.Seconds()/float64(repeatfactor)
	fmt.Println("Day/repeatfactor: ",sec)

	ret := time.Duration(sec*1000000000)
	fmt.Println("Timer tick duration: ", ret)
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

func doEvery(d time.Duration, end time.Time, dto *DTO.UpdateCreatedAtDTO, f func(time.Time, *DTO.UpdateCreatedAtDTO)) {
	//end = end.Add(-2*time.Hour)
	for x := range time.Tick(d) {
		a := end.Sub(time.Now())
		fmt.Println("Razlika ", a)
		dto.CreatedAt = time.Now()
		if a > 0{
			f(x, dto)
		}else{
			break
		}
	}
}

func updateCreatedAt(t time.Time, dto *DTO.UpdateCreatedAtDTO) {

	reqUrl := fmt.Sprintf("http://" + os.Getenv("POST_SERVICE_DOMAIN") + ":" + os.Getenv("POST_SERVICE_PORT") + "/update-createdat")
	jsonPost,_ := json.Marshal(dto)

	resp, err := http.Post(reqUrl,"appliation/json",bytes.NewBuffer(jsonPost))
	if err != nil || resp.StatusCode == 404 {
		fmt.Println("ERROR")
	}
	body, err := ioutil.ReadAll(resp.Body)
	fmt.Println(body)
	fmt.Println("Successfully updated createdat")
}

func updateStoryCreatedAt(t time.Time, dto *DTO.UpdateCreatedAtDTO) {

	reqUrl := fmt.Sprintf("http://" + os.Getenv("POST_SERVICE_DOMAIN") + ":" + os.Getenv("POST_SERVICE_PORT") + "/story/update-agent")
	jsonPost,_ := json.Marshal(dto)

	resp, err := http.Post(reqUrl,"appliation/json",bytes.NewBuffer(jsonPost))
	if err != nil || resp.StatusCode == 404 {
		fmt.Println("ERROR")
	}
	body, err := ioutil.ReadAll(resp.Body)
	fmt.Println(body)
	fmt.Println("Successfully updated createdat")
}

func (repo *AgentRepository) AddCampaignInfluencer(infid string, userid string, id gocql.UUID, start time.Time) error{
	if start.Sub(time.Now()) < 0{
		return gocql.Error{Message: "Campaign has already started!"}
	}
	a := []string{infid}

	if err := repo.Session.Query("UPDATE agentkeyspace.campaigns SET influencers = ? + influencers WHERE id = ? AND userid = ?",
		a, id, userid).Exec(); err != nil {
		fmt.Println("Error while adding influencer to campaign!")
		fmt.Println(err)
		return err
	}
	fmt.Println("Successfully added influencer to campaign!")
	return nil
}


func ParseUUID(input string) (gocql.UUID, error) {
	var u gocql.UUID
	j := 0
	for _, r := range input {
		switch {
		case r == '-' && j&1 == 0:
			continue
		case r >= '0' && r <= '9' && j < 32:
			u[j/2] |= byte(r-'0') << uint(4-j&1*4)
		case r >= 'a' && r <= 'f' && j < 32:
			u[j/2] |= byte(r-'a'+10) << uint(4-j&1*4)
		case r >= 'A' && r <= 'F' && j < 32:
			u[j/2] |= byte(r-'A'+10) << uint(4-j&1*4)
		default:
			return gocql.UUID{}, fmt.Errorf("invalid UUID %q", input)
		}
		j += 1
	}
	if j != 32 {
		return gocql.UUID{}, fmt.Errorf("invalid UUID %q", input)
	}
	return u, nil
}