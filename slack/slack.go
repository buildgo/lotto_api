package slack

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
)

func getSlackToken() string {
	return os.Getenv("SLACK_TOKEN")
}

func getSlackWebhookUrl() string {
	return os.Getenv("SLACK_WEBHOOK_URL")
}

func sendMessageWithWebhook(webhook, message string) {
	url := getSlackWebhookUrl()
	bodyMap := make(map[string]string)
	bodyMap["text"] = message
	jsonString, _ := json.Marshal(bodyMap)

	log.Println(fmt.Sprintf("Send message: %v", string(jsonString)))

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonString))
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
}

func SendMessage(message string) {
	webhook := getSlackWebhookUrl()
	sendMessageWithWebhook(webhook, message)
}
