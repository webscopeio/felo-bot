package slack

import (
	"bytes"
	"fmt"
	"net/http"
)

func (s* Client) CreateGame(trigger_id string, text string) (*http.Response, error) {
	body := []byte(fmt.Sprintf(`{
		"trigger_id": "%s",
		"view": {
				"type": "modal",
				"title": {
						"type": "plain_text",
						"text": "Create Game"
				},
				"submit": {
						"type": "plain_text",
						"text": "Submit"
				},
				"close": {
						"type": "plain_text",
						"text": "Cancel"
				},
				"blocks": [
						{
								"type": "input",
								"block_id": "game_name_block",
								"label": {
										"type": "plain_text",
										"text": "Game Name"
								},
								"element": {
										"type": "plain_text_input",
										"action_id": "game_name",
										"initial_value": "%s",
										"placeholder": {
												"type": "plain_text",
												"text": "Enter game name"
										},
										"multiline": false
								}
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