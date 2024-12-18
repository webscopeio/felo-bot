package slack

import (
	"io"
	"net/http"
	"net/url"
	"strings"
)

const API_URL = "https://slack.com/api"

type Client struct {
	token  string
	url    string
	client *http.Client
}

func New(token string) *Client {
	client := &http.Client{}
	return &Client{
		token:  token,
		url:    API_URL,
		client: client,
	}
}

type Request struct {
	path   string
	method string

	// optional parameters
	params map[string]string
	body   io.Reader
}

func (c *Client) request(r Request) (*http.Response, error) {
	url := c.url + r.path
	for k, v := range r.params {
		if strings.Contains(url, "?") {
			url += "&" + k + "=" + v
		} else {
			url += "?" + k + "=" + v
		}
	}
	req, err := http.NewRequest(r.method, url, r.body)
	if err != nil {
		return nil, err
	}
	if r.method == http.MethodGet {
		req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	} else {
		req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
		req.Header.Add("Content-Type", "application/json")
	}
	req.Header.Set("Authorization", "Bearer "+c.token)
	return c.client.Do(req)
}

func (c *Client) PostMessage(channel string, text string) (*http.Response, error) {
	body := url.Values{}
	body.Set("channel", channel)
	body.Set("text", text)
	return c.request(Request{
		path:   "/chat.postMessage",
		method: http.MethodPost,
		body:   strings.NewReader(body.Encode()),
	})
}

func (c *Client) GetChannelList() (*http.Response, error) {
	return c.request(Request{
		path:   "/conversations.list",
		method: http.MethodGet,
	})
}
