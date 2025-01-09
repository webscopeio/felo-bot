package slack

import (
	"net/http"
)


func (s *Client) GetChannelList() (*http.Response, error) {
	resp, err := s.request(Request{
		path:   "/conversations.list",
		method: http.MethodGet,
	})
	if err != nil {
		return nil, err
	}
	return resp, nil
}

