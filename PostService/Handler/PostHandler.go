package Handler

import (
	"XWS-Nistagram/PostService/DTO"
	"XWS-Nistagram/PostService/Model"
	"XWS-Nistagram/PostService/Service"
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/gocql/gocql"
	"github.com/gorilla/mux"
	"image/jpeg"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
)

type PostHandler struct {
	Service *Service.PostService
}

func (handler *PostHandler) Create(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var postDTO DTO.PostDTO
	err := json.NewDecoder(r.Body).Decode(&postDTO)
	if err != nil{
		fmt.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	id,err := handler.Service.Create(&postDTO)
	if err != nil{
		fmt.Println(err)
		w.WriteHeader(http.StatusExpectationFailed)
		return
	}

	w.WriteHeader(http.StatusCreated)
	_ = json.NewEncoder(w).Encode(id)
}

func (handler *PostHandler) AddPostToFavourites(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var favouriteDTO DTO.FavouriteDTO
	err := json.NewDecoder(r.Body).Decode(&favouriteDTO)
	if err != nil{
		fmt.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	err = handler.Service.AddPostToFavourites(&favouriteDTO)
	if err != nil{
		fmt.Println(err)
		w.WriteHeader(http.StatusExpectationFailed)
		return
	}

	w.WriteHeader(http.StatusCreated)

}

func (handler *PostHandler) AddPostToCollection(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var favouriteDTO DTO.FavouriteDTO
	err := json.NewDecoder(r.Body).Decode(&favouriteDTO)
	if err != nil{
		fmt.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	err = handler.Service.AddPostToCollection(&favouriteDTO)
	if err != nil{
		fmt.Println(err)
		w.WriteHeader(http.StatusExpectationFailed)
		return
	}

	w.WriteHeader(http.StatusCreated)

}

func (handler *PostHandler) AddComment(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var commentDTO DTO.CommentDTO
	err := json.NewDecoder(r.Body).Decode(&commentDTO)
	if err != nil{
		fmt.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	err = handler.Service.AddComment(&commentDTO)
	if err != nil{
		fmt.Println(err)
		w.WriteHeader(http.StatusExpectationFailed)
		return
	}

	w.WriteHeader(http.StatusCreated)

}

func (handler *PostHandler) AddTag(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var tag Model.Tag
	err := json.NewDecoder(r.Body).Decode(&tag)
	if err != nil{
		fmt.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	err = handler.Service.AddTag(&tag)
	if err != nil{
		fmt.Println(err)
		w.WriteHeader(http.StatusExpectationFailed)
		return
	}

	w.WriteHeader(http.StatusCreated)

}

func (handler *PostHandler) DeleteComment(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var comment DTO.CommentDTO
	err := json.NewDecoder(r.Body).Decode(&comment)
	if err != nil{
		fmt.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	err = handler.Service.DeleteComment(&comment)
	if err != nil{
		fmt.Println(err)
		w.WriteHeader(http.StatusExpectationFailed)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

func (handler *PostHandler) RemovePostFromFavourites(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var favouriteDTO DTO.FavouriteDTO
	err := json.NewDecoder(r.Body).Decode(&favouriteDTO)
	if err != nil{
		fmt.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	err = handler.Service.RemovePostFromFavourites(&favouriteDTO)
	if err != nil{
		fmt.Println(err)
		w.WriteHeader(http.StatusExpectationFailed)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

func (handler *PostHandler) RemovePostFromCollection(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var favouriteDTO DTO.FavouriteDTO
	err := json.NewDecoder(r.Body).Decode(&favouriteDTO)
	if err != nil{
		fmt.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	err = handler.Service.RemovePostFromCollection(&favouriteDTO)
	if err != nil{
		fmt.Println(err)
		w.WriteHeader(http.StatusExpectationFailed)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

func (handler *PostHandler) LikePost(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var like Model.Like
	err := json.NewDecoder(r.Body).Decode(&like)
	if err != nil{
		fmt.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	err = handler.Service.LikePost(&like)
	if err != nil{
		fmt.Println(err)
		w.WriteHeader(http.StatusExpectationFailed)
		return
	}

	w.WriteHeader(http.StatusCreated)

}

func (handler *PostHandler) DislikePost(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var dislike Model.Dislike
	err := json.NewDecoder(r.Body).Decode(&dislike)
	if err != nil{
		fmt.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	err = handler.Service.DislikePost(&dislike)
	if err != nil{
		fmt.Println(err)
		w.WriteHeader(http.StatusExpectationFailed)
		return
	}

	w.WriteHeader(http.StatusCreated)

}

func (handler *PostHandler) CheckIfLikeExists(w http.ResponseWriter, r *http.Request)  {
	w.Header().Set("Content-Type", "application/json")
	var like Model.Like
	err := json.NewDecoder(r.Body).Decode(&like)
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	exist := handler.Service.CheckIfLikeExists(&like)
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(exist)
}

func (handler *PostHandler) CheckIfDislikeExists(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")
	var dislike Model.Dislike
	err := json.NewDecoder(r.Body).Decode(&dislike)
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	exist := handler.Service.CheckIfDislikeExists(&dislike)
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(exist)
}

func (handler *PostHandler) FindPostById(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")
	vars := mux.Vars(r)
	postid := vars["id"]
	if postid == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	var postuuid, err = ParseUUID(postid)
	if err != nil{
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	post,_ := handler.Service.FindPostById(postuuid)

	if post == nil{
		w.WriteHeader(http.StatusNotFound)
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(post)

}

func (handler *PostHandler) GetFavouritePosts(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")
	vars := mux.Vars(r)
	userid := vars["id"]
	if userid == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	posts,_ := handler.Service.GetFavouritePosts(userid)

	if posts == nil{
		w.WriteHeader(http.StatusNotFound)
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(posts)

}

func (handler *PostHandler) GetPostsFromCollection(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")
	vars := mux.Vars(r)
	userid := vars["id"]
	collection := vars["collection"]
	if userid == ""  || collection == ""{
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	posts,_ := handler.Service.GetPostsFromCollection(userid, collection)

	if posts == nil{
		w.WriteHeader(http.StatusNotFound)
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(posts)

}

func (handler *PostHandler) FindPostsByUserId(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")
	vars := mux.Vars(r)
	userid := vars["id"]
	if userid == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	posts,_ := handler.Service.FindPostsByUserId(userid)

	if posts == nil{
		w.WriteHeader(http.StatusNotFound)
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(posts)

}

func (handler *PostHandler) FindPostsByTag(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")
	vars := mux.Vars(r)
	tag := vars["tag"]
	if tag == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}else{
		if tag[0:1] != "@" {
			tag = "#" + tag
		}

	}

	posts,_ := handler.Service.FindPostsByTag(tag)

	if posts == nil{
		w.WriteHeader(http.StatusNotFound)
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(posts)

}

func (handler *PostHandler) FindPostsByLocation(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")
	vars := mux.Vars(r)
	location := vars["location"]
	if location == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	posts,_ := handler.Service.FindPostsByLocation(location)

	if posts == nil{
		w.WriteHeader(http.StatusNotFound)
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(posts)

}

func (handler *PostHandler) GetTagsForPost(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")
	vars := mux.Vars(r)
	postid := vars["id"]
	postuuid, err := ParseUUID(postid)

	if err != nil{
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	tags,_ := handler.Service.GetTagsForPost(postuuid)

	if tags == nil{
		w.WriteHeader(http.StatusNotFound)
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(tags)

}

func (handler *PostHandler) GetPureTagsForPost(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")
	vars := mux.Vars(r)
	postid := vars["id"]
	postuuid, err := ParseUUID(postid)

	if err != nil{
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	tags,_ := handler.Service.GetPureTagsForPost(postuuid)

	if tags == nil{
		w.WriteHeader(http.StatusNotFound)
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(tags)

}

func (handler *PostHandler) GetCommentsForPost(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")
	vars := mux.Vars(r)
	postid := vars["id"]
	if postid == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	var postuuid, err = ParseUUID(postid)
	if err != nil{
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	comments,_ := handler.Service.GetCommentsForPost(postuuid)

	if comments == nil{
		w.WriteHeader(http.StatusNotFound)
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(comments)

}

func (handler *PostHandler) GetUserWhoPostedComment(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")
	vars := mux.Vars(r)
	commentid := vars["id"]
	if commentid == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	var commentuuid, err = ParseUUID(commentid)
	if err != nil{
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	username,_ := handler.Service.GetUserWhoPostedComment(commentuuid)

	if username == nil{
		w.WriteHeader(http.StatusNotFound)
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(username)

}

func (handler *PostHandler) GetUsersTaggedOnPost(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")
	vars := mux.Vars(r)
	postid := vars["id"]
	if postid == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	var postuuid, err = ParseUUID(postid)
	if err != nil{
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	userids,_ := handler.Service.GetUsersTaggedOnPost(postuuid)

	if userids == nil{
		w.WriteHeader(http.StatusNotFound)
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(userids)

}

func (handler *PostHandler) GetUsersWhoLikedPost(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")
	vars := mux.Vars(r)
	postid := vars["id"]
	if postid == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	var postuuid, err = ParseUUID(postid)
	if err != nil{
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	userids,_ := handler.Service.GetUsersWhoLikedPost(postuuid)

	if userids == nil{
		w.WriteHeader(http.StatusNotFound)
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(userids)

}

func (handler *PostHandler) GetUsersWhoDislikedPost(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")
	vars := mux.Vars(r)
	postid := vars["id"]
	if postid == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	var postuuid, err = ParseUUID(postid)
	if err != nil{
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	userids,_ := handler.Service.GetUsersWhoDislikedPost(postuuid)

	if userids == nil{
		w.WriteHeader(http.StatusNotFound)
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(userids)

}

func (handler *PostHandler) GetImageOld(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "image/*")
	vars := mux.Vars(r)
	imagepath := vars["id"]
	if imagepath == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	img,err := handler.Service.GetImage(imagepath)
	if err != nil{
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if img == nil{
		w.WriteHeader(http.StatusNotFound)
		return
	}

	buffer := new(bytes.Buffer)
	if err := jpeg.Encode(buffer, img, nil); err != nil {
		fmt.Println("Unable to encode image!")
	}

	w.Header().Set("Content-Length", strconv.Itoa(len(buffer.Bytes())))
	if _, err := w.Write(buffer.Bytes()); err != nil {
		fmt.Println("Unable to write image.")
	}

}

func (handler *PostHandler) UploadImage(w http.ResponseWriter,r *http.Request){
	vars := mux.Vars(r)
	imagePath := vars["id"]
	formKey := vars["formKey"]
	r.ParseMultipartForm(10 << 20)
	file, _, err := r.FormFile(formKey)
	if err != nil {
		fmt.Println("Error Retrieving the File")
		fmt.Println(err)
		return
	}

	defer file.Close()

	dst, err := os.Create("post-documents/" + imagePath +".jpg")
	defer dst.Close()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if _, err := io.Copy(dst, file); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	fmt.Fprintf(w, "Successfully Uploaded File\n")


}

func (handler *PostHandler) GetImage(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "image/jpeg")
	vars := mux.Vars(r)
	imagepath := vars["id"]
	if imagepath == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	file,err := ioutil.ReadFile("post-documents/" + imagepath)
	if err != nil{
		w.WriteHeader(http.StatusBadRequest)
		fmt.Println(err)
		return
	}
	w.Write(file)


}

func (handler *PostHandler) UploadVideo(w http.ResponseWriter,r *http.Request){
	vars := mux.Vars(r)
	imagePath := vars["videoId"]
	formKey := vars["formKey"]
	r.ParseMultipartForm(10 << 20)
	file, _, err := r.FormFile(formKey)
	if err != nil {
		fmt.Println("Error Retrieving the File")
		fmt.Println(err)
		return
	}

	defer file.Close()

	dst, err := os.Create("post-documents/" + imagePath +".mp4")
	defer dst.Close()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if _, err := io.Copy(dst, file); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	fmt.Println("Successfully Uploaded File\n" + dst.Name())

	w.WriteHeader(http.StatusOK)


}

func (handler *PostHandler) GetVideo(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "video/mp4")
	vars := mux.Vars(r)
	imagepath := vars["videoId"]
	if imagepath == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	file,err := ioutil.ReadFile("post-documents/" + imagepath)
	if err != nil{
		w.WriteHeader(http.StatusBadRequest)
		fmt.Println(err)
		return
	}
	w.Write(file)


}


func (handler *PostHandler) GetLikedPostsForUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")
	vars := mux.Vars(r)
	userid := vars["id"]
	if userid == ""{
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	likedPosts,_ := handler.Service.GetLikedPostsForUser(userid)

	if likedPosts == nil{
		w.WriteHeader(http.StatusNotFound)
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(likedPosts)

}

func (handler *PostHandler) GetDislikedPostsForUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")
	vars := mux.Vars(r)
	userid := vars["id"]
	if userid == ""{
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	likedPosts,_ := handler.Service.GetDislikedPostsForUser(userid)

	if likedPosts == nil{
		w.WriteHeader(http.StatusNotFound)
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(likedPosts)
}

func (handler *PostHandler) ReportContent(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var reportedConted DTO.ReportedContentDTO
	err := json.NewDecoder(r.Body).Decode(&reportedConted)

	if err != nil{
		fmt.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	err = handler.Service.ReportContent(&reportedConted)

	if err != nil{
		fmt.Println(err)
		w.WriteHeader(http.StatusExpectationFailed)
		return
	}

	w.WriteHeader(http.StatusCreated)

}

func (handler *PostHandler) GetCollectionsForUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")
	vars := mux.Vars(r)
	userid := vars["id"]
	if userid == ""{
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	collections,_ := handler.Service.GetCollectionsForUser(userid)

	if collections == nil{
		w.WriteHeader(http.StatusNotFound)
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(collections)
}

func (handler *PostHandler) CheckIfPostExistsInFavourites(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")
	vars := mux.Vars(r)
	userid := vars["id"]
	postid := vars["post"]
	if userid == "" || postid == ""{
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	var postuuid, err = ParseUUID(postid)
	if err != nil{
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	exist := handler.Service.CheckIfPostExistsInFavourites(userid,postuuid)
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(exist)
}

func (handler *PostHandler) GetAllCollectionsForPostByUser(w http.ResponseWriter,r  *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")
	vars := mux.Vars(r)
	userid := vars["id"]
	postid := vars["post"]
	if userid == "" || postid == ""{
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	var postuuid, err = ParseUUID(postid)
	if err != nil{
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	collections,_ := handler.Service.GetAllCollectionsForPostByUser(userid,postuuid)
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(collections)
}

func (handler *PostHandler) GetAllPostFeedsForUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	userId := mux.Vars(r)["userId"]


	if userId == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	posts,err := handler.Service.GetAllPostFeedsForUser(userId)
	if err != nil {
		w.WriteHeader(http.StatusExpectationFailed)
		return
	}

	json.NewEncoder(w).Encode(posts)

	w.WriteHeader(http.StatusOK)

}

func (handler *PostHandler) GetLocationForPost(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	postId := mux.Vars(r)["postId"]

	if postId == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	postuuid, err := ParseUUID(postId)
	if err != nil{
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	location,err := handler.Service.GetLocationForPost(postuuid)
	if err != nil {
		w.WriteHeader(http.StatusExpectationFailed)
		return
	}

	json.NewEncoder(w).Encode(location)

	w.WriteHeader(http.StatusOK)

}

func (handler *PostHandler) GetTagSuggestions(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	tag := mux.Vars(r)["tag"]

	if tag == "" {
		tag = "#"
	}else{
		tag = "#" + tag
	}

	tags, err := handler.Service.GetTagSuggestions(tag)
	if err != nil {
		w.WriteHeader(http.StatusExpectationFailed)
		return
	}

	json.NewEncoder(w).Encode(tags)

	w.WriteHeader(http.StatusOK)

}

func (handler *PostHandler) GetAllTags(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	tag := "#"

	tags, err := handler.Service.GetTagSuggestions(tag)
	if err != nil {
		w.WriteHeader(http.StatusExpectationFailed)
		return
	}

	json.NewEncoder(w).Encode(tags)

	w.WriteHeader(http.StatusOK)

}

func (handler *PostHandler) GetLocationSuggestions(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	location := mux.Vars(r)["location"]

	if location == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	locations, err := handler.Service.GetLocationSuggestions(location)
	if err != nil {
		w.WriteHeader(http.StatusExpectationFailed)
		return
	}
	json.NewEncoder(w).Encode(locations)
	w.WriteHeader(http.StatusOK)

}

func (handler *PostHandler) GetAllReportedContents(w http.ResponseWriter,r  *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	contents,err := handler.Service.GetAllReportedContents()
	if err != nil {
		w.WriteHeader(http.StatusExpectationFailed)
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(contents)
}

func (handler *PostHandler) DeleteReportedContent(w http.ResponseWriter,r  *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	vars := mux.Vars(r)
	stringContentId := vars["contentId"]
	userId := vars["contentId"]

	if stringContentId == "" || userId == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	var contentId, err = ParseUUID(stringContentId)
	if err != nil{
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	err = handler.Service.DeleteReportContent(contentId,userId)
	if err != nil {
		w.WriteHeader(http.StatusExpectationFailed)
		return
	}
	w.WriteHeader(http.StatusOK)
}

func (handler *PostHandler) DeletePost(w http.ResponseWriter,r  *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	vars := mux.Vars(r)
	stringPostId := vars["postId"]
	userId := vars["contentId"]
	if stringPostId == "" || userId == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	var postId, err = ParseUUID(stringPostId)
	if err != nil{
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	err = handler.Service.DeletePost(postId,userId)
	if err != nil {
		w.WriteHeader(http.StatusExpectationFailed)
		return
	}
	w.WriteHeader(http.StatusOK)
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
