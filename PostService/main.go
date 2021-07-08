package main

import (
	"XWS-Nistagram/PostService/Handler"
	"XWS-Nistagram/PostService/Repository"
	"XWS-Nistagram/PostService/Service"
	"fmt"
	"github.com/gocql/gocql"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"log"
	"net/http"
	"os"
)

var Session *gocql.Session

func init() {
	fmt.Println("Initialization of cassandra...")
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatal("Error loading .env file")
	}
	cluster := gocql.NewCluster(os.Getenv("POST_SERVICE_HOST"))
	//cluster.ProtoVersion = 4
	//cluster.Keyspace = "postkeyspace"

	Session, err = cluster.CreateSession()
	if err != nil {
		panic(err)
	}
	if err := Session.Query("create keyspace  if not exists postkeyspace with replication = {'class':'SimpleStrategy','replication_factor':1};").Exec(); err != nil {
		fmt.Println("Error while inserting postkeyspace")
		fmt.Println(err)
	}

	if err := Session.Query("DROP TABLE postkeyspace.stories").Exec(); err != nil {
		fmt.Println("Error while dropping table!")
		fmt.Println(err)
	}

	if err := Session.Query("CREATE TABLE if not exists postkeyspace.posts(id timeuuid, userid text, description text, media list<text>, album boolean, repeatcampaign boolean, createdat timestamp, PRIMARY KEY((userid), id)) WITH CLUSTERING ORDER BY (id DESC);").Exec(); err != nil {
		fmt.Println("Error while creating tables!")
		fmt.Println(err)
	}

	if err := Session.Query("CREATE TABLE if not exists postkeyspace.locations(postid uuid, location text, PRIMARY KEY((location), postid));").Exec(); err != nil {
		fmt.Println("Error while creating tables!")
		fmt.Println(err)
	}

	if err := Session.Query("CREATE TABLE if not exists postkeyspace.postcounters(postid uuid, likes counter, dislikes counter, comments counter, media counter, PRIMARY KEY(postid));").Exec(); err != nil {
		fmt.Println("Error while creating tables!")
		fmt.Println(err)
	}

	if err := Session.Query("CREATE TABLE if not exists postkeyspace.comments(id timeuuid, postid uuid, userid text, content text, PRIMARY KEY((postid), id, userid)) WITH CLUSTERING ORDER BY (id DESC);").Exec(); err != nil {
		fmt.Println("Error while creating tables!")
		fmt.Println(err)
	}
	if err := Session.Query("CREATE TABLE if not exists postkeyspace.likes(postid uuid, userid text, PRIMARY KEY((postid, userid)));").Exec(); err != nil {
		fmt.Println("Error while creating tables!")
		fmt.Println(err)
	}
	if err := Session.Query("CREATE TABLE if not exists postkeyspace.dislikes(postid uuid, userid text, PRIMARY KEY((postid, userid)));").Exec(); err != nil {
		fmt.Println("Error while creating tables!")
		fmt.Println(err)
	}
	if err := Session.Query("CREATE TABLE if not exists postkeyspace.tags(postid uuid, tag text, PRIMARY KEY((postid), tag));").Exec(); err != nil {
		fmt.Println("Error while creating tables!")
		fmt.Println(err)
	}
	if err := Session.Query("CREATE TABLE if not exists postkeyspace.tagsDK(postid uuid, tag text, PRIMARY KEY((tag), postid));").Exec(); err != nil {
		fmt.Println("Error while creating tables!")
		fmt.Println(err)
	}
	if err := Session.Query("CREATE TABLE if not exists postkeyspace.favourites(userid text, postid uuid, PRIMARY KEY((userid), postid));").Exec(); err != nil {
		fmt.Println("Error while creating tables!")
		fmt.Println(err)
	}
	if err := Session.Query("CREATE TABLE if not exists postkeyspace.collections(userid text, postid uuid, collection text, PRIMARY KEY((userid), collection, postid));").Exec(); err != nil {
		fmt.Println("Error while creating tables!")
		fmt.Println(err)
	}

	if err := Session.Query("CREATE TABLE if not exists postkeyspace.reported_contents(id uuid, description text, contentid text, userid text, adminid text, PRIMARY KEY((userid, id)));").Exec(); err != nil {
		fmt.Println("Error while creating tables!")
		fmt.Println(err)
	}

	if err := Session.Query("CREATE TABLE if not exists postkeyspace.stories(id timeuuid, userid text, available boolean, image text, highlights boolean, for_close_friends boolean, createdat timestamp, PRIMARY KEY((userid), id)) WITH CLUSTERING ORDER BY (id DESC);").Exec(); err != nil {
		fmt.Println("Error while creating tables!")
		fmt.Println(err)
	}

	fmt.Println("Cassandra well initialized!")
}


