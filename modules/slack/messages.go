package slack

import (
	"net/http"
	"net/url"
	"strings"
)

func (s *Client) PostMessage(channel string, text string) (*http.Response, error) {
	body := url.Values{}
	body.Set("channel", channel)
	body.Set("text", text)
	return s.request(Request{
		path:   "/chat.postMessage",
		method: http.MethodPost,
		body:   strings.NewReader(body.Encode()),
	})
}
