package internal;

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/spf13/viper"
)

func GetConfigFileDir() string {
  var homeDir, err = os.UserHomeDir()

  if err != nil {
    panic(fmt.Errorf("Fatal error getting home directory: %s\n", err))
  }

	return filepath.Join(homeDir, ".hackmd")
}

func LoadConfig() {
	viper.SetConfigType("json")

	viper.SetDefault("hackmdAPIEndpointURL", "https://api.hackmd.io/v1")

	viper.BindEnv("hackmdAPIEndpointURL", "HMD_API_ENDPOINT_URL")
	viper.BindEnv("accessToken", "HMD_API_ACCESS_TOKEN")

	configFilePath := GetConfigFileDir()
	viper.AddConfigPath(configFilePath)

	err := viper.ReadInConfig()
	if err != nil {
		println("Error reading config file: %s\n", err.Error())
	}
}

