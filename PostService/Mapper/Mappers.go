package Mapper

import (
	"XWS-Nistagram/PostService/DTO"
	"XWS-Nistagram/PostService/Model"
)

func ConvertPostDTOToPost(postDTO *DTO.PostDTO) *Model.Post {
	var post Model.Post
	post.ID = postDTO.ID
	post.CreatedAt = postDTO.CreatedAt
	post.Description = postDTO.Description
	post.DislikesCount = postDTO.DislikesCount
	post.LikesCount = postDTO.LikesCount
	post.Image = postDTO.Image
	post.UserID = postDTO.UserID
	post.Comments = postDTO.Comments
	return &post
}

func ConvertCommentDTOToComment(commentDTO *DTO.CommentDTO) *Model.Comment {
	var comment Model.Comment
	comment.ID = commentDTO.ID
	comment.UserID = commentDTO.UserID
	comment.PostID = commentDTO.PostID
	comment.CreatedAt = commentDTO.CreatedAt
	comment.Content = commentDTO.Content
	return &comment
}