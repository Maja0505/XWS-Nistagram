package Repository

import (
	"XWS-Nistagram/PostService/Model"
	"fmt"
	"github.com/gocql/gocql"
	"image"
	"os"
	"time"
)

type PostRepository struct {
	Session gocql.Session
}

func (repo *PostRepository) CreateTables() error{
	if err := repo.Session.Query("DROP TABLE IF EXISTS postkeyspace.posts;").Exec(); err != nil {
		fmt.Println("Error while dropping tables!")
		fmt.Println(err)
		return err
	}
	if err := repo.Session.Query("DROP TABLE IF EXISTS postkeyspace.postcounters;").Exec(); err != nil {
		fmt.Println("Error while dropping tables!")
		fmt.Println(err)
		return err
	}
	if err := repo.Session.Query("DROP TABLE IF EXISTS postkeyspace.comments;").Exec(); err != nil {
		fmt.Println("Error while dropping tables!")
		fmt.Println(err)
		return err
	}
	if err := repo.Session.Query("DROP TABLE IF EXISTS postkeyspace.likes;").Exec(); err != nil {
		fmt.Println("Error while dropping tables!")
		fmt.Println(err)
		return err
	}
	if err := repo.Session.Query("DROP TABLE IF EXISTS postkeyspace.dislikes;").Exec(); err != nil {
		fmt.Println("Error while dropping tables!")
		fmt.Println(err)
		return err
	}
	if err := repo.Session.Query("DROP TABLE IF EXISTS postkeyspace.tags;").Exec(); err != nil {
		fmt.Println("Error while dropping tables!")
		fmt.Println(err)
		return err
	}
	if err := repo.Session.Query("CREATE TABLE postkeyspace.posts(id uuid, userid text, createdat timestamp, description text, image text, location text, PRIMARY KEY((userid, id)));").Exec(); err != nil {
		fmt.Println("Error while creating tables!")
		fmt.Println(err)
		return err
	}
	if err := repo.Session.Query("CREATE TABLE postkeyspace.postcounters(postid uuid, likes counter, dislikes counter, comments counter, PRIMARY KEY(postid));").Exec(); err != nil {
		fmt.Println("Error while creating tables!")
		fmt.Println(err)
		return err
	}
	if err := repo.Session.Query("CREATE TABLE postkeyspace.comments(id uuid, postid uuid, userid text, createdat timestamp, content text, PRIMARY KEY((postid), userid, id));").Exec(); err != nil {
		fmt.Println("Error while creating tables!")
		fmt.Println(err)
		return err
	}
	if err := repo.Session.Query("CREATE TABLE postkeyspace.likes(postid uuid, userid text, PRIMARY KEY((postid, userid)));").Exec(); err != nil {
		fmt.Println("Error while creating tables!")
		fmt.Println(err)
		return err
	}
	if err := repo.Session.Query("CREATE TABLE postkeyspace.dislikes(postid uuid, userid text, PRIMARY KEY((postid, userid)));").Exec(); err != nil {
		fmt.Println("Error while creating tables!")
		fmt.Println(err)
		return err
	}
	if err := repo.Session.Query("CREATE TABLE postkeyspace.tags(postid uuid, tag text, PRIMARY KEY((postid, tag)));").Exec(); err != nil {
		fmt.Println("Error while creating tables!")
		fmt.Println(err)
		return err
	}
	fmt.Println("Successfully dropped and created tables!!")
	return nil
}

func (repo *PostRepository) Create(post *Model.Post) error {
	ID, _ := gocql.RandomUUID()
	if err := repo.Session.Query("INSERT INTO postkeyspace.posts(id, createdat, description, image, userid) VALUES(?, ?, ?, ?, ?)",
		ID, post.CreatedAt, post.Description, post.Image, post.UserID).Exec(); err != nil {
		fmt.Println("Error while creating post!")
		fmt.Println(err)
		return err
	}
	fmt.Println("Successfully created post!!")
	return nil
}

func (repo *PostRepository) AddComment(comment *Model.Comment) error {
	fmt.Println("aa")
	ID, _ := gocql.RandomUUID()
	if err := repo.Session.Query("INSERT INTO postkeyspace.comments(id, postid, userid, createdat, content) VALUES(?, ?, ?, ?, ?)",
		ID, comment.PostID, comment.UserID, comment.CreatedAt, comment.Content).Exec(); err != nil {
		fmt.Println("Error while creating comment!")
		fmt.Println(err)
		return err
	}
	if err := repo.IncrementComments(comment); err != nil {
		return err
	}
	fmt.Println("Successfully created comment!!")
	return nil
}

