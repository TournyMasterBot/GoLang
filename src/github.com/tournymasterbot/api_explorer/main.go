package main

import (
	"fmt"

	webrequesthelper "github.com/tournymasterbot/web_request_helpers"
)

func main() {
	// GET Examples
	fmt.Println("****** GET ******")

	// Specify the request values
	getrequest := webrequesthelper.GetWebRequestConfig{
		URL:     "https://api.github.com/users/tournymasterbot",
		Headers: map[string]string{},
	}

	// Get the web request-response
	getresponse, err := webrequesthelper.Get(getrequest)
	if err != nil {
		fmt.Println(err)
		return
	} else {
		// Output the request-response
		fmt.Println(getresponse)
	}

	// POST Examples
	fmt.Println("****** POST ******")

	postrequest := webrequesthelper.PostWebRequestConfig{
		URL:     "https://httpbin.org/post",
		Headers: map[string]string{},
		PostData: map[string]string{
			"username": "Goher Go",
			"email":    "go@gmail.com",
		},
	}

	postresponse, err := webrequesthelper.Post(postrequest)
	if err != nil {
		fmt.Println(err)
		return
	} else {
		fmt.Println(postresponse)
	}
}
