package slack

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"net/url"
)

type Request struct {
	path   string
	method string

	// optional parameters
	params map[string]string
	body   io.Reader
}

func (s *Client) request(r Request) (*http.Response, error) {
	uri, err := url.Parse(s.API_URL + r.path)
	if err != nil {
		return nil, err
	}
	query := uri.Query()
	for k, v := range r.params {
		query.Set(k, v)
	}
	if r.body != nil {
		// Read and print the body for debugging
		bodyBytes, _ := io.ReadAll(r.body)
		fmt.Printf("Request body: %s\n", string(bodyBytes))
		// Reset the body for the actual request
		r.body = bytes.NewBuffer(bodyBytes)
	}
	uri.RawQuery = query.Encode()
	req, err := http.NewRequest(r.method, uri.String(), r.body)
	if err != nil {
		return nil, err
	}
	if r.method == http.MethodGet {
		req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	} else {
		req.Header.Set("Content-Type", "application/json; charset=utf-8")
	}
	req.Header.Set("Authorization", "Bearer "+s.BOT_TOKEN)
	return s.HTTP_CLIENT.Do(req)
}
