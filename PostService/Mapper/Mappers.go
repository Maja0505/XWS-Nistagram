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
	post.Media = postDTO.Media
	post.UserID = postDTO.UserID
	post.CommentsCount = postDTO.CommentsCount
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

func ConvertReportedContentDTOToReportedContent(reportedContentDTO *DTO.ReportedContentDTO) *Model.ReportedContent{
	var reportedContent Model.ReportedContent
	reportedContent.ID = reportedContentDTO.ID
	reportedContent.Description = reportedContentDTO.Description
	reportedContent.UserID = reportedContentDTO.UserID
	reportedContent.AdminID = reportedContentDTO.AdminID
	reportedContent.ContentID = reportedContentDTO.ContentID
	return &reportedContent
}

func ConvertCreateStoryDTOToPost(storyDTO *DTO.CreateStoryDTO) *Model.Story {
	var story Model.Story
	story.Image = storyDTO.Image
	story.UserID = storyDTO.UserID
	story.ForCloseFriends = storyDTO.ForCloseFriends
	story.Highlights = storyDTO.Highlights
	return &story
}
