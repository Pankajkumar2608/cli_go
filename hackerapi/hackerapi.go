package hackerapi

import (
	"net/http"
	"time"
)

type Client struct {
	httpClient http.Client
}

const baseURL = "http://hn.algolia.com/api/v1"

func NewClient() Client {
	return Client {
		httpClient: http.Client{
			Timeout: time.Minute,
		},
	}
}


