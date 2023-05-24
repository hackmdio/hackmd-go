package internal

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
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

func GetConfigFilePath() string {
	return filepath.Join(GetConfigFileDir(), "config.json")
}

const DEFAULT_API_ENDPOINT_URL = "https://api.hackmd.io/v1"

func LoadConfig() {
	viper.SetConfigType("json")
	viper.SetConfigName("config")

	viper.SetDefault("hackmdAPIEndpointURL", DEFAULT_API_ENDPOINT_URL)

	viper.BindEnv("hackmdAPIEndpointURL", "HMD_API_ENDPOINT_URL")
	viper.BindEnv("accessToken", "HMD_API_ACCESS_TOKEN")

	configFilePath := GetConfigFileDir()
	viper.AddConfigPath(configFilePath)

	err := viper.ReadInConfig()
	if err != nil {
		println("Error reading config file: %s\n", err.Error())
	}
}

// Here we do the serizliaztion of the config file manually
// Because viper is case insensitive, and we want to preserve the case
func SaveConfigToFile() {
	LoadConfig()

	configFilePath := GetConfigFilePath()

	if _, err := os.Stat(configFilePath); os.IsNotExist(err) {
		os.MkdirAll(GetConfigFileDir(), os.ModePerm)
	}

	// build the config object manually
	config := map[string]interface{}{
		"hackmdAPIEndpointURL": viper.GetString("hackmdAPIEndpointURL"),
		"accessToken": viper.GetString("accessToken"),
	}

	/// Serialize the map object to JSON
	jsonBytes, err := json.MarshalIndent(config, "", "  ")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	// Convert the JSON bytes to a string
	jsonString := string(jsonBytes)

	// Write the string to a file
	err = ioutil.WriteFile(configFilePath, []byte(jsonString), 0644)

	if err != nil {
		panic(fmt.Errorf("Fatal error saving config file: %s\n", err))
	}
}

func IsAccessTokenPresent() bool {
	LoadConfig()

	accessToken := viper.GetString("accessToken")

	return accessToken != ""
}

func RemoveAccessTokenFromConfig() {
	LoadConfig()

	viper.Set("accessToken", "")
	SaveConfigToFile()
}

