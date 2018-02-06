package configuration

import (
	"log"
	"sync"

	"github.com/spf13/viper"
)

// AppConfig ...
type AppConfig struct {
	ConsumerToken  string
	ConsumerSecret string
	AccessToken    string
	AccessSecret   string
	DataPath       string
}

var config AppConfig
var configOnce sync.Once

//GetAppConfig app server configuration
func GetAppConfig() AppConfig {
	return config
}

// NewAppConfig configuration setup
func NewAppConfig(appName string) AppConfig {
	// Set config file name: {appName}.json
	viper.SetConfigName("." + appName)
	viper.SetConfigType("json")

	// Tell Viper where to look for config
	viper.AddConfigPath("$HOME/")         // home dir
	viper.AddConfigPath("$HOME/.config/") // config dir
	viper.AddConfigPath("./")             // project root
	viper.AddConfigPath("../")            // for packages

	// Load up config
	err := viper.ReadInConfig()
	if err != nil {
		log.Print("Didn't find config file; falling back on ENV")
	}

	configOnce.Do(func() {
		config = AppConfig{
			ConsumerToken:  viper.GetString("TWITTER_CONSUMER_TOKEN"),
			ConsumerSecret: viper.GetString("TWITTER_CONSUMER_SECRET"),
			AccessToken:    viper.GetString("TWITTER_ACCESS_TOKEN"),
			AccessSecret:   viper.GetString("TWITTER_ACCESS_SECRET"),
			DataPath:       viper.GetString("DATA_PATH"),
		}
	})

	return config
}
