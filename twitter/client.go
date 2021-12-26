package twitter

import (
	"github.com/dghubble/go-twitter/twitter"
	"github.com/dghubble/oauth1"
	"net/http"
)

type Bot struct {
	*twitter.Client
}

type Search = twitter.Search
type Tweet = twitter.Tweet

// NewBot return bot instance
func NewBot(creds *Credentials) (Bot, error) {
	config := oauth1.NewConfig(creds.ConsumerKey, creds.ConsumerSecret)
	token := oauth1.NewToken(creds.AccessToken, creds.AccessTokenSecret)
	httpClient := config.Client(oauth1.NoContext, token)
	client := twitter.NewClient(httpClient)

	verifyParams := &twitter.AccountVerifyParams{
		SkipStatus:   twitter.Bool(true),
		IncludeEmail: twitter.Bool(true),
	}

	_, _, err := client.Accounts.VerifyCredentials(verifyParams)
	if err != nil {
		return Bot{}, err
	}

	return Bot{Client: client}, nil
}

func (b Bot) Tweet(tweetMsg string) (*Tweet, *http.Response, error) {
	return b.Client.Statuses.Update(tweetMsg, nil)
}

func (b Bot) Search(query string) (*Search, *http.Response, error) {
	return b.Client.Search.Tweets(&twitter.SearchTweetParams{
		Query: query,
	})
}
