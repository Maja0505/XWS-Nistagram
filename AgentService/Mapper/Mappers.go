package Mapper

import (
	"XWS-Nistagram/AgentService/DTO"
	"XWS-Nistagram/AgentService/Model"
)

func ConvertCampaignDTOToCampaign(campaignDTO *DTO.CampaignDTO) *Model.Campaign {
	var campaign Model.Campaign
	campaign.DislikesCount = 0
	campaign.LikesCount = 0
	campaign.Media = campaignDTO.Media
	campaign.UserID = campaignDTO.UserID
	campaign.CommentsCount = 0
	campaign.ViewsCount = 0
	campaign.ID = campaignDTO.ID
	campaign.End = campaignDTO.End
	campaign.Start = campaignDTO.Start
	campaign.Links = campaignDTO.Links
	campaign.Repeat = campaignDTO.Repeat
	campaign.RepeatFactor = campaignDTO.RepeatFactor
	campaign.IsPost = campaignDTO.IsPost
	campaign.Location = campaignDTO.Location
	campaign.Description = campaignDTO.Description
	campaign.Tags = campaignDTO.Tags
	return &campaign
}
