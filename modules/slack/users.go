package slack

import (
	"net/http"

	"webscope.io/felo/modules/utils"
)


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

func (s* Client) GetSlackUsers() ([]SlackUser, error) {
	resp, err := s.request(Request{
		path:  "/users.list",
		method: http.MethodGet,
		params: map[string]string{"limit": "200"},
	})
	if err != nil {
		return nil, err
	}
	var data UserListResponse
	if err := utils.MapBodyTo(resp, &data); err != nil {
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
		if err := utils.MapBodyTo(resp, &nextData); err != nil {
			return nil, err
		}
		
		data.Members = append(data.Members, nextData.Members...)
		cursor = nextData.Response_Metadata.NextCursor
	}
	return data.Members, nil
}
