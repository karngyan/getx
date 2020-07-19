package clients

import (
	"errors"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

const (
	durationBetweenRetries = 2 * time.Second
)

type NetworkClient struct{}

// Returns an address for a new NetworkClient
func NewNetworkClient() *NetworkClient {
	return &NetworkClient{}
}

// Gets html byte slice from Uri
func (nc *NetworkClient) GetHtmlBytes(uri string, retryLimit int) ([]byte, error) {
	defer catchPanic("GetHtml()")

	// init retryLeft
	retryLeft := retryLimit

	// initialize a new http client
	client := &http.Client{}
	req, err := http.NewRequest("GET", uri, nil)
	if err != nil {
		return nil,
			errors.New("couldn't create New GET Request " + uri)
	}

	ticker := time.NewTicker(durationBetweenRetries)
	responseReceivedChannel := make(chan bool)
	retriesDoneChannel := make(chan bool)

	var response *http.Response
	log.Println("Initial GET: ", uri)
	performRequest(client,
		req, retriesDoneChannel, &retryLeft, &response)

	if response == nil {
		go func() {
			for {
				select {
				case <-retriesDoneChannel:
					responseReceivedChannel <- true
					return
				case t := <-ticker.C:
					log.Println("GET:", uri, t)
					// perform request
					go performRequest(client,
						req, retriesDoneChannel, &retryLeft, &response)
				}
			}
		}()
	} else {
		responseReceivedChannel <- true
	}

	// wait till response is received
	<-responseReceivedChannel
	ticker.Stop()
	log.Println("Ticker Stopped for: ", uri)

	var htmlBytes []byte

	if response != nil {
		defer response.Body.Close()
		htmlBytes, err = ioutil.ReadAll(response.Body)
		if err != nil {
			log.Println(err)
			return htmlBytes, errors.New("unable to read the response body")
		}
	}

	return htmlBytes, nil
}

// Perform the request and
// send true to done channel based on the response or retryLimit
func performRequest(client *http.Client,
	request *http.Request,
	retriesDoneChannel chan bool,
	retryLeft *int,
	response **http.Response) {
	defer catchPanic("performRequest()")

	resp, err := client.Do(request)
	if err == nil || *retryLeft <= 0 {
		retriesDoneChannel <- true
	}

	*retryLeft -= 1
	*response = resp

}

// Catch panics when they occur
func catchPanic(fnName string) {
	if r := recover(); r != nil {
		log.Println("Error occurred in", fnName, ":", r)
	}
}