func (repo *PostRepository) DeleteComment(comment *Model.Comment) error {
	if err := repo.Session.Query("DELETE FROM postkeyspace.comments where id = ? AND postid = ? AND userid = ? IF EXISTS;",
		comment.ID, comment.PostID, comment.UserID).Exec(); err != nil {
		fmt.Println("Error while deleting comment!")
		fmt.Println(err)
		return err
	}
	if err := repo.DecrementComments(comment); err != nil {
		return err
	}
	fmt.Println("Successfully deleted comment!!")
	return nil
}

func (repo *PostRepository) LikePost(like *Model.Like) error {

	var dislike Model.Dislike
	dislike.PostID = like.PostID
	dislike.UserID = like.UserID

	var liked bool
	var disliked bool
	liked = repo.CheckIfLikeExists(like)
	fmt.Println("Liked: ", liked)
	if liked {
		err := repo.DeleteLike(like)
		if err != nil{
			return err
		}
		if err := repo.DecrementLikes(like); err != nil{
			return err
		}
		return nil
	}
	disliked = repo.CheckIfDislikeExists(&dislike)
	fmt.Println("Disliked: ", disliked)

	if err := repo.Session.Query("INSERT INTO postkeyspace.likes(postid, userid) VALUES(?, ?)",
		like.PostID, like.UserID).Exec(); err != nil {
		fmt.Println("Error while creating like!")
		fmt.Println(err)
		return err
	}

	if err := repo.DeleteDislike(&dislike); err != nil {
		return err
	}
	fmt.Println("Liked: ", liked)
	fmt.Println("Disliked: ", disliked)
	if liked == false && disliked == false {
		if err := repo.IncrementLikes(like); err != nil{
			return err
		}
	}else if liked == true && disliked == false{
	}else if liked == false && disliked == true{
		if err := repo.IncrementLikes(like); err != nil{
			return err
		}
		if err := repo.DecrementDislikes(&dislike); err != nil{
			return err
		}
	}

	fmt.Println("Successfully created like!!")
	return nil
}

func (repo *PostRepository) DislikePost(dislike *Model.Dislike) error {

	var like Model.Like
	like.PostID = dislike.PostID
	like.UserID = dislike.UserID

	var liked bool
	var disliked bool
	liked = repo.CheckIfLikeExists(&like)
	fmt.Println("Liked: ", liked)
	disliked = repo.CheckIfDislikeExists(dislike)
	fmt.Println("Disliked: ", disliked)
	if disliked {
		err := repo.DeleteDislike(dislike)
		if err != nil{
			return err
		}
		if err := repo.DecrementDislikes(dislike); err != nil{
			return err
		}

		return nil
	}

	if err := repo.Session.Query("INSERT INTO postkeyspace.dislikes(postid, userid) VALUES(?, ?)",
		dislike.PostID, dislike.UserID).Exec(); err != nil {
		fmt.Println("Error while creating dislike!")
		fmt.Println(err)
		return err
	}

	if err := repo.DeleteLike(&like); err != nil {
		return err
	}

	if liked == false && disliked == false {
		if err := repo.IncrementDislikes(dislike); err != nil{
			return err
		}
	}else if liked == false && disliked == true{
	}else if liked == true && disliked == false{
		if err := repo.IncrementDislikes(dislike); err != nil{
			return err
		}
		if err := repo.DecrementLikes(&like); err != nil{
			return err
		}
	}

	fmt.Println("Successfully created dislike!!")
	return nil
}

func (repo *PostRepository) IncrementLikes(like *Model.Like) error{
	if err := repo.Session.Query("UPDATE postkeyspace.postcounters SET likes = likes + 1 WHERE postid = ?",
		like.PostID).Exec(); err != nil {
		fmt.Println("Error updating like/dislike counter!")
		fmt.Println(err)
		return err
	}
	return nil
}

func (repo *PostRepository) IncrementDislikes(dislike *Model.Dislike) error{
	if err := repo.Session.Query("UPDATE postkeyspace.postcounters SET dislikes = dislikes + 1 WHERE postid = ?",
		dislike.PostID).Exec(); err != nil {
		fmt.Println("Error updating like/dislike counter!")
		fmt.Println(err)
		return err
	}
	return nil
}

