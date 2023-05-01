package internal;

import (
	HackMDClient "github.com/hackmdio/hackmd-go/pkg/api"
	"github.com/spf13/viper"
)

func GetHackMDClient() *HackMDClient.APIClient {
  LoadConfig()


  var accessToken string = viper.GetString("accessToken")
  var apiEndpoint string = viper.GetString("hackmdAPIEndpointURL")

  return HackMDClient.NewAPIClient(accessToken, HackMDClient.WithAPIEndpointURL(apiEndpoint))
}

