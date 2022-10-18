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
	GetTweetsFromStream(responses chan<- TweetResponse, skipPause chan interface{}) error
}

const (
	// enter your bearertoken for the twitter v2 API here
	BearerToken        = "abc"
	ApiUrl             = "https://api.twitter.com/2/tweets/sample/stream"
	MinPauseDuration   = 100
	MaxPauseDuration   = 10000
	SkipPauseThreshold = 0.9
)

type tweetClient struct {
	ctx context.Context
}

func NewTweetClient(ctx context.Context) TweetClienter {
	return &tweetClient{ctx: ctx}
}

func (tc *tweetClient) GetTweetsFromStream(responses chan<- TweetResponse, skipPause chan interface{}) error {
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

	return tc.processTweet(reader, responses, skipPause)
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

func (tc *tweetClient) processTweet(reader *bufio.Reader, responses chan<- TweetResponse, skipPause chan interface{}) error {
	numTweetsRetrieved := 0
	for {
		line, err := reader.ReadBytes('\n')
		if err != nil {
			fmt.Println("pID #", tc.ctx.Value("processID"), " - Reader error=", err)
			return err
		}

		tc.processLine(&numTweetsRetrieved, responses, skipPause, line)

		select {
		case _ = <-tc.ctx.Done():
			fmt.Println("pID #", tc.ctx.Value("processID"), " - Context done")
			return nil
		default:
			continue
		}
	}
}

func (tc *tweetClient) processLine(numTweetsRetrieved *int, responses chan<- TweetResponse, skipPause chan interface{}, line []byte) {
	(*numTweetsRetrieved)++
	fmt.Println("pID #", tc.ctx.Value("processID"), " - Retrieving tweet", *numTweetsRetrieved)
	responses <- NewTweetResponse(string(line))

	tc.pauseProcessing(skipPause)
}

func (tc *tweetClient) pauseProcessing(skipPause chan interface{}) {
	select {
	case _ = <-skipPause:
		fmt.Println("pID #", tc.ctx.Value("processID"), " - SkipPause message read from the queue.")
		return
	default:
	}

	pauseDuration := rand.Intn(MaxPauseDuration-MinPauseDuration) + MinPauseDuration
	if pauseDuration >= (SkipPauseThreshold * MaxPauseDuration) {
		fmt.Println("pID #", tc.ctx.Value("processID"), " - Pause threshold surpassed.  Pushing skipPause message.")
		skipPause <- true
	}

	time.Sleep(time.Duration(pauseDuration) * time.Millisecond)
}
