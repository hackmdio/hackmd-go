package internal;

import (
	"fmt"
	"os"

	"github.com/spf13/viper"
)

func GetConfigFilePath() string {
  var homeDir, err = os.UserHomeDir()

  if err != nil {
    panic(fmt.Errorf("Fatal error getting home directory: %s\n", err))
  }

  return fmt.Sprintf("%s/.config/hackmd/config.json", homeDir)
}

func LoadConfig() {
	viper.SetConfigType("json")

	viper.SetDefault("hackmdAPIEndpointURL", "https://api.hackmd.io/v1")

	viper.BindEnv("hackmdAPIEndpointURL", "HMD_API_ENDPOINT_URL")
	viper.BindEnv("accessToken", "HMD_API_ACCESS_TOKEN")

	configFilePath := GetConfigFilePath()

	if _, err := os.Stat(configFilePath); err == nil {
		println("Using config file: " + configFilePath)
		viper.SetConfigFile(configFilePath)

		err := viper.ReadInConfig()
		if err != nil {
			panic(fmt.Errorf("Fatal error reading config file: %s\n", err))
		}
	} else {
		viper.AutomaticEnv()
	}
}

