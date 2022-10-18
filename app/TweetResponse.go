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

	for _, v := range regexp.MustCompile(`#\w+\b`).FindAllString(rawMessage, -1) {
		hashtags[v] += 1
	}
	return hashtags
}
