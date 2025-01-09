package modules

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
)

type Slack struct {
	BOT_TOKEN  string
	API_URL string
	HTTP_CLIENT *http.Client
}


type Request struct {
	path   string
	method string

	// optional parameters
	params map[string]string
	body   io.Reader
}

func (s *Slack) request(r Request) (*http.Response, error) {
	uri, err := url.Parse(s.API_URL + r.path)
	if err != nil {
		return nil, err
	}
	query := uri.Query()
	for k, v := range r.params {
		query.Set(k, v)
	}
	uri.RawQuery = query.Encode()
	req, err := http.NewRequest(r.method, uri.String(), r.body)
	if err != nil {
		return nil, err
	}
	if r.method == http.MethodGet {
		req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	} else {
		req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
		req.Header.Add("Content-Type", "application/json")
	}
	req.Header.Set("Authorization", "Bearer "+s.BOT_TOKEN)
	return s.HTTP_CLIENT.Do(req)
}

func mapBodyTo[T interface{}](resp *http.Response, to *T) error {
	if (resp.StatusCode != http.StatusOK) {
		return fmt.Errorf("error fetching slack users: %s", resp.Status)
	}
	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	return json.Unmarshal(data, to)
}

func (s *Slack) PostMessage(channel string, text string) (*http.Response, error) {
	body := url.Values{}
	body.Set("channel", channel)
	body.Set("text", text)
	return s.request(Request{
		path:   "/chat.postMessage",
		method: http.MethodPost,
		body:   strings.NewReader(body.Encode()),
	})
}

func (s *Slack) GetChannelList() (*http.Response, error) {
	resp, err := s.request(Request{
		path:   "/conversations.list",
		method: http.MethodGet,
	})
	if err != nil {
		return nil, err
	}
	return resp, nil
}

type SlackUser struct {
	Id string `json:"id"`
	Name string `json:"name"`
	Profile struct {
		RealName string `json:"real_name"`
	} `json:"profile"`
}

type UserListResponse struct {
	Ok bool `json:"ok"`
	Members []SlackUser `json:"members"`
	Response_Metadata struct {
		NextCursor string `json:"next_cursor"`
	} `json:"response_metadata"`
}

func (s* Slack) GetSlackUsers() ([]SlackUser, error) {
	resp, err := s.request(Request{
		path:  "/users.list",
		method: http.MethodGet,
		params: map[string]string{"limit": "200"},
	})
	if err != nil {
		return nil, err
	}
	var data UserListResponse
	if err := mapBodyTo(resp, &data); err != nil {
		return nil, err
	}
	cursor := data.Response_Metadata.NextCursor
	for cursor != "" {
		resp, err := s.request(Request{
			path:   "/users.list",
			method: http.MethodGet,
			params: map[string]string{
				"limit":  "200",
				"cursor": cursor,
			},
		})
		if err != nil {
			return nil, err
		}
		
		var nextData UserListResponse
		if err := mapBodyTo(resp, &nextData); err != nil {
			return nil, err
		}
		
		data.Members = append(data.Members, nextData.Members...)
		cursor = nextData.Response_Metadata.NextCursor
	}
	return data.Members, nil
}

func (s* Slack) PostMatch() (*http.Response, error) {
		body := []byte(`{
			"type": "modal",
			"title": {
				"type": "plain_text",
				"text": "Submit a match result",
			},
			"submit": {
				"type": "plain_text",
				"text": "Submit",
			},
			close: {
				"type": "plain_text",
				"text": "Cancel",
			},
			"blocks": [
				{
					"type": "header",
					"text": {
						"type": "plain_text",
						"text": "Home Team",
					}
				},
				{
					"type": "section",

				}
			]
		}`)
		return s.request(Request{
			path: "/views.open",
			method: http.MethodPost,
			body: bytes.NewBuffer(body),
		})
}

func (s* Slack) CreateGame(trigger_id string, text string) (*http.Response, error) {
	body := []byte(fmt.Sprintf(`{
	"trigger_id": "%s",
	"view": {
			"type": "modal",
			"title": {
				"type": "plain_text",
				"text": "Create a new game type"
			},
			"submit": {
				"type": "plain_text",
				"text": "Submit"
			},
			"close": {
				"type": "plain_text",
				"text": "Cancel"
			},
			"submit_disabled": true,
			"blocks": [
				{
					"type": "input",
					"initial_value": "%s",
					"label": {
						"type": "plain_text",
						"text": "Game Name"
					},
					"element": {
						"type": "plain_text_input",
						"action_id": "game_name",
						"placeholder": {
							"type": "plain_text",
							"text": "Enter game name"
						},
						"multiline": false
					},
					"optional": false
				}
			],
			"callback_id": "create_game"
		}
	}`, trigger_id, text))
		return s.request(Request{
			path: "/views.open",
			method: http.MethodPost,
			body: bytes.NewBuffer(body),
		})
}