package Repository

import (
	"XWS-Nistagram/PostService/Model"
	"github.com/gocql/gocql"
	"time"
)

type StoryRepository struct {
	Session gocql.Session
}

func (repo *StoryRepository) Create(story *Model.Story) error {
	ID, _ := gocql.RandomUUID()
	createdAt := time.Now()
	expiredAt := createdAt.Add(time.Hour * time.Duration(24))
	if err := repo.Session.Query("INSERT INTO postkeyspace.stories(id,userid, createdat, expiredat , image , for_close_friends,highlights) VALUES(?, ?, ?, ?, ?, ?, ?)",
		ID, story.UserID, createdAt, expiredAt, story.Image,story.ForCloseFriends,story.Highlights).Exec(); err != nil {
		return err
	}

	return nil
}

func (repo *StoryRepository) SetStoryForHighlights(id gocql.UUID) error {
	err := repo.Session.Query("update postkeyspace.stories set highlights = True where id=?;",id).Exec()

	if err != nil {
		return err
	}

	return nil
}

func (repo *StoryRepository) GetAllStoriesByUser(userId string) (*[]Model.Story,error){
	var stories []Model.Story
	m := map[string]interface{}{}
	query := "select * from postkeyspace.stories where userid=? allow filtering;"
	iter := repo.Session.Query(query,userId).Iter()

	for iter.MapScan(m){
		stories = append(stories, Model.Story{
			ID: m["id"].(gocql.UUID),
			UserID: m["userid"].(string),
			CreatedAt: m["createdat"].(time.Time),
			ExpiredAt: m["expiredat"].(time.Time),
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
	now := time.Now()
	query := "select * from postkeyspace.stories where userid=? and for_close_friends=False and expiredat > ?  allow filtering;"
	iter := repo.Session.Query(query,userId,now).Iter()

	for iter.MapScan(m){
		stories = append(stories, Model.Story{
			ID: m["id"].(gocql.UUID),
			UserID: m["userid"].(string),
			CreatedAt: m["createdat"].(time.Time),
			ExpiredAt: m["expiredat"].(time.Time),
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
	now := time.Now()
	query := "select * from postkeyspace.stories where userid=? and expiredat > ?  allow filtering;"
	iter := repo.Session.Query(query,userId,now).Iter()

	for iter.MapScan(m){
		stories = append(stories, Model.Story{
			ID: m["id"].(gocql.UUID),
			UserID: m["userid"].(string),
			CreatedAt: m["createdat"].(time.Time),
			ExpiredAt: m["expiredat"].(time.Time),
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
	query := "select * from postkeyspace.stories where userid=? and highlights=True  allow filtering;"
	iter := repo.Session.Query(query,userId).Iter()

	for iter.MapScan(m){
		stories = append(stories, Model.Story{
			ID: m["id"].(gocql.UUID),
			UserID: m["userid"].(string),
			CreatedAt: m["createdat"].(time.Time),
			ExpiredAt: m["expiredat"].(time.Time),
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
	query := "select * from postkeyspace.stories where userid=? and for_close_friends=False and expiredat > ? allow filtering;"
	iter := repo.Session.Query(query,userId,now).Iter()
	for iter.MapScan(map[string]interface{}{}){
		return true
	}
	return false
}

func (repo *StoryRepository) CheckDoesUserHaveAnyNotExpiredStoryForCloseFriends(userId string) bool{
	now := time.Now()
	query := "select * from postkeyspace.stories where userid=? and expiredat > ? allow filtering;;"
	iter := repo.Session.Query(query,userId,now).Iter()
	for iter.MapScan(map[string]interface{}{}){
		return true
	}
	return false
}








