package main

import (
	"errors"
	"fmt"
	"net/http"
)

type hitResult struct {
	url    string
	status string
}

var errRequestFailed = errors.New("request failed")

func main() {
	c := make(chan hitResult)
	urls := []string{
		"https://www.airbnb.com/",
		"https://www.google.com/",
		"https://www.amazon.com/",
		"https://www.reddit.com/",
		"https://soundcloud.com/",
		"https://www.facebook.com/",
		"https://www.instagram.com/",
	}

	for _, url := range urls {
		go hitUrl(url, c)
	}
}

func hitUrl(url string, c chan<- hitResult) {
	fmt.Println("Checking:", url)
	resp, err := http.Get(url)
	status := "OK"
	if err != nil || resp.StatusCode >= 400 {
		status = "FAILED"
	}
	c <- hitResult{url: url, status: status}
}
