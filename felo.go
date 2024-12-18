package main

import (
	"fmt"
	"os"
	"strings"
	"webscope.io/felo/modules/server"
	"webscope.io/felo/modules/slack"
)

type ENV struct {
	BOT_TOKEN string
	HOST string
	ENV string
}

func readEnv() (ENV, error) {
	file, err := os.ReadFile(".env")
	if err != nil {
		fmt.Println("Error reading .env file")
		panic(err)
	}
	lines := strings.Split(string(file), "\n")
	envMap := make(map[string]string)

	for _, line := range lines {
		if len(line) == 0 {
			continue
		}
		parts := strings.Split(line, "=")
		if len(parts) != 2 {
			continue
		}
		key := strings.TrimSpace(parts[0])
		value := strings.TrimSpace(parts[1])
		if len(key) == 0 || len(value) == 0 {
			return ENV{}, fmt.Errorf("invalid .env file. Please check the file and correct the mistakes")
		}
		envMap[key] = value
	}

	envKeys := []string{
		"BOT_TOKEN",
		"HOST",
		"ENV",
	}

	for _, key := range envKeys {
		if _, ok := envMap[key]; !ok {
			return ENV{}, fmt.Errorf("missing key %s in .env file", key)
		}
	}

	return ENV{
		BOT_TOKEN: envMap["BOT_TOKEN"],
		HOST: envMap["HOST"],
		ENV: envMap["ENV"],
	}, nil
}

func main() {
	env, err := readEnv()
	if err != nil {
		fmt.Println("Error reading .env file")
		panic(err)
	}
	client := slack.New(env.BOT_TOKEN)
	server.New(env.HOST, env.ENV, client)
}
