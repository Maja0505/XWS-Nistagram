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

func (service *PostService) DeleteComment(comment *Model.Comment) error {
	err := service.Repo.DeleteComment(comment)
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

func (service *PostService) FindPostsByUserId(userid string) ( *[]Model.Post, error) {
	posts,err := service.Repo.FindPostsByUserId(userid)
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
func (service *PostService) GetUsersWhoLikedPost(postid gocql.UUID) ( *[]DTO.UserByUsernameDTO, error) {
	userids, err := service.Repo.GetUsersWhoLikedPost(postid)

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

/*func (service *PostService) GetAllLikesForPost(postid string) error {
	err := service.Repo.GetAllLikesForPost(postid)
	if err != nil{
		fmt.Println(err)
		return  err
	}
	return nil
}*/
