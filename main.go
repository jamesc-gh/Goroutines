package main

import (
	"context"
	"fmt"
	"sync"

	"example.com/goroutines/app"
)

const (
	MaxTweets = 100
)

func main() {
	RunConcurrencyTest()
}

// Concurrency example with the twitter streaming API, using goroutines, channels, waitGroups, and context
func RunConcurrencyTest() {
	ctx, cancel := context.WithCancel(context.Background())
	responses := make(chan app.TweetResponse)
	var wg sync.WaitGroup

	for i := 1; i <= 20; i++ {
		initiateProcessing(ctx, &wg, i, responses)
	}

	processResponses(responses, cancel)

	wg.Wait()
	fmt.Println("Processing completed")
}

func processResponses(responses chan app.TweetResponse, cancel context.CancelFunc) {
	receivedTweetNumber := 0
	hashtags := make(map[string]int)

	for resp := range responses {
		receivedTweetNumber++
		hashtags = appendHashtags(hashtags, resp.Hashtags)

		if receivedTweetNumber%10 == 0 {
			fmt.Println("CURRENT PROGRESS: tweets received =", receivedTweetNumber, ", hashtags =", hashtags)
		}

		if receivedTweetNumber >= MaxTweets {
			fmt.Println("Maximum number of tweets received.  Ending processing.")
			cancel()
			break
		}
	}

	fmt.Println("FINAL PROGRESS: tweets received =", receivedTweetNumber, ", hashtags =", hashtags)
}

func initiateProcessing(ctx context.Context, wg *sync.WaitGroup, i int, responses chan app.TweetResponse) {
	wg.Add(1)
	go func() {
		defer wg.Done()
		app.NewTweetClient(context.WithValue(ctx, "processID", i)).GetTweetsFromStream(responses)
	}()
}

func appendHashtags(current, new map[string]int) map[string]int {
	for k := range new {
		if _, ok := current[k]; !ok {
			current[k] = 1
		} else {
			curr := current[k]
			curr++
			current[k] = curr
		}
	}

	return current
}
