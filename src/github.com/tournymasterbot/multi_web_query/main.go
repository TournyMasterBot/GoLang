package main

import (
	"fmt"

	webrequesthelper "github.com/tournymasterbot/web_request_helpers"
)

func main() {
	// Process a list of urls using the specified callback
	processURLs(
		urlCallback,
		"https://api.github.com/users/tournymasterbot",
		"https://api.github.com/users/JayyJayy",
		"https://api.github.com/users/EpicGames",
	)
}

// Variadic function demo
func processURLs(callback func(string), urls ...string) {
	for _, url := range urls {
		fmt.Println("*****")
		fmt.Println("Fetching: " + url)
		callback(url)
	}
}

// Process url result demo
func urlCallback(url string) {
	getrequest := webrequesthelper.GetWebRequestConfig{
		URL:     url,
		Headers: map[string]string{},
	}
	result, err := webrequesthelper.Get(getrequest)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Data:", result)
}
