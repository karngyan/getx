package clients

import (
	"errors"
	"io/ioutil"
	"log"
	"net/http"
)

type NetworkClient struct{}

// Returns an address for a new NetworkClient
func NewNetworkClient() *NetworkClient {
	return &NetworkClient{}
}

// Gets html byte slice from Uri
func (nc *NetworkClient) GetHtmlBytes(uri string) ([]byte, error) {
	defer catchPanic("GetHtml()")
	// initialize a new http client
	client := &http.Client{}
	req, err := http.NewRequest("GET", uri, nil)
	if err != nil {
		return nil, errors.New("couldn't create New GET Request " + uri)
	}

	// perform request
	resp, err := client.Do(req)
	if err != nil {
		return nil, errors.New("couldn't perform GET request to " + uri)
	}

	defer resp.Body.Close()
	bytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, errors.New("unable to read the response body")
	}

	return bytes, nil
}

// Catch panics when they occur
func catchPanic(fnName string) {
	if r := recover(); r != nil {
		log.Println("Error occurred in", fnName, ":", r)
	}
}
