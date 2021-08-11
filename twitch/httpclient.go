package twitch

import (
	"fmt"
	"os"
	"phillzbot/domain"

	"github.com/Adeithe/go-twitch/api/helix"
	"github.com/Adeithe/go-twitch/api/kraken"
)

var (
	apiInstance *domain.TwitchAPI
)

func NewAPIClient() (*domain.TwitchAPI, error) {
	if apiInstance != nil {
		return apiInstance, nil
	}

	api_client_id := os.Getenv("API_CLIENT_ID")
	if api_client_id == "" {
		return nil, fmt.Errorf("error: domain.TwitchAPI client id not found")
	}

	api_token := os.Getenv("API_ACCESS_TOKEN")
	if api_token == "" {
		return nil, fmt.Errorf("error: domain.TwitchAPI token not found")
	}

	krakenAPI := kraken.New(api_client_id, api_token)
	helixAPI := helix.New(api_client_id, api_token)

	apiInstance = &domain.TwitchAPI{
		Kraken: krakenAPI,
		Helix:  helixAPI,
	}

	return apiInstance, nil
}
