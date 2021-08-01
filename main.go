package main

import (
	"fmt"
	"log"
	"os"

	"jsonpbot-go/twitch"

	"github.com/joho/godotenv"
)

func LoadEnvVars(keys []string) map[string]string {
	var envVals = make(map[string]string)
	err := godotenv.Load()

	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	for _, key := range keys {
		fmt.Println(key)
		envVals[key] = os.Getenv(key)
	}

	return envVals
}

func main() {
	_ = godotenv.Load()
	c := twitch.Chat{}
	c.Init()
}
