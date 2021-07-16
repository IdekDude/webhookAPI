package handlers

import (
	"fmt"
	"github.com/joho/godotenv"
	"os"
	"strings"
)

func LoadWebhooks(key string) []string {
	if err := godotenv.Load(".env"); err != nil {
		fmt.Println(err)
	}
	webhook := os.Getenv(key)

	hookArray := strings.Split(webhook, ",")
	return hookArray
}