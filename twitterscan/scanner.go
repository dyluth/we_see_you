package twitterscan

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/url"
	"os"

	"github.com/ChimeraCoder/anaconda"
)

// ScannerImpl is an implementation that talks to twitter
type ScannerImpl struct {
	api *anaconda.TwitterApi
}

type creds struct {
	ConsumerKey       string `json:"consumer_key"`
	ConsumerSecret    string `json:"consumer_secret"`
	AccessToken       string `json:"access_token"`
	AccessTokenSecret string `json:"access_token_secret"`
}

func load(credsPath string) (creds, error) {
	credsFile, err := os.Open(credsPath)
	if err != nil {
		return creds{}, err
	}
	defer credsFile.Close()
	byteValue, err := ioutil.ReadAll(credsFile)
	if err != nil {
		return creds{}, err
	}

	var c creds
	err = json.Unmarshal(byteValue, &c)
	return c, err
}

func Init(credsPath string) (ScannerImpl, error) {
	creds, err := load(credsPath)
	if err != nil {
		return ScannerImpl{}, err
	}

	anaconda.SetConsumerKey(creds.ConsumerKey)
	anaconda.SetConsumerSecret(creds.ConsumerSecret)
	api := anaconda.NewTwitterApi(creds.AccessToken, creds.AccessTokenSecret)
	scanner := ScannerImpl{api: api}

	return scanner, nil
}

func (scanner *ScannerImpl) scan() {
	v := url.Values{}
	s := scanner.api.UserStream(v)

	for t := range s.C {
		switch v := t.(type) {
		case anaconda.Tweet:
			fmt.Printf("%-15s: %s\n", v.User.ScreenName, v.Text)
		case anaconda.EventTweet:
			switch v.Event.Event {
			case "favorite":
				sn := v.Source.ScreenName
				tw := v.TargetObject.Text
				fmt.Printf("Favorited by %-15s: %s\n", sn, tw)
			case "unfavorite":
				sn := v.Source.ScreenName
				tw := v.TargetObject.Text
				fmt.Printf("UnFavorited by %-15s: %s\n", sn, tw)
			}
		}
	}
}
