package slack

import (
	"net/http"
)

func (s *Client) CreateGame(trigger_id string, text string) (*http.Response, error) {
	body := CreateView(trigger_id, View{
		Type:   "modal",
		Title:  PlainText("Create Game"),
		Submit: PlainText("Submit"),
		Close:  PlainText("Cancel"),
		Blocks: []Block{
			{
				Type:    "input",
				BlockId: "game_name_block",
				Label:   PlainText("Game Name"),
				Element: Element{
					Type:         "plain_text_input",
					ActionId:     "game_name",
					InitialValue: text,
					Placeholder:  PlainText("Enter game name"),
					Multiline:    false,
				},
			},
		},
	})
	return s.request(Request{
		path:   "/views.open",
		method: http.MethodPost,
		body:   body,
	})
}
