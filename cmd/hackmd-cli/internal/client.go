package internal;

import (
	HackMDClient "github.com/hackmdio/hackmd-go/pkg/api"
	"github.com/spf13/viper"
	"github.com/AlecAivazis/survey/v2"
)

func GetHackMDClient() *HackMDClient.APIClient {
  LoadConfig()

  accessToken := viper.GetString("accessToken")
  apiEndpoint := viper.GetString("hackmdAPIEndpointURL")

  if accessToken == "" {
    accessToken = askForAccessToken()
    viper.Set("accessToken", accessToken)
    SaveConfigToFile()
  }

  return HackMDClient.NewAPIClient(accessToken, HackMDClient.WithAPIEndpointURL(apiEndpoint))
}

func askForAccessToken() string {
  var accessToken string
  prompt := &survey.Password{
    Message: "Please enter your access token:",
  }
  survey.AskOne(prompt, &accessToken)

  return accessToken
}

