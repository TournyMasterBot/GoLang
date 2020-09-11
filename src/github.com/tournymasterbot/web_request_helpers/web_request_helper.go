package webrequesthelper

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

// GetWebRequestConfig : config
type GetWebRequestConfig struct {
	URL     string
	Headers map[string]string
}

// PostWebRequestConfig : config
type PostWebRequestConfig struct {
	URL      string
	Headers  map[string]string
	PostData map[string]string
}

// Get : Returns a string result from the requested web request
func Get(request GetWebRequestConfig) (string, error) {
	response, err := GetBytes(request)
	if err != nil {
		fmt.Println(err)
		return "", err
	}
	if response != nil || len(response) > 0 {
		return string(response), nil
	}
	return "", nil
}

// GetBytes : Returns a byte array from the requested web request
func GetBytes(request GetWebRequestConfig) ([]byte, error) {
	// Issue Get Request
	response, err := http.Get(request.URL)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	// Get request response
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}

// Post : Returns a string response from the requested web request
func Post(request PostWebRequestConfig) (string, error) {
	response, err := PostBytes(request)
	if err != nil {
		return "", err
	}
	return string(response), err
}

// PostBytes : Returns a byte array from the requested web request
func PostBytes(request PostWebRequestConfig) ([]byte, error) {
	// Serialize the post data
	reqBody, err := json.Marshal(request.PostData)
	if err != nil {
		return nil, err
	}

	// Issue post request
	resp, err := http.Post(request.URL, "application/json", bytes.NewBuffer(reqBody))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	// Read post response
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	return body, nil
}
