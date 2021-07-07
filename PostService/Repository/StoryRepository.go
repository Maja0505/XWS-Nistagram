package Repository

import (
	"XWS-Nistagram/PostService/Model"
	"fmt"
	"github.com/gocql/gocql"
	"time"
)

type StoryRepository struct {
	Session gocql.Session
}

func (repo *StoryRepository) Create(story *Model.Story) error {
	ID := gocql.TimeUUID()
	if err := repo.Session.Query("INSERT INTO postkeyspace.stories(id, userid, available, image, for_close_friends, highlights) VALUES(?, ?, ?, ?, ?, ?)",
		ID, story.UserID, true, story.Image,story.ForCloseFriends,story.Highlights).Exec(); err != nil {
		return err
	}
	if err := repo.Session.Query("UPDATE postkeyspace.stories USING TTL 86400 SET available = True WHERE id = ? AND userid = ?",
		ID, story.UserID).Exec(); err != nil {
		return err
	}
	fmt.Println("Successfully created story!")

	return nil
}

func (repo *StoryRepository) SetStoryForHighlights(id gocql.UUID) error {
	err := repo.Session.Query("update postkeyspace.stories set highlights = True where id=?;",id).Exec()

	if err != nil {
		return err
	}
	fmt.Println("Successfully created highlight!")

	return nil
}

func (repo *StoryRepository) GetAllStoriesByUser(userId string) (*[]Model.Story,error){
	var stories []Model.Story
	m := map[string]interface{}{}
	query := "select * from postkeyspace.stories where userid=?;"
	iter := repo.Session.Query(query,userId).Iter()

	for iter.MapScan(m){
		a := m["id"].(gocql.UUID)
		b := a.Time()
		stories = append(stories, Model.Story{
			ID: m["id"].(gocql.UUID),
			UserID: m["userid"].(string),
			CreatedAt: b,
			Available: m["available"].(bool),
			Image: m["image"].(string),
			Highlights: m["highlights"].(bool),
			ForCloseFriends: m["for_close_friends"].(bool),
		})
		m = map[string]interface{}{}
	}

	return &stories,nil

}

func (repo *StoryRepository) GetAllNotExpiredStoriesByUser(userId string) (*[]Model.Story,error){
	var stories []Model.Story
	m := map[string]interface{}{}
	query := "select * from postkeyspace.stories where userid=? and available = true and for_close_friends=False allow filtering;"
	iter := repo.Session.Query(query,userId).Iter()

	for iter.MapScan(m){
		a := m["id"].(gocql.UUID)
		b := a.Time()
		stories = append(stories, Model.Story{
			ID: m["id"].(gocql.UUID),
			UserID: m["userid"].(string),
			CreatedAt: b,
			Available: m["available"].(bool),
			Image: m["image"].(string),
			Highlights: m["highlights"].(bool),
			ForCloseFriends: m["for_close_friends"].(bool),
		})
		m = map[string]interface{}{}
	}

	return &stories,nil

}

func (repo *StoryRepository) GetAllStoriesForCloseFriendsByUser(userId string) (*[]Model.Story,error){
	var stories []Model.Story
	m := map[string]interface{}{}
	query := "select * from postkeyspace.stories where userid=? and available = true allow filtering;"
	iter := repo.Session.Query(query,userId).Iter()

	for iter.MapScan(m){
		a := m["id"].(gocql.UUID)
		b := a.Time()
		stories = append(stories, Model.Story{
			ID: m["id"].(gocql.UUID),
			UserID: m["userid"].(string),
			CreatedAt: b,
			Available: m["available"].(bool),
			Image: m["image"].(string),
			Highlights: m["highlights"].(bool),
			ForCloseFriends: m["for_close_friends"].(bool),
		})
		m = map[string]interface{}{}
	}

	return &stories,nil

}

func (repo *StoryRepository) GetAllHighlightsStoriesByUser(userId string) (*[]Model.Story,error){
	var stories []Model.Story
	m := map[string]interface{}{}
	query := "select * from postkeyspace.stories where userid=? and highlights=True allow filtering;"
	iter := repo.Session.Query(query,userId).Iter()

	for iter.MapScan(m){
		a := m["id"].(gocql.UUID)
		b := a.Time()
		stories = append(stories, Model.Story{
			ID: m["id"].(gocql.UUID),
			UserID: m["userid"].(string),
			CreatedAt: b,
			Available: m["available"].(bool),
			Image: m["image"].(string),
			Highlights: m["highlights"].(bool),
			ForCloseFriends: m["for_close_friends"].(bool),
		})
		m = map[string]interface{}{}
	}

	return &stories,nil

}

func (repo *StoryRepository) CheckDoesUserHaveAnyNotExpiredStory(userId string) bool{
	now := time.Now()
	query := "select * from postkeyspace.stories where userid=? and for_close_friends=False and available = true allow filtering;"
	iter := repo.Session.Query(query,userId,now).Iter()
	for iter.MapScan(map[string]interface{}{}){
		return true
	}
	return false
}

func (repo *StoryRepository) CheckDoesUserHaveAnyNotExpiredStoryForCloseFriends(userId string) bool{
	now := time.Now()
	query := "select * from postkeyspace.stories where userid=? and available = true  allow filtering;"
	iter := repo.Session.Query(query,userId,now).Iter()
	for iter.MapScan(map[string]interface{}{}){
		return true
	}
	return false
}








