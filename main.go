package main

import (
	"context"
	"fmt"
	"sync"

	"example.com/goroutines/app"
)

const (
	MaxTweets    = 100
	NumOfClients = 20
)

func main() {
	RunConcurrencyTest()
}

// Concurrency example with the twitter streaming API, using goroutines, channels, waitGroups, and context
func RunConcurrencyTest() {
	ctx, cancel := context.WithCancel(context.Background())
	responses := make(chan app.TweetResponse)
	var wg sync.WaitGroup
	wg.Add(1)

	go func() {
		initiateProcessing(ctx, &wg, responses)
		processResponses(responses, cancel)
	}()

	wg.Wait()
	fmt.Println("Processing completed")
}

func initiateProcessing(ctx context.Context, wg *sync.WaitGroup, responses chan app.TweetResponse) {
	skipPause := make(chan interface{}, NumOfClients)
	for i := 1; i <= NumOfClients; i++ {
		pid := i
		wg.Add(1)
		go func() {
			defer wg.Done()
			app.NewTweetClient(context.WithValue(ctx, "processID", pid)).GetTweetsFromStream(responses, skipPause)
		}()
	}
	wg.Done()
}

func processResponses(responses chan app.TweetResponse, cancel context.CancelFunc) {
	receivedTweetNumber := 0
	hashtags := make(map[string]int)

	for resp := range responses {
		receivedTweetNumber++
		for k := range resp.Hashtags {
			hashtags[k] += 1
		}

		if receivedTweetNumber%10 == 0 {
			fmt.Println("CURRENT PROGRESS: tweets received =", receivedTweetNumber, ", hashtags =", hashtags)
		}

		if receivedTweetNumber >= MaxTweets {
			fmt.Println("Maximum number of tweets received.  Ending processing.")
			fmt.Println("FINAL PROGRESS: tweets received =", receivedTweetNumber, ", hashtags =", hashtags)
			cancel()
			break
		}
	}
}