func initPostRepo(session *gocql.Session) *Repository.PostRepository{
	return &Repository.PostRepository{Session: *session}
}

func initStoryRepo(session *gocql.Session) *Repository.StoryRepository{
	return &Repository.StoryRepository{Session: *session}
}

func initPostService(postRepo *Repository.PostRepository) *Service.PostService{
	return &Service.PostService{Repo : *postRepo}
}

func initStoryService(storyRepo *Repository.StoryRepository) *Service.StoryService{
	return &Service.StoryService{Repo : *storyRepo}
}

func initHandler(service *Service.PostService) *Handler.PostHandler{
	return &Handler.PostHandler{Service: service}
}

func initStoryHandler(service *Service.StoryService) *Handler.StoryHandler{
	return &Handler.StoryHandler{Service: service}
}


func handleFunc(handler *Handler.PostHandler,router *mux.Router){

	router.HandleFunc("/create", handler.Create).Methods("POST")
	router.HandleFunc("/add-comment", handler.AddComment).Methods("POST")
	router.HandleFunc("/delete-comment", handler.DeleteComment).Methods("POST")
	router.HandleFunc("/like-post", handler.LikePost).Methods("POST")
	router.HandleFunc("/dislike-post", handler.DislikePost).Methods("POST")
	router.HandleFunc("/add-to-favourites", handler.AddPostToFavourites).Methods("POST")
	router.HandleFunc("/add-to-collection", handler.AddPostToCollection).Methods("POST")
	router.HandleFunc("/remove-post-from-favourites", handler.RemovePostFromFavourites).Methods("POST")
	router.HandleFunc("/remove-post-from-collection", handler.RemovePostFromCollection).Methods("POST")
	router.HandleFunc("/get-one-post/{id}", handler.FindPostById).Methods("GET")
	router.HandleFunc("/get-all-by-userid/{id}", handler.FindPostsByUserId).Methods("GET")
	router.HandleFunc("/get-comments-for-post/{id}", handler.GetCommentsForPost).Methods("GET")
	router.HandleFunc("/get-users-tagged-on-post/{id}", handler.GetUsersTaggedOnPost).Methods("GET")
	router.HandleFunc("/get-users-who-liked-post/{id}", handler.GetUsersWhoLikedPost).Methods("GET")
	router.HandleFunc("/get-users-who-disliked-post/{id}", handler.GetUsersWhoDislikedPost).Methods("GET")
	router.HandleFunc("/get-user-who-posted-comment/{id}", handler.GetUserWhoPostedComment).Methods("GET")
	router.HandleFunc("/get-image/{id}", handler.GetImage).Methods("GET")
	router.HandleFunc("/add-tag", handler.AddTag).Methods("POST")
	router.HandleFunc("/get-tags-for-post/{id}", handler.GetTagsForPost).Methods("GET")
	router.HandleFunc("/get-pure-tags-for-post/{id}", handler.GetPureTagsForPost).Methods("GET")
	router.HandleFunc("/get-all-by-tag/{tag}", handler.FindPostsByTag).Methods("GET")
	router.HandleFunc("/get-all-by-location/{location}", handler.FindPostsByLocation).Methods("GET")
	router.HandleFunc("/get-favourite-posts/{id}", handler.GetFavouritePosts).Methods("GET")
	router.HandleFunc("/get-posts-from-collection/{id}/{collection}", handler.GetPostsFromCollection).Methods("GET")
	router.HandleFunc("/upload-image/{id}/{formKey}",handler.UploadImage).Methods("POST")
	router.HandleFunc("/like-exists", handler.CheckIfLikeExists).Methods("PUT")
	router.HandleFunc("/dislike-exists", handler.CheckIfDislikeExists).Methods("PUT")
	router.HandleFunc("/get-liked-posts-for-user/{id}", handler.GetLikedPostsForUser).Methods("GET")
	router.HandleFunc("/get-disliked-posts-for-user/{id}", handler.GetDislikedPostsForUser).Methods("GET")
	router.HandleFunc("/report-content", handler.ReportContent).Methods("POST")
	router.HandleFunc("/video-upload/{videoId}/{formKey}", handler.UploadVideo).Methods("POST")
	router.HandleFunc("/video-get/{videoId}", handler.GetVideo).Methods("GET")
	router.HandleFunc("/get-collections-for-user/{id}", handler.GetCollectionsForUser).Methods("GET")
	router.HandleFunc("/post-exists-in-favourites/{id}/{post}", handler.CheckIfPostExistsInFavourites).Methods("GET")
	router.HandleFunc("/get-all-collections-for-post-by-user/{id}/{post}", handler.GetAllCollectionsForPostByUser).Methods("GET")
	router.HandleFunc("/get-all-post-feeds-for-user/{userId}", handler.GetAllPostFeedsForUser).Methods("GET")
	router.HandleFunc("/get-tag-suggestions/{tag}", handler.GetTagSuggestions).Methods("GET")
	router.HandleFunc("/get-all-tags", handler.GetAllTags).Methods("GET")
	router.HandleFunc("/get-location-for-post/{postId}", handler.GetLocationForPost).Methods("GET")
	router.HandleFunc("/get-location-suggestions/{location}", handler.GetLocationSuggestions).Methods("GET")
	router.HandleFunc("/update-createdat", handler.UpdatePostCreatedAt).Methods("POST")
}

