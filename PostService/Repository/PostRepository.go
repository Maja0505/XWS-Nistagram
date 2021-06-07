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

func (repo *PostRepository) Create(post *Model.Post) error {
	ID, _ := gocql.RandomUUID()
	if err := repo.Session.Query("INSERT INTO postkeyspace.postsv8(id, createdat, description, image, userid) VALUES(?, ?, ?, ?, ?)",
		ID, post.CreatedAt, post.Description, post.Image, post.UserID).Exec(); err != nil {
		fmt.Println("Error while creating post!")
		fmt.Println(err)
		return err
	}
	fmt.Println("Successfully created post!!")
	return nil
}

func (repo *PostRepository) AddComment(comment *Model.Comment) error {
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
	disliked = repo.CheckIfDislikeExists(&dislike)
	fmt.Println("Disliked: ", disliked)

	if err := repo.Session.Query("INSERT INTO postkeyspace.postlikes(postid, userid) VALUES(?, ?)",
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
		fmt.Println("eto")
		if err := repo.IncrementLikes(like); err != nil{
			return err
		}
	}else if liked == true && disliked == false{
		fmt.Println("zasto")
	}else if liked == false && disliked == true{
		fmt.Println("zato")
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

	if err := repo.Session.Query("INSERT INTO postkeyspace.postdislikes(postid, userid) VALUES(?, ?)",
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
	if err := repo.Session.Query("UPDATE postkeyspace.postsv8counters SET likes = likes + 1 WHERE postid = ?",
		like.PostID).Exec(); err != nil {
		fmt.Println("Error updating like/dislike counter!")
		fmt.Println(err)
		return err
	}
	return nil
}

func (repo *PostRepository) IncrementDislikes(dislike *Model.Dislike) error{
	if err := repo.Session.Query("UPDATE postkeyspace.postsv8counters SET dislikes = dislikes + 1 WHERE postid = ?",
		dislike.PostID).Exec(); err != nil {
		fmt.Println("Error updating like/dislike counter!")
		fmt.Println(err)
		return err
	}
	return nil
}

func (repo *PostRepository) DecrementLikes(like *Model.Like) error{
	if err := repo.Session.Query("UPDATE postkeyspace.postsv8counters SET likes = likes - 1 WHERE postid = ?",
		like.PostID).Exec(); err != nil {
		fmt.Println("Error updating like/dislike counter!")
		fmt.Println(err)
		return err
	}
	return nil
}

func (repo *PostRepository) DecrementDislikes(dislike *Model.Dislike) error{
	if err := repo.Session.Query("UPDATE postkeyspace.postsv8counters SET dislikes = dislikes - 1 WHERE postid = ?",
		dislike.PostID).Exec(); err != nil {
		fmt.Println("Error updating like/dislike counter!")
		fmt.Println(err)
		return err
	}
	return nil
}
func (repo *PostRepository) IncrementComments(comment *Model.Comment) error {
	if err := repo.Session.Query("UPDATE postkeyspace.postsv8counters SET comments = comments + 1 WHERE postid = ?",
		comment.PostID).Exec(); err != nil {
		fmt.Println("Error updating comments counter!")
		fmt.Println(err)
		return err
	}
	return nil
}
func (repo *PostRepository) DecrementComments(comment *Model.Comment) error {
	if err := repo.Session.Query("UPDATE postkeyspace.postsv8counters SET comments = comments - 1 WHERE postid = ?",
		comment.PostID).Exec(); err != nil {
		fmt.Println("Error updating comments counter!")
		fmt.Println(err)
		return err
	}
	return nil
}
func (repo *PostRepository) CheckIfLikeExists(like *Model.Like) bool {
	var likes int
	repo.Session.Query("SELECT COUNT(*) FROM postkeyspace.postlikes WHERE postid = ? AND userid = ? LIMIT 1",
		like.PostID,like.UserID).Iter().Scan(&likes)
	if likes == 1 {
		return true
	}else{
		return false
	}

}
func (repo *PostRepository) CheckIfDislikeExists(dislike *Model.Dislike) bool {
	var dislikes int
	repo.Session.Query("SELECT COUNT(*) FROM postkeyspace.postdislikes WHERE postid = ? AND userid = ? LIMIT 1",
		dislike.PostID, dislike.UserID).Iter().Scan(&dislikes)
	if dislikes == 1 {
		return true
	}else{
		return false
	}

}

func (repo *PostRepository) DeleteLike(like *Model.Like) error {
	if err := repo.Session.Query("DELETE FROM postkeyspace.postlikes WHERE postid = ? AND userid = ? IF EXISTS;",
		like.PostID, like.UserID).Exec(); err != nil {
		fmt.Println("Error while deleting like!")
		fmt.Println(err)
		return err
	}
	return nil
}

func (repo *PostRepository) DeleteDislike(dislike *Model.Dislike) error {
	if err := repo.Session.Query("DELETE FROM postkeyspace.postdislikes WHERE postid = ? AND userid = ? IF EXISTS;",
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

	iter := repo.Session.Query("SELECT * FROM postkeyspace.postsv8 WHERE id=? allow filtering", postid).Iter()
	for iter.MapScan(m) {
		posts = append(posts, Model.Post{
			ID:        m["id"].(gocql.UUID),
			CreatedAt: m["createdat"].(time.Time),
			Description:  m["description"].(string),
			UserID:       m["userid"].(gocql.UUID),
			Image: m["image"].(string),
		})
		m = map[string]interface{}{}
	}
	return &posts[0],nil
}

func (repo *PostRepository) FindPostsByUserId(userid gocql.UUID) ( *[]Model.Post, error){

	var posts []Model.Post
	m := map[string]interface{}{}

	iter := repo.Session.Query("SELECT * FROM postkeyspace.postsv8 WHERE userid=?", userid).Iter()
	for iter.MapScan(m) {
		posts = append(posts, Model.Post{
			ID:        m["id"].(gocql.UUID),
			CreatedAt: m["createdat"].(time.Time),
			Description:  m["description"].(string),
			UserID:       m["userid"].(gocql.UUID),
			Image: m["image"].(string),
		})
		m = map[string]interface{}{}
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
			UserID:       m["userid"].(gocql.UUID),
			Content: m["content"].(string),
		})
		m = map[string]interface{}{}
	}
	return &comments,nil
}

func (repo *PostRepository) GetUsersWhoLikedPost(postid gocql.UUID) ( *[]gocql.UUID, error) {
	var userids []gocql.UUID
	var useruuid gocql.UUID
	iter := repo.Session.Query("SELECT userid FROM postkeyspace.postlikes WHERE postid=?", postid).Iter()
	for iter.Scan(&useruuid){
		userids = append(userids, useruuid)
	}
	return &userids, nil
}

func (repo *PostRepository) GetUsersWhoDislikedPost(postid gocql.UUID) ( *[]gocql.UUID, error) {
	var userids []gocql.UUID
	var useruuid gocql.UUID
	iter := repo.Session.Query("SELECT userid FROM postkeyspace.postdislikes WHERE postid=?", postid).Iter()
	for iter.Scan(&useruuid){
		userids = append(userids, useruuid)
	}
	return &userids, nil
}

func (repo *PostRepository) GetImage(imagepath string) (image.Image, error){
	img, err := LoadImage(imagepath)
	if err != nil{
		return nil, err
	}
	fmt.Println("Image successfuly loaded")
	return img, nil
}

func LoadImage(imagepath string) (image.Image, error){
	f, err := os.Open(imagepath)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	img, fmtName, err := image.Decode(f)
	if err != nil {
		return nil, err
	}
	fmt.Println(fmtName)
	return img, nil
}