func (repo *PostRepository) DecrementLikes(like *Model.Like) error{
	if err := repo.Session.Query("UPDATE postkeyspace.postcounters SET likes = likes - 1 WHERE postid = ?",
		like.PostID).Exec(); err != nil {
		fmt.Println("Error updating like/dislike counter!")
		fmt.Println(err)
		return err
	}
	return nil
}

func (repo *PostRepository) DecrementDislikes(dislike *Model.Dislike) error{
	if err := repo.Session.Query("UPDATE postkeyspace.postcounters SET dislikes = dislikes - 1 WHERE postid = ?",
		dislike.PostID).Exec(); err != nil {
		fmt.Println("Error updating like/dislike counter!")
		fmt.Println(err)
		return err
	}
	return nil
}
func (repo *PostRepository) IncrementComments(comment *Model.Comment) error {
	if err := repo.Session.Query("UPDATE postkeyspace.postcounters SET comments = comments + 1 WHERE postid = ?",
		comment.PostID).Exec(); err != nil {
		fmt.Println("Error updating comments counter!")
		fmt.Println(err)
		return err
	}
	return nil
}
func (repo *PostRepository) DecrementComments(comment *Model.Comment) error {
	if err := repo.Session.Query("UPDATE postkeyspace.postcounters SET comments = comments - 1 WHERE postid = ?",
		comment.PostID).Exec(); err != nil {
		fmt.Println("Error updating comments counter!")
		fmt.Println(err)
		return err
	}
	return nil
}
func (repo *PostRepository) CheckIfLikeExists(like *Model.Like) bool {
	var likes int
	repo.Session.Query("SELECT COUNT(*) FROM postkeyspace.likes WHERE postid = ? AND userid = ? LIMIT 1",
		like.PostID,like.UserID).Iter().Scan(&likes)
	if likes == 1 {
		return true
	}else{
		return false
	}

}
func (repo *PostRepository) CheckIfDislikeExists(dislike *Model.Dislike) bool {
	var dislikes int
	repo.Session.Query("SELECT COUNT(*) FROM postkeyspace.dislikes WHERE postid = ? AND userid = ? LIMIT 1",
		dislike.PostID, dislike.UserID).Iter().Scan(&dislikes)
	if dislikes == 1 {
		return true
	}else{
		return false
	}

}

func (repo *PostRepository) DeleteLike(like *Model.Like) error {
	if err := repo.Session.Query("DELETE FROM postkeyspace.likes WHERE postid = ? AND userid = ? IF EXISTS;",
		like.PostID, like.UserID).Exec(); err != nil {
		fmt.Println("Error while deleting like!")
		fmt.Println(err)
		return err
	}
	return nil
}

func (repo *PostRepository) DeleteDislike(dislike *Model.Dislike) error {
	if err := repo.Session.Query("DELETE FROM postkeyspace.dislikes WHERE postid = ? AND userid = ? IF EXISTS;",
		dislike.PostID, dislike.UserID).Exec(); err != nil {
		fmt.Println("Error while deleting like!")
		fmt.Println(err)
		return err
	}
	return nil
}


/*func (repo *PostRepository) GetAllLikesForPost(postid string) error{
	var likes []Model.Like
	m := map[string]interface{}{}

	iter := repo.Session.Query("SELECT * FROM postkeyspace.postlikes WHERE postid = ?", postid).Iter()
	for iter.MapScan(m) {
		likes = append(likes, Model.Like{
			PostID:        m["postid"].(int),
			UserID: m["userid"].(int),
		})
		m = map[string]interface{}{}
	}

	Conv, _ := json.MarshalIndent(likes, "", " ")
	fmt.Println(string(Conv))
	return nil
}*/

func (repo *PostRepository) FindPostById(postid gocql.UUID) ( *Model.Post, error){
	var posts []Model.Post
	m := map[string]interface{}{}
	m2 := map[string]interface{}{}

	iter := repo.Session.Query("SELECT * FROM postkeyspace.posts WHERE id=? ALLOW FILTERING", postid).Iter()
	iter2 := repo.Session.Query("SELECT * FROM postkeyspace.postcounters WHERE postid=?", postid).Iter()
	fmt.Println(iter.NumRows())
	fmt.Println(iter2.NumRows())
	for i:=0; i<iter.NumRows(); i++{
		iter.MapScan(m)
		iter2.MapScan(m2)
		if iter2.NumRows() == 1{
			var a int64 = m2["likes"].(int64)
			var b int64 = m2["likes"].(int64)
			var c int64 = m2["likes"].(int64)
			var post = Model.Post{
				ID:        m["id"].(gocql.UUID),
				CreatedAt: m["createdat"].(time.Time),
				Description:  m["description"].(string),
				UserID:       m["userid"].(string),
				Image: m["image"].(string),
				LikesCount: a,
				DislikesCount: b,
				CommentsCount: c,
			}
			posts = append(posts, post)
			m = map[string]interface{}{}
			m2 = map[string]interface{}{}
		}else {
			var post = Model.Post{
				ID:          m["id"].(gocql.UUID),
				CreatedAt:   m["createdat"].(time.Time),
				Description: m["description"].(string),
				UserID:      m["userid"].(string),
				Image:       m["image"].(string),
			}

			posts = append(posts, post)
			m = map[string]interface{}{}
			m2 = map[string]interface{}{}
		}
	}
	return &posts[0],nil

}