func handleStoryFunc(handler *Handler.StoryHandler,router *mux.Router){

	router.HandleFunc("/story/create", handler.Create).Methods("POST")
	router.HandleFunc("/story/set-for-highlights/{id}", handler.SetStoryForHighlights).Methods("PUT")

	router.HandleFunc("/story/all/{userId}", handler.GetAllStoriesByUser).Methods("GET")
	router.HandleFunc("/story/all-not-expired/{userId}", handler.GetAllNotExpiredStoriesByUser).Methods("GET")
	router.HandleFunc("/story/all-for-close-friends/{userId}", handler.GetAllStoriesForCloseFriendsByUser).Methods("GET")
	router.HandleFunc("/story/all-highlights/{userId}", handler.GetAllHighlightsStoriesByUser).Methods("GET")
	router.HandleFunc("/story/all-follows-with-stories/{userId}", handler.GetAllFollowsWithStories).Methods("GET")

	router.HandleFunc("/story/video-upload/{videoId}", handler.UploadVideo).Methods("POST")
	router.HandleFunc("/story/video-get/{videoId}", handler.GetVideo).Methods("GET")
	router.HandleFunc("/story/image-upload/{imageId}", handler.UploadImage).Methods("POST")

	router.HandleFunc("/story/update-agent", handler.UpdateStoryAvailabilityAndDate).Methods("POST")

}

func main(){
	fmt.Println("\n----------------MAIN----------------\n")
	postRepo := initPostRepo(Session)
	postRepo.InitTrie()
	postRepo.InitLocationsTrie()
	postService := initPostService(postRepo)
	handler := initHandler(postService)


	storyRepo := initStoryRepo(Session)
	storyService := initStoryService(storyRepo)
	storyHandler := initStoryHandler(storyService)

	router := mux.NewRouter().StrictSlash(true)
	handleFunc(handler,router)
	handleStoryFunc(storyHandler,router)

	headers := handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization"})
	methods := handlers.AllowedMethods([]string{"GET", "POST", "PUT", "HEAD", "OPTIONS"})
	origins := handlers.AllowedOrigins([]string{"*"})

	fmt.Println("\nServer running...")
	log.Fatal(http.ListenAndServe(":" + os.Getenv("POST_SERVICE_PORT"), handlers.CORS(headers, methods, origins)(router)))

}