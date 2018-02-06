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
	// config := oauth1.NewConfig("6lHa4wdDL3Nt7Ijd4DAQkP627", "NoBFqKhoAMMonF3BTxvxjm1ckhiHF9A9tlaDuA6ejmMtI9k154")
	// token := oauth1.NewToken("960391604170207234-Ta3bAgH5gK8jf8EyeieNFl5XkRL6m62", "TdUoLU7PKy92NXU4ndd5HnGsmL0Ia95JfmorB4omPMubf")
	httpClient := config.Client(oauth1.NoContext, token)

	// Twitter client
	client = twttr.NewClient(httpClient)
}

// GetClient ...
func GetClient() *twttr.Client {
	return client
}
