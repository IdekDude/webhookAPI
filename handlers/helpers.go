package handlers

import (
	"os"
	"strings"
)

func LoadWebhooks(key string) []string {
	webhook := os.Getenv(key)

	hookArray := strings.Split(webhook, ",")
	return hookArray
}