package slack

import "net/http"

type Client struct {
	BOT_TOKEN  string
	API_URL string
	HTTP_CLIENT *http.Client
}
