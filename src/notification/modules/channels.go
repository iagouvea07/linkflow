package modules

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
)

var webhook string = "https://discord.com/api/webhooks/1473026622121574452/0VLooYL5jsSgARaDpEi5YlguHEzIMiYSvppjjHX824Um3hwtp_K8q84YKjMh6HmoUcsQ"

type Notification struct {
	Content string 	`json:"content"`
}

func (notification Notification)SendNotificationToDiscord(message *kafka.Message) error{

	notification.Content = string(message.Value)

	body, _ := json.Marshal(notification)
	payload := bytes.NewBuffer(body)

	res, err := http.Post(webhook, "application/json", payload)

	if err != nil {
		log.Fatal(err)
		return err
	}

	io.ReadAll(res.Body)

	defer res.Body.Close()

	log.Printf("Content: %s", notification.Content)
	return nil
}

func (notification Notification) UseNotification(channel string, message *kafka.Message) {
	switch channel {
		case "discord":
			notification.SendNotificationToDiscord(message)
		case "*":
			fmt.Println("channel not found!")
	}
}