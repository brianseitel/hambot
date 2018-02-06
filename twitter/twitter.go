package twitter

import (
	"github.com/brianseitel/hambot/configuration"
	twttr "github.com/dghubble/go-twitter/twitter"
	"github.com/dghubble/oauth1"
)

var client *twttr.Client

func init() {
	appConfig := configuration.NewAppConfig("hambot")

	config := oauth1.NewConfig(appConfig.ConsumerToken, appConfig.ConsumerSecret)
	token := oauth1.NewToken(appConfig.AccessToken, appConfig.AccessSecret)
	httpClient := config.Client(oauth1.NoContext, token)

	// Twitter client
	client = twttr.NewClient(httpClient)
}

// GetClient ...
func GetClient() *twttr.Client {
	return client
}
