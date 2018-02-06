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
		text = markov.MainChain.Generate()
		text = strings.Replace(text, "\n", "", -1)
		text = strings.Replace(text, "\r", "", -1)
		text = strings.Replace(text, "\t", "", -1)
		text = strings.Replace(text, "  ", " ", -1)

		if len(text) < 30 {
			fmt.Println("too short")
			continue
		}

		if len(text) > 280 {
			text = text[0:280]
		}

		if !strings.ContainsAny(text, ".!?;") {
			fmt.Println("No terminal punctuation")
			continue
		}

		last := strings.LastIndex(text, ".")
		if last > 0 {
			text = text[0 : last+1]
			fmt.Println(text)
			break
		}

		last = strings.LastIndex(text, "!")
		if last > 0 {
			text = text[0 : last+1]
			fmt.Println(text)
			break
		}

		last = strings.LastIndex(text, ";")
		if last > 0 {
			text = text[0 : last+1]
			fmt.Println(text)
			break
		}

		last = strings.LastIndex(text, "?")
		if last > 0 {
			text = text[0 : last+1]
			fmt.Println(text)
			break
		}
	}

	_, _, err := c.Statuses.Update(text, nil)
	if err != nil {
		panic(err)
	}
}

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
