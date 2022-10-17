package app

import (
	"regexp"
	"strings"
)

type TweetResponse struct {
	Response string
	Hashtags map[string]int
}

func NewTweetResponse(rawMessage string) TweetResponse {
	hashtags := getHashtags(rawMessage)
	return TweetResponse{Response: rawMessage, Hashtags: hashtags}
}

func getHashtags(rawMessage string) map[string]int {
	hashtags := make(map[string]int)

	if !strings.Contains(rawMessage, "#") {
		return hashtags
	}

	regex := regexp.MustCompile(`#\w+\b`)
	matches := regex.FindAllString(rawMessage, -1)

	for _, v := range matches {
		if _, ok := hashtags[v]; !ok {
			hashtags[v] = 1
		} else {
			curr := hashtags[v]
			curr++
			hashtags[v] = curr
		}
	}

	return hashtags
}