func (repo *PostRepository) FindPostsByUserId(userid string) ( *[]Model.Post, error){
	var posts []Model.Post
	m := map[string]interface{}{}
	m2 := map[string]interface{}{}

	iter := repo.Session.Query("SELECT * FROM postkeyspace.posts WHERE userid=? ALLOW FILTERING", userid).Iter()
	for iter.MapScan(m) {
		iter2 := repo.Session.Query("SELECT * FROM postkeyspace.postcounters WHERE postid=?", m["id"].(gocql.UUID)).Iter()
		iter2.MapScan(m2)
		if iter2.NumRows() == 1{
			var a int64 = m2["likes"].(int64)
			var b int64 = m2["likes"].(int64)
			var c int64 = m2["likes"].(int64)
			var post = Model.Post{
				ID:        m["id"].(gocql.UUID),
				CreatedAt: m["createdat"].(time.Time),
				Description:  m["description"].(string),
				UserID:       m["userid"].(string),
				Image: m["image"].(string),
				LikesCount: a,
				DislikesCount: b,
				CommentsCount: c,
			}
			posts = append(posts, post)
			m = map[string]interface{}{}
			m2 = map[string]interface{}{}
		}else {
			var post = Model.Post{
				ID:          m["id"].(gocql.UUID),
				CreatedAt:   m["createdat"].(time.Time),
				Description: m["description"].(string),
				UserID:      m["userid"].(string),
				Image:       m["image"].(string),
			}

			posts = append(posts, post)
			m = map[string]interface{}{}
			m2 = map[string]interface{}{}
		}
	}
	return &posts,nil
}

func (repo *PostRepository) GetCommentsForPost(postid gocql.UUID) ( *[]Model.Comment, error) {
	var comments []Model.Comment
	m := map[string]interface{}{}

	iter := repo.Session.Query("SELECT * FROM postkeyspace.comments WHERE postid=?", postid).Iter()
	for iter.MapScan(m) {
		comments = append(comments, Model.Comment{
			ID:        m["id"].(gocql.UUID),
			PostID: m["postid"].(gocql.UUID),
			CreatedAt: m["createdat"].(time.Time),
			UserID:       m["userid"].(string),
			Content: m["content"].(string),
		})
		m = map[string]interface{}{}
	}
	return &comments,nil
}

func (repo *PostRepository) GetUsersWhoLikedPost(postid gocql.UUID) ( *[]string, error) {
	var userids []string
	var userid string
	iter := repo.Session.Query("SELECT userid FROM postkeyspace.likes WHERE postid=? ALLOW FILTERING", postid).Iter()
	for iter.Scan(&userid){
		fmt.Println(userid)
		userids = append(userids, userid)
	}
	return &userids, nil
}

func (repo *PostRepository) GetUsersWhoDislikedPost(postid gocql.UUID) ( *[]string, error) {
	var userids []string
	var userid string
	iter := repo.Session.Query("SELECT userid FROM postkeyspace.dislikes WHERE postid=? ALLOW FILTERING", postid).Iter()
	for iter.Scan(&userid){
		userids = append(userids, userid)
	}
	return &userids, nil
}

func (repo *PostRepository) GetImage(imagepath string) (image.Image, error){
	var directory string = "Images/"
	var imgpath string = directory + imagepath
	img, err := LoadImage(imgpath)
	if err != nil{
		return nil, err
	}
	fmt.Println("Image successfuly loaded!")
	return img, nil
}

func LoadImage(imagepath string) (image.Image, error){
	f, err := os.Open(imagepath)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	img, _, err := image.Decode(f)
	if err != nil {
		return nil, err
	}
	return img, nil
}