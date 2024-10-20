package main

import (
	"fmt"
	"os"

	"github.com/go-resty/resty/v2"
)

// sendLineNotification sends a notification to LINE
func sendLineNotification(message string) error {
	token := os.Getenv("LINE_NOTIFY_TOKEN") // Assume the token is set in .env

	client := resty.New()
	resp, err := client.R().
		SetHeader("Authorization", "Bearer "+token).
		SetFormData(map[string]string{"message": message}).
		Post("https://notify-api.line.me/api/notify")

	if err != nil {
		return fmt.Errorf("failed to send notification: %w", err)
	}

	if resp.StatusCode() != 200 {
		return fmt.Errorf("LINE Notify returned error: %s", resp.String())
	}

	return nil
}
