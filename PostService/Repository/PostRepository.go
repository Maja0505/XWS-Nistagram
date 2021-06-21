package Repository

import (
	"XWS-Nistagram/PostService/DTO"
	"XWS-Nistagram/PostService/Model"
	"fmt"
	"github.com/gocql/gocql"
	"image"
	"os"
	"strings"
	"time"
)

type PostRepository struct {
	Session gocql.Session
}

func (repo *PostRepository) CreateTables() error{
	/*if err := repo.Session.Query("DROP TABLE IF EXISTS postkeyspace.posts;").Exec(); err != nil {
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
	if err := repo.Session.Query("DROP TABLE IF EXISTS postkeyspace.tagsDK;").Exec(); err != nil {
		fmt.Println("Error while dropping tables!")
		fmt.Println(err)
		return err
	}
	if err := repo.Session.Query("DROP TABLE IF EXISTS postkeyspace.favourites;").Exec(); err != nil {
		fmt.Println("Error while dropping tables!")
		fmt.Println(err)
		return err
	}
	if err := repo.Session.Query("DROP TABLE IF EXISTS postkeyspace.collections;").Exec(); err != nil {
		fmt.Println("Error while dropping tables!")
		fmt.Println(err)
		return err
	}

	*/if err := repo.Session.Query("CREATE TABLE if not exists postkeyspace.posts(id uuid, userid text, createdat timestamp, description text, image text, location text, PRIMARY KEY((userid, id)));").Exec(); err != nil {
		fmt.Println("Error while creating tables!")
		fmt.Println(err)
		return err
	}
	if err := repo.Session.Query("CREATE TABLE if not exists postkeyspace.postcounters(postid uuid, likes counter, dislikes counter, comments counter, PRIMARY KEY(postid));").Exec(); err != nil {
		fmt.Println("Error while creating tables!")
		fmt.Println(err)
		return err
	}
	if err := repo.Session.Query("CREATE TABLE if not exists postkeyspace.comments(id uuid, postid uuid, userid text, createdat timestamp, content text, PRIMARY KEY((postid), userid, id));").Exec(); err != nil {
		fmt.Println("Error while creating tables!")
		fmt.Println(err)
		return err
	}
	if err := repo.Session.Query("CREATE TABLE if not exists postkeyspace.likes(postid uuid, userid text, PRIMARY KEY((postid, userid)));").Exec(); err != nil {
		fmt.Println("Error while creating tables!")
		fmt.Println(err)
		return err
	}
	if err := repo.Session.Query("CREATE TABLE if not exists postkeyspace.dislikes(postid uuid, userid text, PRIMARY KEY((postid, userid)));").Exec(); err != nil {
		fmt.Println("Error while creating tables!")
		fmt.Println(err)
		return err
	}
	if err := repo.Session.Query("CREATE TABLE if not exists postkeyspace.tags(postid uuid, tag text, PRIMARY KEY((postid), tag));").Exec(); err != nil {
		fmt.Println("Error while creating tables!")
		fmt.Println(err)
		return err
	}
	if err := repo.Session.Query("CREATE TABLE if not exists postkeyspace.tagsDK(postid uuid, tag text, PRIMARY KEY((tag), postid));").Exec(); err != nil {
		fmt.Println("Error while creating tables!")
		fmt.Println(err)
		return err
	}
	if err := repo.Session.Query("CREATE TABLE if not exists postkeyspace.favourites(userid text, postid uuid, PRIMARY KEY((userid), postid));").Exec(); err != nil {
		fmt.Println("Error while creating tables!")
		fmt.Println(err)
		return err
	}
	if err := repo.Session.Query("CREATE TABLE if not exists postkeyspace.collections(userid text, postid uuid, collection text, PRIMARY KEY((userid), collection, postid));").Exec(); err != nil {
		fmt.Println("Error while creating tables!")
		fmt.Println(err)
		return err
	}

	if err := repo.Session.Query("CREATE TABLE if not exists postkeyspace.reported_contents(id uuid, description text, contentid text, userid text, adminid text, PRIMARY KEY((userid, id)));").Exec(); err != nil {
		fmt.Println("Error while creating tables!")
		fmt.Println(err)
		return err
	}
	fmt.Println("Successfully dropped and created tables!!")

	return nil
}

