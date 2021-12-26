package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/rsnullptr/gowitter/twitter"
	"log"
	"os"
	"strconv"
	"time"
)

type configuration struct {
	Tweets []string `json:"tweets"`
}

func main() {
	fmt.Println("gowitter v0.0.1")

	file, err := os.ReadFile(os.Getenv("TWEETS_JSON"))
	if err != nil {
		panic(err)
	}
	var cfg configuration
	err = json.Unmarshal(file, &cfg)
	if err != nil {
		panic(err)
	}

	intervalStr := os.Getenv("INTERVAL_MIN")
	if intervalStr == "" {
		intervalStr = "45"
	}

	interval, err := strconv.ParseInt(intervalStr, 10, 64)
	if err != nil {
		panic(err)
	}

	creds := twitter.Credentials{
		AccessToken:       os.Getenv("ACCESS_TOKEN"),
		AccessTokenSecret: os.Getenv("ACCESS_TOKEN_SECRET"),
		ConsumerKey:       os.Getenv("CONSUMER_KEY"),
		ConsumerSecret:    os.Getenv("CONSUMER_KEY_SECRET"),
	}

	client, err := twitter.NewBot(&creds)
	if err != nil {
		panic(err)
	}

	if len(cfg.Tweets) == 0 {
		panic(errors.New("no tweets provided"))
	}

	chrono := time.NewTicker(time.Minute * time.Duration(interval))
	i := 0
	for {
		select {
		case <-chrono.C:
			_, _, err := client.Tweet(cfg.Tweets[i])
			if err != nil {
				log.Printf("unable to tweet index %d; err: %s", i, err.Error())
			}

			log.Printf("tweeted index %d with success.", i)

			i++
			if i >= len(cfg.Tweets) {
				log.Printf("index is %d, reseting tweet index.", i)
				i = 0
			}
		}
	}

}
