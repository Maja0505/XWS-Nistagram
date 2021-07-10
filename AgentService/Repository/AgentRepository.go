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

	/*if err := repo.Session.Query("DROP TABLE IF EXISTS agentkeyspace.campaigns;").Exec(); err != nil {
		fmt.Println("Error while dropping tables!")
		fmt.Println(err)
		return err
	}*/

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
	campaign.ID = ID
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

func (repo *AgentRepository) CreateCampaignRequest(userid string, campaignid gocql.UUID) error {
	if err := repo.Session.Query("INSERT INTO agentkeyspace.campaignrequest(campaignid, userid) VALUES(?, ?)",
		campaignid, userid).Exec(); err != nil {
		fmt.Println("Error while creating campaign request!")
		fmt.Println(err)
		return err
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

		_, err := repo.CreateFullPost(campaign)
		if err != nil{
			fmt.Println(err)
		}
		fmt.Println("Prosappp")
		err = repo.CreatePostsForInfluencers(campaign)
		if err != nil{
			fmt.Println(err)
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
		err := repo.CreateAgentStoriesOnce(campaign)
		if err != nil{
			fmt.Println(err)
		}
		err = repo.CreateInfluencerStoriesOnce(campaign)
		if err != nil{
			fmt.Println(err)
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

		uuid, err := repo.CreateFullPost(campaign)
		if err != nil{
			fmt.Println(err)
		}
		err = repo.CreatePostsForInfluencers(campaign)
		if err != nil{
			fmt.Println(err)
		}
		updateDto := DTO.UpdateCreatedAtDTO{
			ID: uuid,
			UserID: campaign.UserID,
			CreatedAt: time.Now(),
		}
		go doEvery(tick, campaign.End, &updateDto, updateCreatedAt)
	})
	fmt.Println("Successfully created influencer and agent posts!!!")
	return nil
}

func (repo *AgentRepository) StartRepeatTimerStory(in time.Time, campaign *Model.Campaign, tick time.Duration) error {
	timerdur, err := repo.CalculateTimerDuration(in)
	if err != nil{
		fmt.Println(err)
		return err
	}
	time.AfterFunc(timerdur, func() {

		if err := repo.CreateAgentStories(campaign, tick); err != nil{
			fmt.Println(err)
		}
		if err := repo.CreateInfluencerStories(campaign, tick); err != nil{
			fmt.Println(err)
		}
	})
	fmt.Println("Successfully created influencer and agent stories!!!")
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

func (repo *AgentRepository) CreateAgentPost(campaign *Model.Campaign) (gocql.UUID,error) {

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
		IsCampaign: true,
	}
	jsonPost,_ := json.Marshal(Postdto)

	resp, err := http.Post(reqUrl,"appliation/json",bytes.NewBuffer(jsonPost))
	if err != nil || resp.StatusCode == 404 {
		fmt.Println(err)
		return gocql.UUID{}, err
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil || resp.StatusCode == 404 {
		fmt.Println(err)
		return gocql.UUID{}, err
	}
	var a string
	json.Unmarshal(body, &a)
	fmt.Println(a)
	uuid, err := ParseUUID(a)
	if err != nil{
		fmt.Println(err)
		return gocql.UUID{}, err
	}
	return uuid, nil
}

func (repo *AgentRepository) CreateInfluencerPost(campaign *Model.Campaign, infid *string) (gocql.UUID,error) {

	fmt.Println("CreateInfluencerPost")
	var isAlbum = false
	if len(campaign.Media) == 1 {
		isAlbum = false
	} else {
		isAlbum = true
	}
	reqUrl := fmt.Sprintf("http://" + os.Getenv("POST_SERVICE_DOMAIN") + ":" + os.Getenv("POST_SERVICE_PORT") + "/create")
	Postdto := DTO.PostDTO{
		Description: campaign.Description,
		Media:       campaign.Media,
		UserID:      *infid,
		MediaCount:  int64(len(campaign.Media)),
		Album:       isAlbum,
		Location:    campaign.Location,
		IsCampaign: true,
	}
	jsonPost, _ := json.Marshal(Postdto)

	resp, err := http.Post(reqUrl, "appliation/json", bytes.NewBuffer(jsonPost))
	if err != nil || resp.StatusCode == 404 {
		fmt.Println(err)
		return gocql.UUID{}, err
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil || resp.StatusCode == 404 {
		fmt.Println(err)
		return gocql.UUID{}, err
	}
	var a string
	json.Unmarshal(body, &a)
	fmt.Println(a)
	uuid, err := ParseUUID(a)
	if err != nil {
		fmt.Println(err)
		return gocql.UUID{}, err
	}
	return uuid, nil
}

func (repo *AgentRepository) CreateLinks(userid *string, links *[]string, uuid *gocql.UUID) error {
	reqUrl := fmt.Sprintf("http://" + os.Getenv("POST_SERVICE_DOMAIN") + ":" + os.Getenv("POST_SERVICE_PORT") + "/add-links")
	addLinksDTO := DTO.UpdateLinksDTO{
		ID:     *uuid,
		UserID: *userid,
		Links:  *links,
	}

	jsonaddLinks, _ := json.Marshal(addLinksDTO)

	resp, err := http.Post(reqUrl, "appliation/json", bytes.NewBuffer(jsonaddLinks))
	if err != nil || resp.StatusCode == 404 {
		fmt.Println(err)
		return err
	}
	_, err = ioutil.ReadAll(resp.Body)
	if err != nil || resp.StatusCode == 404 {
		fmt.Println(err)
		return err
	}
	return nil
}
func (repo *AgentRepository) CreateTags(tags []string, uuid *gocql.UUID) error {
	for i := range tags {
		reqUrl := fmt.Sprintf("http://" + os.Getenv("POST_SERVICE_DOMAIN") + ":" + os.Getenv("POST_SERVICE_PORT") + "/add-tag")
		tag := Model.Tag{
			Tag:    tags[i],
			PostID: *uuid,
		}
		jsonTag, _ := json.Marshal(tag)

		resp, err := http.Post(reqUrl, "appliation/json", bytes.NewBuffer(jsonTag))
		if err != nil || resp.StatusCode == 404 {
			fmt.Println(err)
			return err
		}
		_, err = ioutil.ReadAll(resp.Body)
		if err != nil {
			fmt.Println(err)
			return err
		}
	}
	return nil
}
func (repo *AgentRepository) CreateFullPost(campaign *Model.Campaign) (gocql.UUID, error) {
	uuid, err := repo.CreateAgentPost(campaign)
	if err != nil {
		fmt.Println(err)
		return gocql.UUID{}, err
	}

	err = repo.CreateLinks(&campaign.UserID, &campaign.Links, &uuid)
	if err != nil {
		fmt.Println(err)
		return gocql.UUID{}, err
	}

	repo.CreateTags(campaign.Tags, &uuid)
	if err != nil {
		fmt.Println(err)
		return gocql.UUID{}, err
	}
	return uuid, nil
}

func (repo *AgentRepository) CreateAgentStory(campaign *Model.Campaign, i int) (gocql.UUID, error){

	reqUrl := fmt.Sprintf("http://" + os.Getenv("POST_SERVICE_DOMAIN") + ":" + os.Getenv("POST_SERVICE_PORT") + "/story/create")

	fmt.Println(campaign.Media[i])
	storyDTO := DTO.StoryDTO{
		Media:           campaign.Media[i],
		UserID:          campaign.UserID,
		ForCloseFriends: false,
		Highlights:      false,
		Link: campaign.Links[i],
	}
	jsonPost, _ := json.Marshal(storyDTO)
	fmt.Println(storyDTO.Media)
	fmt.Println(storyDTO.Link)

	resp, err := http.Post(reqUrl, "appliation/json", bytes.NewBuffer(jsonPost))
	if err != nil || resp.StatusCode == 404 {
		fmt.Println(err)
		return gocql.UUID{}, err
	}
	body, err := ioutil.ReadAll(resp.Body)
	var a string
	json.Unmarshal(body, &a)
	fmt.Println(a)
	uuid, err := ParseUUID(a)
	if err != nil {
		fmt.Println(err)
		return gocql.UUID{}, err
	}
	return uuid, nil
}

func (repo *AgentRepository) CreateInfluencerStory(campaign *Model.Campaign, i int, infid string) (gocql.UUID, error){

	reqUrl := fmt.Sprintf("http://" + os.Getenv("POST_SERVICE_DOMAIN") + ":" + os.Getenv("POST_SERVICE_PORT") + "/story/create")

	storyDTO := DTO.StoryDTO{
		Media:           campaign.Media[i],
		UserID:          infid,
		ForCloseFriends: false,
		Highlights:      false,
	}
	jsonPost, _ := json.Marshal(storyDTO)

	resp, err := http.Post(reqUrl, "appliation/json", bytes.NewBuffer(jsonPost))
	if err != nil || resp.StatusCode == 404 {
		fmt.Println(err)
		return gocql.UUID{}, err
	}
	body, err := ioutil.ReadAll(resp.Body)
	var a string
	json.Unmarshal(body, &a)
	fmt.Println(a)
	uuid, err := ParseUUID(a)
	if err != nil {
		fmt.Println(err)
		return gocql.UUID{}, err
	}
	return uuid, nil
}

func (repo *AgentRepository) CreateAgentStories(campaign *Model.Campaign, tick time.Duration) error {
	for i := range campaign.Media {
		uuid, err := repo.CreateAgentStory(campaign, i)
		if err != nil {
			fmt.Println(err)
			return err
		}
		updto := DTO.UpdateCreatedAtDTO{
			ID:        uuid,
			UserID:    campaign.UserID,
			CreatedAt: time.Now(),
		}
		go doEvery(tick, campaign.End, &updto, updateStoryCreatedAt)
	}
	return nil
}

func (repo *AgentRepository) CreateAgentStoriesOnce(campaign *Model.Campaign) error {
	for i := range campaign.Media {
		_, err := repo.CreateAgentStory(campaign, i)
		if err != nil {
			fmt.Println(err)
			return err
		}
	}
	return nil
}

func (repo *AgentRepository) CreateInfluencerStories(campaign *Model.Campaign, tick time.Duration) error {
	a, err := repo.GetCampaignInfluencers(campaign.ID)
	if err != nil {
		fmt.Println(err)
		return err
	}
	for _, s := range a {
		for i := range campaign.Media {
			uuid, err := repo.CreateInfluencerStory(campaign, i, s)
			if err != nil {
				fmt.Println(err)
				return err
			}
			updto := DTO.UpdateCreatedAtDTO{
				ID:        uuid,
				UserID:    s,
				CreatedAt: time.Now(),
			}
			go doEvery(tick, campaign.End, &updto, updateStoryCreatedAt)
		}
	}
	return nil
}

func (repo *AgentRepository) CreateInfluencerStoriesOnce(campaign *Model.Campaign) error {
	a, err := repo.GetCampaignInfluencers(campaign.ID)
	if err != nil {
		fmt.Println(err)
		return err
	}
	for _, s := range a {
		for i := range campaign.Media {
			_, err := repo.CreateInfluencerStory(campaign, i, s)
			if err != nil {
				fmt.Println(err)
				return err
			}
		}
	}
	return nil
}

func (repo *AgentRepository) CreatePostsForInfluencers(campaign *Model.Campaign) error {
	a, err := repo.GetCampaignInfluencers(campaign.ID)
	if err != nil {
		fmt.Println(err)
		return err
	}
	for _, s := range a {
		fmt.Println(s)
		uuid, err := repo.CreateInfluencerPost(campaign, &s)
		if err != nil {
			fmt.Println(err)
			return err
		}

		err = repo.CreateLinks(&s, &campaign.Links, &uuid)
		if err != nil {
			fmt.Println(err)
			return err
		}

		repo.CreateTags(campaign.Tags, &uuid)
		if err != nil {
			fmt.Println(err)
			return err
		}
	}
	return nil
}

func doEvery(d time.Duration, end time.Time, dto *DTO.UpdateCreatedAtDTO, f func(time.Time, *DTO.UpdateCreatedAtDTO)) {
	end = end.Add(-2*time.Hour)
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

func (repo *AgentRepository) AddCampaignInfluencer(infid string, userid string, id gocql.UUID) error{
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

func (repo *AgentRepository) GetCampaignInfluencers(campaignid gocql.UUID) ([]string, error){

	fmt.Println("CC ", campaignid)
	m := map[string]interface{}{}
	iter := repo.Session.Query("SELECT * FROM agentkeyspace.campaigns WHERE id = ? ALLOW FILTERING",
		campaignid).Iter();
	//if iter.NumRows() == 0{
	//	return nil, gocql.Error{Message: "Campaign not found!"}
	//}
	iter.MapScan(m)
	var a = m["influencers"].([]string)

	fmt.Println("Successfully loaded influencers from campaign!")
	return a, nil
}

func (repo *AgentRepository) GetCampaignsForUser(userid string) ( *[]Model.Campaign, error){
	var campaigns []Model.Campaign
	m := map[string]interface{}{}

	iter := repo.Session.Query("SELECT * FROM agentkeyspace.campaigns WHERE userid=?", userid).Iter()
	fmt.Println(iter.NumRows())
	if iter.NumRows() == 0{
		return nil, nil
	}
	for iter.MapScan(m) {
		var a = m["end"].(time.Time)
		if a.Sub(time.Now()) > 0 {
			var a = m["id"].(gocql.UUID)
			b := a.Time()
			var campaign = Model.Campaign{
				ID:           m["id"].(gocql.UUID),
				CreatedAt:    b,
				//Description:  m["description"].(string),
				UserID:       m["userid"].(string),
				Media:        m["media"].([]string),
				Links:        m["links"].([]string),
				Repeat:       m["repeat"].(bool),
				IsPost:       m["ispost"].(bool),
				Start:        m["start"].(time.Time),
				End:          m["end"].(time.Time),
				RepeatFactor: m["repeatfactor"].(int),
				Influencers:  m["influencers"].([]string),
			}
			campaigns = append(campaigns, campaign)
			m = map[string]interface{}{}
		}
	}
	return &campaigns, nil
}

func (repo *AgentRepository) GetCampaignRequests(userid string) ( *[]DTO.RequestDTO, error){
	var campaignRequests []DTO.RequestDTO
	m := map[string]interface{}{}

	iter := repo.Session.Query("SELECT * FROM agentkeyspace.campaignrequest WHERE userid=?", userid).Iter()
	if iter.NumRows() == 0{
		return nil, nil
	}
	for iter.MapScan(m) {
			var campaign = DTO.RequestDTO{
				CampaignID:   m["campaignid"].(gocql.UUID),
				UserID:       m["userid"].(string),
			}
		campaignRequests = append(campaignRequests, campaign)
		m = map[string]interface{}{}
	}
	return &campaignRequests, nil
}

func (repo *AgentRepository) GetCampaignById(campaignidString string) ( *Model.Campaign, error){

	campaignid,err := ParseUUID(campaignidString)
	if err != nil {
		return nil, err
	}

	var campaignFirst Model.Campaign
	m := map[string]interface{}{}

	iter := repo.Session.Query("SELECT * FROM agentkeyspace.campaigns WHERE id = ? ALLOW FILTERING;",
		campaignid).Iter()
	fmt.Println(iter.NumRows())
	if iter.NumRows() == 0{
		return nil, nil
	}
	for iter.MapScan(m) {
		var a = m["start"].(time.Time)
		if a.Sub(time.Now()) > 0 {
			var a = m["id"].(gocql.UUID)
			b := a.Time()
			var campaign = Model.Campaign{
				ID:           m["id"].(gocql.UUID),
				CreatedAt:    b,
				//Description:  m["description"].(string),
				UserID:       m["userid"].(string),
				Media:        m["media"].([]string),
				Links:        m["links"].([]string),
				Repeat:       m["repeat"].(bool),
				IsPost:       m["ispost"].(bool),
				Start:        m["start"].(time.Time),
				End:          m["end"].(time.Time),
				RepeatFactor: m["repeatfactor"].(int),
				Influencers:  m["influencers"].([]string),
			}
			campaignFirst = campaign
			return &campaignFirst, nil
		}
	}
	return nil, nil
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