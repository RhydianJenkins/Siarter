package httpClient

import (
	"encoding/json"
	"net/http"
	"time"
)

var client = &http.Client{Timeout: 10 * time.Second}

type Client struct {
	url  string
	mock bool
}

type Boat struct {
	LAT       string
	LON       string
	COURSE    string
	SPEED     string
	STATUS    string
	TIMESTAMP string
	MMSI      string
}

func NewClient(url string, mock bool) *Client {
	return &Client{url, mock}
}

func getMockResponse() []*Boat {
	return []*Boat{
		{
			"50.388360",
			"-5.319272",
			"226",
			"68",
			"99",
			"2022-06-03T10:51:59",
			"232007968",
		},
	}
}

func (c *Client) Get() ([]*Boat, error) {
	if c.mock {
		return getMockResponse(), nil
	}

	r, clientError := client.Get(c.url)
	if clientError != nil {
		return nil, clientError
	}

	defer r.Body.Close()

	response := []*Boat{}
	decodeError := json.NewDecoder(r.Body).Decode(&response)
	if decodeError != nil {
		return nil, decodeError
	}

	return response, nil
}
