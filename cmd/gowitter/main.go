package main

import (
	"fmt"
	"github.com/rsnullptr/gowitter/twitter"
	"log"
	"os"
)

func main() {
	fmt.Println("Go-Twitter Bot v0.01")
	creds := twitter.Credentials{
		AccessToken:       os.Getenv("ACCESS_TOKEN"),
		AccessTokenSecret: os.Getenv("ACCESS_TOKEN_SECRET"),
		ConsumerKey:       os.Getenv("CONSUMER_KEY"),
		ConsumerSecret:    os.Getenv("CONSUMER_KEY_SECRET"),
	}

	client, err := twitter.NewClient(&creds)
	if err != nil {
		log.Println("error getting Twitter Client")
		log.Println(err)
	}

	fmt.Printf("%+v\n", client)
}
