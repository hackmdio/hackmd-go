package internal;

import (
	HackMDClient "github.com/hackmdio/hackmd-go/pkg/api"
	"github.com/spf13/viper"
)

func GetHackMDClient() *HackMDClient.APIClient {
  LoadConfig()

  accessToken := viper.GetString("accessToken")
  apiEndpoint := viper.GetString("hackmdAPIEndpointURL")

  return HackMDClient.NewAPIClient(accessToken, HackMDClient.WithAPIEndpointURL(apiEndpoint))
}

