package server

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/gin-gonic/gin"
	"webscope.io/felo/modules/slack"
	"webscope.io/felo/modules/supabase"
)

func eventsHandler(ctx *gin.Context, client *slack.Client, db *supabase.DB) {
	req := ctx.Request
	defer req.Body.Close()
	body, err := io.ReadAll(req.Body)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, Response{Status: "error", Ok: false, Message: "Error reading request body"})
		return
	}

	var eventBody struct {
		Type      string `json:"type"`
		Challenge string `json:"challenge"`
	}

	if err := json.Unmarshal(body, &eventBody); err != nil {
		ctx.JSON(http.StatusBadRequest, Response{Status: "error", Ok: false, Message: "Error parsing request body"})
		return
	}

	eventType := eventBody.Type
	switch eventType {
	case "url_verification":
		challenge := eventBody.Challenge
		ctx.JSON(http.StatusOK, gin.H{"challenge": challenge})
	default:
		ctx.JSON(http.StatusOK, gin.H{"status": "success", "ok": true})
	}
}
