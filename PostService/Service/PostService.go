package Service

import (
	"XWS-Nistagram/PostService/DTO"
	"XWS-Nistagram/PostService/Mapper"
	"XWS-Nistagram/PostService/Model"
	"XWS-Nistagram/PostService/Repository"
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/gocql/gocql"
	"image"
	"io/ioutil"
	"net/http"
	"os"
	"sort"
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
	err := service.Repo.AddPostToFavourites(favouriteDTO)
	if err != nil{
		fmt.Println(err)
		return  err
	}
	return nil
}

func (service *PostService) AddPostToCollection(favouriteDTO *DTO.FavouriteDTO) error {
	err := service.Repo.AddPostToCollection(favouriteDTO)
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
	err := service.Repo.RemovePostFromCollection(favourite)
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

func (service *PostService) CheckIfLikeExists(like *Model.Like) bool {
	exist := service.Repo.CheckIfLikeExists(like)
	return exist
}

func (service *PostService) CheckIfDislikeExists(dislike *Model.Dislike) bool {
	exist := service.Repo.CheckIfDislikeExists(dislike)
	return exist
}

func (service *PostService) FindPostById(postid gocql.UUID) ( *Model.Post, error) {
	post,err := service.Repo.FindPostById(postid)
	if err != nil{
		fmt.Println(err)
		return  nil, err
	}
	return post, err
}

func (service *PostService) GetUserWhoPostedComment(commentid gocql.UUID) ( *[]DTO.UserByUsernameDTO, error) {
	username,err := service.Repo.GetUserWhoPostedComment(commentid)
	if err != nil{
		fmt.Println(err)
		return  nil, err
	}
	reqUrl := fmt.Sprintf("http://" + os.Getenv("USER_SERVICE_DOMAIN") + ":" + os.Getenv("USER_SERVICE_PORT") + "/convert-usernames")

	type UsernamesDTO struct {
		Usernames []string
	}

	usernamesDTO := UsernamesDTO{}
	usernamesDTO.Usernames = append(usernamesDTO.Usernames, *username)
	fmt.Println(usernamesDTO.Usernames)
	jsonUserids,_ := json.Marshal(usernamesDTO)

	resp, err := http.Post(reqUrl,"appliation/json",bytes.NewBuffer(jsonUserids))
	if err != nil || resp.StatusCode == 404 {
		return nil,err
	}
	body, err := ioutil.ReadAll(resp.Body)
	fmt.Println(body)
	var data []DTO.UserByUsernameDTO
	err = json.Unmarshal(body, &data)
	fmt.Println(data)
	if err != nil{
		return nil, err
	}
	return &data, nil
}

func (service *PostService) GetTagsForPost(postid gocql.UUID) ( *[]Model.Tag, error) {
	tags,err := service.Repo.GetTagsForPost(postid)
	if err != nil{
		fmt.Println(err)
		return  nil, err
	}
	return tags, err
}

func (service *PostService) GetPureTagsForPost(postid gocql.UUID) ( *[]Model.Tag, error) {
	tags,err := service.Repo.GetPureTagsForPost(postid)
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

func (service *PostService) GetUsersTaggedOnPost(postid gocql.UUID) ( *[]DTO.UserByUsernameDTO, error) {
	usernames, err := service.Repo.GetUsersTaggedOnPost(postid)
	if err != nil{
		fmt.Println(err)
		return  nil, err
	}
	fmt.Println(usernames)

	reqUrl := fmt.Sprintf("http://" + os.Getenv("USER_SERVICE_DOMAIN") + ":" + os.Getenv("USER_SERVICE_PORT") + "/convert-usernames")

	type UsernamesDTO struct {
		Usernames []string
	}

	usernamesDTO := UsernamesDTO{}
	usernamesDTO.Usernames = *usernames
	jsonUserids,_ := json.Marshal(usernamesDTO)

	resp, err := http.Post(reqUrl,"appliation/json",bytes.NewBuffer(jsonUserids))
	if err != nil || resp.StatusCode == 404 {
		return nil,err
	}
	body, err := ioutil.ReadAll(resp.Body)
	var data []DTO.UserByUsernameDTO
	err = json.Unmarshal(body, &data)
	if err != nil{
		return nil, err
	}
	return &data, nil
}

func (service *PostService) GetUsersWhoLikedPost(postid gocql.UUID) ( *[]DTO.UserByUsernameDTO, error) {
	userids, err := service.Repo.GetUsersWhoLikedPost(postid)

	if err != nil{
		fmt.Println(err)
		return  nil, err
	}
	reqUrl := fmt.Sprintf("http://" + os.Getenv("USER_SERVICE_DOMAIN") + ":" + os.Getenv("USER_SERVICE_PORT") + "/convert-user-ids") //namestiti da moze i lokalno

	type UserIdsDTO struct {
		UserIds []string
	}

	userIdsDto := UserIdsDTO{}
	userIdsDto.UserIds = *userids
	jsonUserids,_ := json.Marshal(userIdsDto)

	resp, err := http.Post(reqUrl,"appliation/json",bytes.NewBuffer(jsonUserids))
	if err != nil || resp.StatusCode == 404 {
		return nil,err
	}
	body, err := ioutil.ReadAll(resp.Body)
	var data []DTO.UserByUsernameDTO
	err = json.Unmarshal(body, &data)
	if err != nil{
		return nil, err
	}
	return &data, nil
}
func (service *PostService) GetUsersWhoDislikedPost(postid gocql.UUID) ( *[]DTO.UserByUsernameDTO, error) {
	userids, err := service.Repo.GetUsersWhoDislikedPost(postid)

	if err != nil{
		fmt.Println(err)
		return  nil, err
	}
	reqUrl := fmt.Sprintf("http://user-service:8080/convert-user-ids") //namestiti da moze i lokalno

	type UserIdsDTO struct {
		UserIds []string
	}

	userIdsDto := UserIdsDTO{}
	userIdsDto.UserIds = *userids
	jsonUserids,_ := json.Marshal(userIdsDto)

	resp, err := http.Post(reqUrl,"appliation/json",bytes.NewBuffer(jsonUserids))
	if err != nil || resp.StatusCode == 404 {
		return nil,err
	}
	body, err := ioutil.ReadAll(resp.Body)
	var data []DTO.UserByUsernameDTO
	err = json.Unmarshal(body, &data)
	if err != nil{
		return nil, err
	}
	return &data, nil
}

func (service *PostService) GetImage(imagepath string) ( image.Image, error) {
	img, err := service.Repo.GetImage(imagepath)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	return img, err
}

func (service *PostService) GetLikedPostsForUser(userid string) (*[]Model.Post, error) {
	likedPost, err := service.Repo.GetLikedPostsForUser(userid)

	if err != nil{
		fmt.Println(err)
		return  nil, err
	}

	return likedPost, err
}

func (service *PostService) GetDislikedPostsForUser(userid string) (*[]Model.Post, error) {
	dislikedPost, err := service.Repo.GetDislikedPostsForUser(userid)

	if err != nil{
		fmt.Println(err)
		return  nil, err
	}

	return dislikedPost, err
}

func (service *PostService) ReportContent(reportedContentDTO *DTO.ReportedContentDTO) error {
	reportedContent := Mapper.ConvertReportedContentDTOToReportedContent(reportedContentDTO)
	err := service.Repo.ReportContent(reportedContent)
	if err != nil{
		fmt.Println(err)
		return  err
	}
	return nil
}

func (service *PostService) GetCollectionsForUser(userid string) (*[]string,error) {
	collections, err := service.Repo.GetCollectionsForUser(userid)

	if err != nil{
		fmt.Println(err)
		return  nil, err
	}

	return collections, err
}

func (service *PostService) CheckIfPostExistsInFavourites(userid string, postid gocql.UUID) bool {
	exist := service.Repo.CheckIfPostIsInFavourites(userid,postid)
	return exist
}

func (service *PostService) GetAllCollectionsForPostByUser(userid string, postuuid gocql.UUID) (*[]string,error) {
	collections, err := service.Repo.GetAllCollectionsForPostByUser(userid,postuuid)

	if err != nil{
		fmt.Println(err)
		return  nil, err
	}

	return collections, err
}

func (service *PostService) GetAllPostFeedsForUser(userid string) ( *[]Model.Post, error){

	var postsByAllNotMutedFollowedUsers []Model.Post

	reqUrl := fmt.Sprintf("http://" + os.Getenv("USER_FOLLOWERS_SERVICE_DOMAIN") + ":" + os.Getenv("USER_FOLLOWERS_SERVICE_PORT") + "/allNotMutedFollows/" + userid)

	resp, err := http.Get(reqUrl)
	if err != nil || resp.StatusCode == 404 {
		return nil,err
	}
	body, err := ioutil.ReadAll(resp.Body)
	var notMutedFollowedUsers []string
	err = json.Unmarshal(body, &notMutedFollowedUsers)
	if err != nil{
		return nil, err
	}

	for _,userId := range notMutedFollowedUsers {

		postsByOneUser,err := service.Repo.FindPostsByUserId(userId)

		if err != nil {
			return nil, err
		}
		if postsByOneUser != nil {
			postsByAllNotMutedFollowedUsers = append(postsByAllNotMutedFollowedUsers, *postsByOneUser...)
		}
	}

	var feedSlice FeedSlice
	feedSlice = postsByAllNotMutedFollowedUsers
	sort.Sort(feedSlice)
	postsByAllNotMutedFollowedUsers = feedSlice

	return &postsByAllNotMutedFollowedUsers,nil

}


//for sorting post feeds by created time
type FeedSlice []Model.Post

func (f FeedSlice) Len() int {
	return len(f)
}

func (f FeedSlice) Less(i, j int) bool {
	return f[i].CreatedAt.After(f[j].CreatedAt)
}

func (f FeedSlice) Swap(i, j int) {
	f[i], f[j] = f[j], f[i]
}
