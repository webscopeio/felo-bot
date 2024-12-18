package main

import (
	// std
	"fmt"
	"os"
	// local modules
	"webscope.io/felo/modules/server"
	"webscope.io/felo/modules/slack"
	// external modules
	"github.com/joho/godotenv"
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

	localEnv, envErr := godotenv.Read(".env")

	for _, key := range envKeys {
		// Atempt to get OS level ENV variable
		val := os.Getenv(key)
		if val == "" {
			// If not found, attempt to read from .env file
			if envErr != nil || localEnv[key] == "" {
				return ENV{}, fmt.Errorf("missing environment variable %s", key)
			} else {
				val = localEnv[key]
			}
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
