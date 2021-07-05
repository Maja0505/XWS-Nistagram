package Mapper

import (
	"XWS-Nistagram/PostService/DTO"
	"XWS-Nistagram/PostService/Model"
	"strconv"
	"time"
)

func ConvertPostDTOToPost(postDTO *DTO.PostDTO) *Model.Post {
	var post Model.Post
	post.CreatedAt = time.Now()
	post.Description = postDTO.Description
	post.DislikesCount = 0
	post.LikesCount = 0
	post.Media = postDTO.Media
	post.UserID = postDTO.UserID
	post.CommentsCount = 0
	post.Location = postDTO.Location
	post.ID = postDTO.ID
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

func ConvertStoryListToStoryDTOList(stories *[]Model.Story) *[]DTO.StoryDTO{
	var storiesDtos []DTO.StoryDTO

	for _, story := range *stories {
		storiesDtos = append(storiesDtos,*convertStoryToStoryDTO(&story))
	}

	return &storiesDtos
}

func convertStoryToStoryDTO(s *Model.Story) *DTO.StoryDTO {
	var storyDto DTO.StoryDTO

	storyDto.ID = s.ID
	storyDto.UserID = s.UserID
	storyDto.Duration = 10
	storyDto.ForCloseFriends = s.ForCloseFriends
	storyDto.Highlights = s.Highlights
	storyDto.Media = s.Image
	if s.Image[len(s.Image) - 3 : len(s.Image)] == "jpg" {
		storyDto.Type = "image"
	}else{
		storyDto.Type = "video"
	}
	storyDto.Subheading = "Posted " + calculateMinutes(s.CreatedAt) + " ago"

	return &storyDto
}

func calculateMinutes(createdAt time.Time) string {
	var stringDuration string
	now := time.Now()
	durationMinutes := now.Minute() - createdAt.Minute()
	durationHours := now.Hour() - createdAt.Hour()
	if durationHours > 0{
		stringDuration = strconv.Itoa(durationHours) + "h"
	}else{
		stringDuration = strconv.Itoa(durationMinutes) + "min"
	}

	return stringDuration

}