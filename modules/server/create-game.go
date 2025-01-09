package server

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"webscope.io/felo/modules/slack"
	"webscope.io/felo/modules/supabase"
	"webscope.io/felo/modules/utils"
)

func createGameHandler(ctx *gin.Context, slack *slack.Client, db *supabase.DB) {
	triggerId := ctx.Request.PostFormValue("trigger_id")
	text := ctx.Request.PostFormValue("text")

	resp, err := slack.CreateGame(triggerId, text)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, Response{Status: "error", Ok: false, Message: "Error creating game: " + err.Error()})
		return
	}
	var responseBody struct {
		View struct {
			State struct {
				Values map[string]map[string]struct {
					Value string `json:"value"`
				} `json:"values"`
			} `json:"state"`
		} `json:"view"`
	}

	err = utils.MapBodyTo(resp, &responseBody)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, Response{Status: "error", Ok: false, Message: "Error reading response body: " + err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, Response{Status: "success", Ok: true, Data: resp})
}
