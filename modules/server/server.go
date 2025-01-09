package server

import (
	"github.com/gin-gonic/gin"
	"webscope.io/felo/modules/slack"
	"webscope.io/felo/modules/supabase"
)

/* ---------------------------------- Setup --------------------------------- */

type Server struct {}

type Response struct {
	Status string      `json:"status"`
	Ok     bool        `json:"ok"`
	Message string      `json:"message,omitempty"`
	Data   interface{} `json:"data,omitempty"`
}

type HandlerCallback = func(func(ctx *gin.Context, client *slack.Client, db *supabase.DB)) gin.HandlerFunc

// Callback creator that enriches standard gin route handler with the slack and supabase clients
func createHandlerWithClients(client *slack.Client, db *supabase.DB) HandlerCallback {
	return func(callbackHandlerDef func(ctx *gin.Context, client *slack.Client, db *supabase.DB)) gin.HandlerFunc {
		return func(ctx *gin.Context) {
			callbackHandlerDef(ctx, client, db)
		}
	}
}

func (s* Server) New(env string, port string, client *slack.Client, db *supabase.DB) {
	if env == "production" {
		gin.SetMode(gin.ReleaseMode)
	}
	router := gin.New()
	handler := createHandlerWithClients(client, db)

	router.POST("/slack/create-game", handler(createGameHandler))
	router.POST("/slack/events", handler(eventsHandler))

	if port == "" {
		port = "8080"
	}

	if err := router.Run(":" + port); err != nil {
		panic(err)
	}
}