func (repo *PostRepository) Create(post *Model.Post) error {
	//ID, _ := gocql.RandomUUID()
	if err := repo.Session.Query("INSERT INTO postkeyspace.posts(id, createdat, description, image, userid) VALUES(?, ?, ?, ?, ?)",
		post.ID, post.CreatedAt, post.Description, post.Image, post.UserID).Exec(); err != nil {
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

func (repo *PostRepository) AddPostToFavourites(favourite *DTO.FavouriteDTO) error {
	if err := repo.Session.Query("INSERT INTO postkeyspace.favourites(userid, postid) VALUES(?, ?) IF NOT EXISTS",
		favourite.UserID, favourite.PostID).Exec(); err != nil {
		fmt.Println("Error while adding post to favourites!")
		fmt.Println(err)
		return err
	}
	fmt.Println("Successfully added post to favourites!!")
	return nil
}

func (repo *PostRepository) AddPostToCollection(favourite *DTO.FavouriteDTO) error {
	if repo.CheckIfPostIsInFavourites(favourite.UserID, favourite.PostID) == false{
		return gocql.Error{Message: "Post is not in favourites!!"}
	}
	if err := repo.Session.Query("INSERT INTO postkeyspace.collections(userid, postid, collection) VALUES(?, ?, ?) IF NOT EXISTS",
		favourite.UserID, favourite.PostID, favourite.Collection).Exec(); err != nil {
		fmt.Println("Error while adding post to collection: ", favourite.Collection)
		fmt.Println(err)
		return err
	}
	fmt.Println("Successfully added post to collection: ", favourite.Collection)
	return nil
}

func (repo *PostRepository) AddTag(tag *Model.Tag) error {
	if err := repo.Session.Query("INSERT INTO postkeyspace.tags(tag, postid) VALUES(?, ?)",
		tag.Tag, tag.PostID).Exec(); err != nil {
		fmt.Println("Error while creating tag!")
		fmt.Println(err)
		return err
	}
	if err := repo.Session.Query("INSERT INTO postkeyspace.tagsDK(tag, postid) VALUES(?, ?)",
		tag.Tag, tag.PostID).Exec(); err != nil {
		fmt.Println("Error while creating tag!")
		fmt.Println(err)
		return err
	}
	fmt.Println("Successfully created tag!!")
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

func (repo *PostRepository) RemovePostFromFavourites(favourite *DTO.FavouriteDTO) error {
	if err := repo.Session.Query("DELETE FROM postkeyspace.favourites where userid = ? AND postid = ? IF EXISTS;",
		favourite.UserID, favourite.PostID).Exec(); err != nil {
		fmt.Println("Error while deleting post from favourites!")
		fmt.Println(err)
		return err
	}
	collections,err := repo.GetAllCollectionsForPostByUser(favourite.UserID,favourite.PostID)
	if err != nil{
		fmt.Println(err)
	}
	for _, collection := range *collections {
		favourite.Collection = collection
		err = repo.RemovePostFromCollection(favourite)
		if err != nil{
			fmt.Println(err)
		}
	}

	fmt.Println("Successfully deleted post from favourites!!")
	return nil
}

func (repo *PostRepository) RemovePostFromCollection(favourite *DTO.FavouriteDTO) error {
	if err := repo.Session.Query("DELETE FROM postkeyspace.collections where userid = ? AND postid = ? AND collection = ? IF EXISTS;",
		favourite.UserID, favourite.PostID, favourite.Collection).Exec(); err != nil {
		fmt.Println("Error while deleting post from collection: ", favourite.Collection)
		fmt.Println(err)
		return err
	}
	fmt.Println("Successfully deleted post from collection: ", favourite.Collection)
	return nil
}

func (repo *PostRepository) LikePost(like *Model.Like) error {

	var dislike Model.Dislike
	dislike.PostID = like.PostID
	dislike.UserID = like.UserID
	var liked bool
	var disliked bool

	liked = repo.CheckIfLikeExists(like)
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
	if err := repo.Session.Query("INSERT INTO postkeyspace.likes(postid, userid) VALUES(?, ?)",
		like.PostID, like.UserID).Exec(); err != nil {
		fmt.Println("Error while creating like!")
		fmt.Println(err)
		return err
	}
	if err := repo.DeleteDislike(&dislike); err != nil {
		return err
	}
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
	disliked = repo.CheckIfDislikeExists(dislike)
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

func (repo *PostRepository) CheckIfPostIsInFavourites(userid string, postid gocql.UUID) bool {
	var exists int
	repo.Session.Query("SELECT COUNT(*) FROM postkeyspace.favourites WHERE userid = ? AND postid = ? LIMIT 1",
		userid, postid).Iter().Scan(&exists)
	if exists == 1 {
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

func (repo *PostRepository) FindPostById(postid gocql.UUID) ( *Model.Post, error){
	var posts []Model.Post
	m := map[string]interface{}{}
	m2 := map[string]interface{}{}

	iter := repo.Session.Query("SELECT * FROM postkeyspace.posts WHERE id=? ALLOW FILTERING", postid).Iter()
	if iter.NumRows() == 0{
		return nil, gocql.Error{Message: "No post found"}
	}
	iter2 := repo.Session.Query("SELECT * FROM postkeyspace.postcounters WHERE postid=?", postid).Iter()
	for i:=0; i<iter.NumRows(); i++{
		iter.MapScan(m)
		iter2.MapScan(m2)
		if iter2.NumRows() == 1{
			var a int64 = m2["likes"].(int64)
			var b int64 = m2["dislikes"].(int64)
			var c int64 = m2["comments"].(int64)
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
	if iter.NumRows() == 0{
		return nil, gocql.Error{Message: "No post found"}
	}
	for iter.MapScan(m) {
		iter2 := repo.Session.Query("SELECT * FROM postkeyspace.postcounters WHERE postid=?", m["id"].(gocql.UUID)).Iter()
		/*if iter2.NumRows() == 0{
			//continue
		}*/
		iter2.MapScan(m2)
		if iter2.NumRows() == 1{
			var a int64 = m2["likes"].(int64)
			var b int64 = m2["dislikes"].(int64)
			var c int64 = m2["comments"].(int64)
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

func (repo *PostRepository) FindPostsByTag(tag string) ( *[]Model.Post, error){
	var tags []Model.Tag
	var posts []Model.Post
	m := map[string]interface{}{}

	iter := repo.Session.Query("SELECT * FROM postkeyspace.tagsDK WHERE tag=?", tag).Iter()
	for iter.MapScan(m) {
		var tag = Model.Tag{
			PostID:      m["postid"].(gocql.UUID),
			Tag:   		 m["tag"].(string),
		}
		tags = append(tags, tag)
		m = map[string]interface{}{}
	}
	for i:=0; i< len(tags); i++{
		var post,err = repo.FindPostById(tags[i].PostID)
		if err != nil{
			continue
		}
		if post != nil {
			posts = append(posts, *post)
		}
	}
	return &posts, nil
}

func (repo *PostRepository) GetFavouritePosts(userid string) ( *[]Model.Post, error){
	var postids []gocql.UUID
	var posts []Model.Post
	m := map[string]interface{}{}

	iter := repo.Session.Query("SELECT * FROM postkeyspace.favourites WHERE userid=?", userid).Iter()
	for iter.MapScan(m) {
		var postid = m["postid"].(gocql.UUID)
		postids = append(postids, postid)
		m = map[string]interface{}{}
	}
	for i:=0; i< len(postids); i++{
		var post,err = repo.FindPostById(postids[i])
		if err != nil{
			continue
		}
		if post != nil {
			posts = append(posts, *post)
		}
	}
	return &posts, nil
}

func (repo *PostRepository) GetPostsFromCollection(userid string, collection string) ( *[]Model.Post, error){
	var postids []gocql.UUID
	var posts []Model.Post
	m := map[string]interface{}{}

	iter := repo.Session.Query("SELECT * FROM postkeyspace.collections WHERE userid=? AND collection=?", userid, collection).Iter()
	for iter.MapScan(m) {
		var postid = m["postid"].(gocql.UUID)
		postids = append(postids, postid)
		m = map[string]interface{}{}
	}
	for i:=0; i< len(postids); i++{
		var post,err = repo.FindPostById(postids[i])
		if err != nil{
			continue
		}
		if post != nil {
			posts = append(posts, *post)
		}
	}
	return &posts, nil
}

func (repo *PostRepository) GetTagsForPost(postid gocql.UUID) ( *[]Model.Tag, error) {
	var tags []Model.Tag
	m := map[string]interface{}{}

	iter := repo.Session.Query("SELECT * FROM postkeyspace.tags WHERE postid=?", postid).Iter()
	for iter.MapScan(m) {
			var tag = Model.Tag{
				PostID:      m["postid"].(gocql.UUID),
				Tag:   		 m["tag"].(string),
			}
			tags = append(tags, tag)
			m = map[string]interface{}{}
	}
	return &tags,nil
}

func (repo *PostRepository) GetPureTagsForPost(postid gocql.UUID) ( *[]Model.Tag, error) {
	var tags []Model.Tag
	m := map[string]interface{}{}

	iter := repo.Session.Query("SELECT * FROM postkeyspace.tags WHERE postid=?", postid).Iter()
	for iter.MapScan(m) {
		var tag = Model.Tag{
			PostID:      m["postid"].(gocql.UUID),
			Tag:   		 m["tag"].(string),
		}
		if strings.HasPrefix(tag.Tag, "@") == false{
			tags = append(tags, tag)
		}
		m = map[string]interface{}{}
	}
	return &tags,nil
}

func (repo *PostRepository) GetUsersTaggedOnPost(postid gocql.UUID) (*[]string, error){
	var tags, err = repo.GetTagsForPost(postid)
	if err != nil{
		return nil, err
	}
	var usernames []string
	for _, tag := range *tags{
		if strings.HasPrefix(tag.Tag,"@"){
			var username = strings.SplitAfter(tag.Tag,"@")
			usernames = append(usernames, username[1])
		}
	}
	return &usernames, nil
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

func (repo *PostRepository) GetUserWhoPostedComment(commentid gocql.UUID) ( *string, error){
	iter:= repo.Session.Query("SELECT * FROM postkeyspace.comments WHERE id=? ALLOW FILTERING", commentid).Iter()
	if iter.NumRows() == 0 {
		return nil, gocql.Error{Message: "Comment does not exist!"}
	}
	var username string
	m := map[string]interface{}{}
	for iter.MapScan(m) {
		username = m["userid"].(string)
		m = map[string]interface{}{}
	}
	if username == ""{
		return nil, gocql.Error{Message: "Username empty!"}
	}
	return &username, nil
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
	var directory string = "post-documents/"
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

func (repo *PostRepository) GetLikedPostsForUser(userid string) ( *[]Model.Post, error) {

	var postids []gocql.UUID
	var posts []Model.Post
	m := map[string]interface{}{}
	iter := repo.Session.Query("SELECT * FROM postkeyspace.likes WHERE userid=? ALLOW FILTERING", userid).Iter()

	for iter.MapScan(m) {
		var postid = m["postid"].(gocql.UUID)
		postids = append(postids, postid)
		m = map[string]interface{}{}
	}
	for i:=0; i< len(postids); i++{
		var post,err = repo.FindPostById(postids[i])
		if err != nil{
			continue
		}
		if post != nil {
			posts = append(posts, *post)
		}
	}
	return &posts, nil
}

func (repo *PostRepository) GetDislikedPostsForUser(userid string) (*[]Model.Post, error) {

	var postids []gocql.UUID
	var posts []Model.Post
	m := map[string]interface{}{}
	iter := repo.Session.Query("SELECT * FROM postkeyspace.dislikes WHERE userid=? ALLOW FILTERING", userid).Iter()

	for iter.MapScan(m) {
		var postid = m["postid"].(gocql.UUID)
		postids = append(postids, postid)
		m = map[string]interface{}{}
	}
	for i:=0; i< len(postids); i++{
		var post,err = repo.FindPostById(postids[i])
		if err != nil{
			continue
		}
		if post != nil {
			posts = append(posts, *post)
		}
	}
	return &posts, nil
}

func (repo *PostRepository) ReportContent(content *Model.ReportedContent) error {
	ID, _ := gocql.RandomUUID()
	if err := repo.Session.Query("INSERT INTO postkeyspace.reported_contents(id, description, contentid, userid, adminid) VALUES(?, ?, ?, ?, ?)",
		ID, content.Description, content.ContentID, content.UserID, content.AdminID).Exec(); err != nil {
		fmt.Println("Error while creating report!")
		fmt.Println(err)
		return err
	}
	fmt.Println("Successfully created report!!")
	return nil
}

func (repo *PostRepository) GetCollectionsForUser(userid string) (*[] string, error) {
	var collections []string
	var collection string
	iter := repo.Session.Query("SELECT collection FROM postkeyspace.collections WHERE userid=? ALLOW FILTERING", userid).Iter()
	for iter.Scan(&collection){
		collections = append(collections, collection)
	}

	return &collections, nil
}

func (repo *PostRepository) GetAllCollectionsForPostByUser(userid string, postuuid gocql.UUID) ( *[]string, error) {
	var collections []string
	var collection string
	iter := repo.Session.Query("SELECT collection FROM postkeyspace.collections WHERE userid=? and postid=? ALLOW FILTERING", userid,postuuid).Iter()
	for iter.Scan(&collection){
		collections = append(collections, collection)
	}

	return &collections, nil
}




