package Service

import (
	"XWS-Nistagram/PostService/DTO"
	"XWS-Nistagram/PostService/Mapper"
	"XWS-Nistagram/PostService/Model"
	"XWS-Nistagram/PostService/Repository"
	"fmt"
	"github.com/gocql/gocql"
	"image"
)

type PostService struct {
	Repo Repository.PostRepository
}

func (service *PostService) Create(postDTO *DTO.PostDTO) error {
	post := Mapper.ConvertPostDTOToPost(postDTO)
	err := service.Repo.Create(post)
	if err != nil{
		fmt.Println(err)
		return  err
	}
	return nil
}

func (service *PostService) AddComment(commentDTO *DTO.CommentDTO) error {
	comment := Mapper.ConvertCommentDTOToComment(commentDTO)
	err := service.Repo.AddComment(comment)
	if err != nil{
		fmt.Println(err)
		return  err
	}
	return nil
}

func (service *PostService) AddTag(tag *Model.Tag) error {
	err := service.Repo.AddTag(tag)
	if err != nil{
		fmt.Println(err)
		return  err
	}
	return nil
}
func (service *PostService) AddPostToFavourites(favouriteDTO *DTO.FavouriteDTO) error {
	err := service.Repo.AddPostToFavourites(favouriteDTO.PostID, favouriteDTO.UserID)
	if err != nil{
		fmt.Println(err)
		return  err
	}
	return nil
}

func (service *PostService) AddPostToCollection(favouriteDTO *DTO.FavouriteDTO) error {
	err := service.Repo.AddPostToCollection(favouriteDTO.PostID, favouriteDTO.UserID, favouriteDTO.Collection)
	if err != nil{
		fmt.Println(err)
		return  err
	}
	return nil
}

func (service *PostService) DeleteComment(commentDTO *DTO.CommentDTO) error {
	comment := Mapper.ConvertCommentDTOToComment(commentDTO)
	err := service.Repo.DeleteComment(comment)
	if err != nil{
		fmt.Println(err)
		return  err
	}
	return nil
}

func (service *PostService) RemovePostFromFavourites(favourite *DTO.FavouriteDTO) error {
	err := service.Repo.RemovePostFromFavourites(favourite)
	if err != nil{
		fmt.Println(err)
		return  err
	}
	return nil
}

func (service *PostService) RemovePostFromCollection(favourite *DTO.FavouriteDTO) error {
	err := service.Repo.RemovePostFromFavourites(favourite)
	if err != nil{
		fmt.Println(err)
		return  err
	}
	return nil
}

func (service *PostService) LikePost(like *Model.Like) error {
	err := service.Repo.LikePost(like)
	if err != nil{
		fmt.Println(err)
		return  err
	}
	return nil
}

func (service *PostService) DislikePost(dislike *Model.Dislike) error {
	err := service.Repo.DislikePost(dislike)
	if err != nil{
		fmt.Println(err)
		return  err
	}
	return nil
}

func (service *PostService) CheckIfLikeExists(like *Model.Like) error {
	err := service.Repo.CheckIfLikeExists(like)
	if err == true{
	}else{
	}
	return nil
}

func (service *PostService) FindPostById(postid gocql.UUID) ( *Model.Post, error) {
	post,err := service.Repo.FindPostById(postid)
	if err != nil{
		fmt.Println(err)
		return  nil, err
	}
	return post, err
}

func (service *PostService) GetTagsForPost(postid gocql.UUID) ( *[]Model.Tag, error) {
	tags,err := service.Repo.GetTagsForPost(postid)
	if err != nil{
		fmt.Println(err)
		return  nil, err
	}
	return tags, err
}

func (service *PostService) GetFavouritePosts(userid string) ( *[]Model.Post, error) {
	posts,err := service.Repo.GetFavouritePosts(userid)
	if err != nil{
		fmt.Println(err)
		return  nil, err
	}
	return posts, err
}
func (service *PostService) GetPostsFromCollection(userid string, collection string) ( *[]Model.Post, error) {
	posts,err := service.Repo.GetPostsFromCollection(userid, collection)
	if err != nil{
		fmt.Println(err)
		return  nil, err
	}
	return posts, err
}

func (service *PostService) FindPostsByUserId(userid string) ( *[]Model.Post, error) {
	posts,err := service.Repo.FindPostsByUserId(userid)
	if err != nil{
		fmt.Println(err)
		return  nil, err
	}
	return posts, err
}

func (service *PostService) FindPostsByTag(tag string) ( *[]Model.Post, error) {
	posts,err := service.Repo.FindPostsByTag(tag)
	if err != nil{
		fmt.Println(err)
		return  nil, err
	}
	return posts, err
}

func (service *PostService) GetCommentsForPost(postid gocql.UUID) ( *[]Model.Comment, error) {
	comments, err := service.Repo.GetCommentsForPost(postid)
	if err != nil{
		fmt.Println(err)
		return  nil, err
	}
	return comments, err
}
func (service *PostService) GetUsersWhoLikedPost(postid gocql.UUID) ( *[]gocql.UUID, error) {
	userids, err := service.Repo.GetUsersWhoLikedPost(postid)
	if err != nil{
		fmt.Println(err)
		return  nil, err
	}
	return userids, err
}
func (service *PostService) GetUsersWhoDislikedPost(postid gocql.UUID) ( *[]gocql.UUID, error) {
	userids, err := service.Repo.GetUsersWhoDislikedPost(postid)
	if err != nil{
		fmt.Println(err)
		return  nil, err
	}
	return userids, err
}

func (service *PostService) GetImage(imagepath string) ( image.Image, error) {
	img, err := service.Repo.GetImage(imagepath)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	return img, err
}

/*func (service *PostService) GetAllLikesForPost(postid string) error {
	err := service.Repo.GetAllLikesForPost(postid)
	if err != nil{
		fmt.Println(err)
		return  err
	}
	return nil
}*/
