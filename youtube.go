/*
This is the YouTube package level comment.

*/

package youtube

import (
	"context"
	"fmt"

	"google.golang.org/api/option"
	"google.golang.org/api/youtube/v3"
)

// InitYouTubeAPI creates YouTube API Service with given API key
func InitYouTubeAPI(apiKey string) (*youtube.Service, error) {
	ctx := context.Background()
	youtubeService, err := youtube.NewService(ctx, option.WithAPIKey(apiKey))

	if err != nil {
		return nil, fmt.Errorf("InitYouTubeAPI: %v", err)
	}
	return youtubeService, nil
}

// ChannelById returns the information (parts) of channel (id)
func ChannelById(service *youtube.Service, part []string, id string) (*youtube.Channel, error) {

	channelResponse, err := service.Channels.List(part).Id(id).Do()

	if err != nil {
		return nil, fmt.Errorf("ChannelById: %v", err)
	}

	if len(channelResponse.Items) == 0 {
		return nil, fmt.Errorf("ChannelById: Empty channelResponse.Items")
	}

	channel := channelResponse.Items[0]

	return channel, nil
}
