package main

import (
	// std
	"fmt"
	"net/http"

	// internal modules
	"webscope.io/felo/modules"
	"webscope.io/felo/modules/utils"
)

func main() {
	env, err := utils.ReadEnv()
	if err != nil {
		fmt.Println("Error reading .env file")
		panic(err)
	}
	supabase := &modules.Supabase{
		SUPABASE_KEY: env.SUPABASE_KEY,
		SUPABASE_URL: env.SUPABASE_URL,
	}
	db, err := supabase.Init(nil)
	if err != nil {
		fmt.Printf("Error initializing Supabase client")
		panic(err)
	}
	slack := &modules.Slack{
		BOT_TOKEN: env.BOT_TOKEN,
		API_URL: "https://slack.com/api",
		HTTP_CLIENT: &http.Client{},
	}
	server := &modules.Server{}
	server.New(env.ENV, env.PORT, slack, db)
}
