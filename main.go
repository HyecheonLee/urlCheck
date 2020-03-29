package main

import (
	"fmt"
	"net/http"
)

type requestResult struct {
	url    string
	status string
	err    error
}

func main() {
	c := make(chan requestResult)
	urls := []string{
		"https://www.airbnb.com/",
		"https://www.google.com/",
		"https://www.amazon.com/",
		"https://www.reddit.com/",
		"https://www.google.com/",
		"https://soundcloud.com/",
		"https://www.facebook.com/",
		"https://www.instagram.com/",
		"https://academy.nomadcoders.co/",
	}
	for _, url := range urls {
		go hitUrl(url, c)
	}
	for i := 0; i < len(urls); i++ {
		fmt.Println(<-c)
	}
}
func hitUrl(url string, c chan<- requestResult) {
	fmt.Println("Checking:", url)
	resp, err := http.Get(url)
	if err != nil {
		c <- requestResult{err: err}
	} else {
		c <- requestResult{url: url, status: resp.Status}
	}
}
