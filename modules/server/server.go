package server

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"webscope.io/felo/modules/slack"
)

type Response struct {
	Status string      `json:"status"`
	Ok     bool        `json:"ok"`
	Data   interface{} `json:"data,omitempty"`
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
	resp := Response{Status: "success", Ok: true, Data: "Hello from Felo go app. Received /match command!"}
	client.PostMessage("C0859BH2E2W", "Hello from Felo go app. Received /match command!")
	ctx.JSON(http.StatusOK, resp)
}

func eventsHandler(ctx *gin.Context, client *slack.Client) {
	evenType := ctx.PostForm("type")
	switch evenType {
		case "url_verification":
			challenge := ctx.PostForm("challenge")
			ctx.JSON(http.StatusOK, gin.H{"challenge": challenge})
		default:
			ctx.JSON(http.StatusOK, gin.H{"status": "success", "ok": true})
		}
}

func New(env string, port string, client *slack.Client) {
	if env == "production" {
		gin.SetMode(gin.ReleaseMode)
	}
	router := gin.New()

	handler := createHandlerWithClient(client)

	router.GET("/slack/match", handler(matchHandler))
	router.POST("/slack/events", handler(eventsHandler))

	if port == "" {
		port = "8080"
	}

	if err := router.Run(":" + port); err != nil {
		panic(err)
	}
}
