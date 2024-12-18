package server

import (
	"net/http"
	"os"
	"strings"

	"webscope.io/felo/modules/slack"
	"github.com/gin-gonic/gin"
)

type Response struct {
	Status string `json:"status"`
	Ok bool `json:"ok"`
	Data interface{} `json:"data,omitempty"`
}

// Callback creator that enriches standard gin route handler with the slack client
func createHandlerWithClient(client *slack.Client) func(func(ctx *gin.Context, client *slack.Client)) gin.HandlerFunc {
	return func(fn func(ctx *gin.Context, client *slack.Client)) gin.HandlerFunc {
		return func(ctx *gin.Context) {
			fn(ctx, client)
		}
	}
}

func matchHandler(ctx *gin.Context, client *slack.Client) {
	resp := Response{ Status: "success", Ok: true, Data: "Hello from Felo go app. Received /match command!" }
	client.PostMessage("C0859BH2E2W", "Hello from Felo go app. Received /match command!")
	ctx.JSON(http.StatusOK, resp)
}

func New(host string, env string, client *slack.Client) {
	if env == "production" {
		gin.SetMode(gin.ReleaseMode)
	}
	router := gin.New()

	handler := createHandlerWithClient(client)

	router.GET("/slack/match", handler(matchHandler))

	serverUrl := host
	if strings.Contains(host, "localhost") && !strings.Contains(host, ":") {
		port := os.Getenv("PORT")
		if port == "" {
			port = "8080"
		}
		serverUrl = host + ":" + port
	}
	if err := router.Run(serverUrl); err != nil {
		panic(err)
	}
}
