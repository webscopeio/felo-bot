package utils

import (
	//std
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

type ENV struct {
	BOT_TOKEN    string
	ENV          string
	PORT         string
	SUPABASE_URL string
	SUPABASE_KEY string
}

func ReadEnv() (ENV, error) {
	envMap := map[string]string{}
	envKeys := []string{
		"BOT_TOKEN",
		"ENV",
		"PORT",
		"SUPABASE_URL",
		"SUPABASE_KEY",
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
		BOT_TOKEN:    envMap["BOT_TOKEN"],
		ENV:          envMap["ENV"],
		PORT:         envMap["PORT"],
		SUPABASE_URL: envMap["SUPABASE_URL"],
		SUPABASE_KEY: envMap["SUPABASE_KEY"],
	}, nil
}
