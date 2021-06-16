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
	err = handler.Service.Create(&postDTO)
	if err != nil{
		fmt.Println(err)
		w.WriteHeader(http.StatusExpectationFailed)
		return
	}

	w.WriteHeader(http.StatusCreated)

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
	}

	posts,_ := handler.Service.FindPostsByTag(tag)

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
	w.Header().Set("Content-Type", "image/jpeg")
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
	r.ParseMultipartForm(10 << 20)
	file, _, err := r.FormFile("myFile")
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

/*func (handler *PostHandler) GetAllLikesForPost(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Ulaziii")
	w.Header().Set("Content-Type", "application/json")
	postid := mux.Vars(r)["id"]
	i, err := strconv.Atoi(postid)
	fmt.Println("Link dobar ", i)
		if err != nil {
			fmt.Println(err)
			w.WriteHeader(http.StatusBadRequest)
			return
		}
	err2 := handler.Service.GetAllLikesForPost(postid)
	if err2 != nil{
	fmt.Println(err2)
	w.WriteHeader(http.StatusExpectationFailed)
	return
	}

	w.WriteHeader(http.StatusCreated)

}*/

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
