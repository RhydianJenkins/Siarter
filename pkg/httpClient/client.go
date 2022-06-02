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

type ApiResponse []*Boat

func NewClient(url string, mock bool) *Client {
	return &Client{url, mock}
}

func getMockResponse() *ApiResponse {
	return &ApiResponse{
		&Boat{
			"123",
			"123",
			"123",
			"123",
			"123",
			"123",
			"123",
		},
	}
}

func (c *Client) Get() (*ApiResponse, error) {
	if c.mock {
		return getMockResponse(), nil
	}

	r, clientError := client.Get(c.url)
	if clientError != nil {
		return nil, clientError
	}

	defer r.Body.Close()

	response := &ApiResponse{}
	decodeError := json.NewDecoder(r.Body).Decode(response)
	if decodeError != nil {
		return nil, decodeError
	}

	return response, nil
}
