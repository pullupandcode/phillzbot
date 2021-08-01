package utils

import (
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	env map[string]string
}

var c *Config

func init() {
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}
}

func GetVar(key string) string {
	envVar, ok := c.env[key]

	if ok {
		return envVar
	}

	envVar = os.Getenv(key)
	c.env[key] = envVar
	return envVar
}
