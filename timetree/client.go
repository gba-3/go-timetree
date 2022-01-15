package timetree

import (
	"net/http"
	"sync"
)

type service struct {
	cli *Client
}

type Client struct {
	clientMu  sync.Mutex
	client    *http.Client
	BaseURL   string
	token     string
	sv        service
	Calendars *CalendarService
}

func NewClient(httpClient *http.Client, token string) *Client {
	if httpClient == nil {
		httpClient = &http.Client{}
	}
	baseURL := TIMETREE_BASE_URL
	c := &Client{
		client:  httpClient,
		BaseURL: baseURL,
		token:   token,
	}
	c.sv.cli = c
	c.Calendars = (*CalendarService)(&c.sv)
	return c
}

func (c *Client) Client() *http.Client {
	c.clientMu.Lock()
	defer c.clientMu.Unlock()
	return c.client
}
