package slack

import (
	"bytes"
	"encoding/json"
)

type View struct {
	Type   string    `json:"type"`
	Title  plainText `json:"title"`
	Submit plainText `json:"submit"`
	Close  plainText `json:"close"`
	Blocks []Block   `json:"blocks"`
}

type Block struct {
	Type    string    `json:"type"`
	BlockId string    `json:"block_id"`
	Label   plainText `json:"label"`
	Element Element   `json:"element"`
}

type plainText struct {
	Type string `json:"type"`
	Text string `json:"text"`
}

type Element struct {
	Type         string `json:"type"`
	ActionId     string `json:"action_id"`
	InitialValue string `json:"initial_value"`
	Placeholder  struct {
		Type string `json:"type"`
		Text string `json:"text"`
	}
	Multiline bool `json:"multiline"`
}

type viewPayload struct {
	TriggerID string `json:"trigger_id"`
	View      View   `json:"view"`
}

func CreateView(trigger_id string, view View) *bytes.Buffer {
	payload := viewPayload{
		TriggerID: trigger_id,
		View:      view,
	}

	jsonBytes, err := json.Marshal(payload)
	if err != nil {
		return bytes.NewBuffer([]byte{})
	}
	return bytes.NewBuffer(jsonBytes)
}

func PlainText(text string) plainText {
	return plainText{Type: "plain_text", Text: text}
}
