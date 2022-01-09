package timetree

import (
	"net/http"
)

type Client struct {
	client  *http.Client
	BaseURL string
	token   string
}

func NewClient(httpClient *http.Client, token string) *Client {
	if httpClient == nil {
		httpClient = &http.Client{}
	}
	baseURL := TIMETREE_BASE_URL
	return &Client{
		client:  httpClient,
		BaseURL: baseURL,
		token:   token,
	}

}
