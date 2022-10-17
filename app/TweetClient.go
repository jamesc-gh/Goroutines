package app

import (
	"bufio"
	"context"
	"errors"
	"fmt"
	"math/rand"
	"net/http"
	"time"
)

type TweetClienter interface {
	GetTweetsFromStream(responses chan<- TweetResponse) error
}

const (
	// enter your bearertoken for the twitter v2 API here
	BearerToken      = "abc"
	ApiUrl           = "https://api.twitter.com/2/tweets/sample/stream"
	MinPauseDuration = 100
	MaxPauseDuration = 10000
)

type tweetClient struct {
	ctx context.Context
}

func NewTweetClient(ctx context.Context) TweetClienter {
	return &tweetClient{ctx: ctx}
}

func (tc *tweetClient) GetTweetsFromStream(responses chan<- TweetResponse) error {
	req, err := tc.getRequest()
	if err != nil {
		fmt.Println("pID #", tc.ctx.Value("processID"), " - Request error=", err)
		return err
	}

	reader, err := tc.getResponse(req)
	if err != nil {
		fmt.Println("pID #", tc.ctx.Value("processID"), " - Response error=", err)
		return err
	}

	return tc.processTweet(reader, responses)
}

func (tc *tweetClient) getRequest() (*http.Request, error) {
	req, err := http.NewRequest("GET", ApiUrl, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Add("Authorization", "Bearer "+BearerToken)

	return req, nil
}

func (tc *tweetClient) getResponse(req *http.Request) (*bufio.Reader, error) {
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	} else if resp.StatusCode >= 400 {
		return nil, errors.New(resp.Status)
	}

	return bufio.NewReader(resp.Body), nil
}

func (tc *tweetClient) processTweet(reader *bufio.Reader, responses chan<- TweetResponse) error {
	numTweetsRetrieved := 0
	for {
		line, err := reader.ReadBytes('\n')
		if err != nil {
			fmt.Println("pID #", tc.ctx.Value("processID"), " - Reader error=", err)
			return err
		}

		tc.processLine(&numTweetsRetrieved, responses, line)

		select {
		case _ = <-tc.ctx.Done():
			fmt.Println("pID #", tc.ctx.Value("processID"), " - Context done")
			return nil
		default:
			continue
		}
	}
}

func (tc *tweetClient) processLine(numTweetsRetrieved *int, responses chan<- TweetResponse, line []byte) {
	(*numTweetsRetrieved)++
	fmt.Println("pID #", tc.ctx.Value("processID"), " - Retrieving tweet", *numTweetsRetrieved)
	responses <- NewTweetResponse(string(line))

	pauseDuration := rand.Intn(MaxPauseDuration-MinPauseDuration) + MinPauseDuration
	time.Sleep(time.Duration(pauseDuration) * time.Millisecond)
}
