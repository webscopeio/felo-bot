package main

import (
	"fmt"
	"os"
	"webscope.io/felo/modules/server"
	"webscope.io/felo/modules/slack"
)

type ENV struct {
	BOT_TOKEN string
	ENV       string
	PORT      string
}

func readEnv() (ENV, error) {
	envMap := map[string]string{}
	envKeys := []string{
		"BOT_TOKEN",
		"ENV",
		"PORT",
	}

	for _, key := range envKeys {
		val := os.Getenv(key)
		if val == "" {
			return ENV{}, fmt.Errorf("Missing environment variable %s", key)
		}
		envMap[key] = val
	}
	return ENV{
		BOT_TOKEN: envMap["BOT_TOKEN"],
		ENV:       envMap["ENV"],
		PORT:      envMap["PORT"],
	}, nil
}

func main() {
	env, err := readEnv()
	if err != nil {
		fmt.Println("Error reading .env file")
		panic(err)
	}
	client := slack.New(env.BOT_TOKEN)
	server.New(env.ENV, env.PORT, client)
}
