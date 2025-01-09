package modules

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"

	"github.com/gin-gonic/gin"
)

/* ---------------------------------- Setup --------------------------------- */

type Server struct {}

type Response struct {
	Status string      `json:"status"`
	Ok     bool        `json:"ok"`
	Message string      `json:"message,omitempty"`
	Data   interface{} `json:"data,omitempty"`
}

// Callback creator that enriches standard gin route handler with the slack and supabase clients
func createHandlerWithClients(slack *Slack, db *DB) func(func(ctx *gin.Context, slack *Slack, db *DB)) gin.HandlerFunc {
	return func(fn func(ctx *gin.Context, slack *Slack, db *DB)) gin.HandlerFunc {
		return func(ctx *gin.Context) {
			fn(ctx, slack, db)
		}
	}
}

func eventsHandler(ctx *gin.Context, client *Slack, db *DB) {
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

func (s* Server) New(env string, port string, client *Slack, db *DB) {
	if env == "production" {
		gin.SetMode(gin.ReleaseMode)
	}
	router := gin.New()
	handler := createHandlerWithClients(client, db)

	router.POST("/slack/match", handler(matchHandler))
	router.POST("/slack/create-game", handler(createGameHandler))
	router.POST("/slack/events", handler(eventsHandler))
	router.GET("/db/test", handler(testHandler))

	if port == "" {
		port = "8080"
	}

	if err := router.Run(":" + port); err != nil {
		panic(err)
	}
}

/* -------------------------------- Endpoints ------------------------------- */

func testHandler(ctx *gin.Context, client *Slack, db *DB) {
	response, err := client.GetSlackUsers()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, Response{Status: "error", Ok: false, Message: "Error fetching slack users" + err.Error()})
		return
	}
	
	ctx.JSON(http.StatusOK, Response{Status: "success", Ok: true, Data: response})
}

func matchHandler(ctx *gin.Context, client *Slack, db *DB) {
	resp := Response{Status: "success", Ok: true, Data: "Hello from Felo go app. Received /match command!"}
	client.PostMessage("C0859BH2E2W", "Hello from Felo go app. Received /match command!")
	ctx.JSON(http.StatusOK, resp)
}

func createGameHandler(ctx *gin.Context, slack *Slack, db *DB) {
	var payload struct {
		TriggerId string `json:"trigger_id"`
		Text string `json:"text"`
	}
	var buf bytes.Buffer
	if _, err := io.Copy(&buf, ctx.Request.Body); err != nil {
		ctx.JSON(http.StatusBadRequest, Response{Status: "error", Ok: false, Message: "Error reading request body"})
		return
	}
	ctx.Request.Body = io.NopCloser(&buf)

	if err := ctx.Request.ParseForm(); err != nil {
		ctx.JSON(http.StatusBadRequest, Response{Status: "error", Ok: false, Message: "Error parsing form"})
		return
	}
	payload.TriggerId = ctx.Request.Form.Get("trigger_id")
	payload.Text = ctx.Request.Form.Get("text")

	resp, err := slack.CreateGame(payload.TriggerId, payload.Text)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, Response{Status: "error", Ok: false, Message: "Error creating game" + err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, Response{Status: "success", Ok: true, Data: resp})
}