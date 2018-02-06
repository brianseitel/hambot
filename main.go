package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	"github.com/brianseitel/hambot/configuration"
	"github.com/brianseitel/hambot/markov"
	"github.com/brianseitel/hambot/twitter"
)

func main() {
	c := twitter.GetClient()

	markov.Init()

	papers := loadData()
	for _, m := range papers {
		markov.MainChain.Build(strings.Trim(m+".", " \t\n\r"))
	}

	var text string
	for {
		text := markov.MainChain.Generate()
		text = strings.Replace(text, "\n", "", -1)
		text = strings.Replace(text, "\r", "", -1)
		text = strings.Replace(text, "\t", "", -1)
		text = strings.Replace(text, "  ", " ", -1)

		// Don't do anything less than 30 chars
		if len(text) < 30 {
			fmt.Println("too short")
			continue
		}

		// Cap at 200 characters or so
		if len(text) > 280 {
			text = text[0:280]
		}

		// Detect whether to terminate or not
		if !strings.ContainsAny(text, ".!?;") {
			fmt.Println("No terminal punctuation")
			continue
		}

		// Find the end of the phrase, then break
		text, end := detectEnd(text)
		if end {
			break
		}
	}

	// Send it!
	_, _, err := c.Statuses.Update(text, nil)
	if err != nil {
		panic(err)
	}
}

// Detect the end of a segment
func detectEnd(text string) (string, bool) {
	ends := []string{".", "!", ";", "?"}
	for _, s := range ends {
		last := strings.LastIndex(text, s)
		if last > 0 {
			text = text[0 : last+1]
			return text, true
		}
	}

	return text, false
}

// Load data from file
func loadData() []string {
	config := configuration.GetAppConfig()

	f, err := os.Open(config.DataPath)
	if err != nil {
		panic(err)
	}

	body, err := ioutil.ReadAll(f)
	if err != nil {
		panic(err)
	}

	return strings.Split(string(body), ". ")
}
