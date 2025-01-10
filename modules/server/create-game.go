package server

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"webscope.io/felo/modules/slack"
	"webscope.io/felo/modules/supabase"
)

func createGameHandler(ctx *gin.Context, slack *slack.Client, db *supabase.DB) {
	triggerId := ctx.Request.PostFormValue("trigger_id")
	text := ctx.Request.PostFormValue("text")

	resp, err := slack.CreateGame(triggerId, text)

	fmt.Printf("CreateGame response: %v", resp)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, Response{Status: "error", Ok: false, Message: "Error creating game: " + err.Error()})
		return
	}
	
	ctx.JSON(http.StatusOK, Response{Status: "success", Ok: true, Data: resp})
}
