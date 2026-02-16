package modules

import (
	"fmt"
	"time"

	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
)

func ReadMessage() {
	consumer, err := kafka.NewConsumer(&kafka.ConfigMap{
		"bootstrap.servers": "localhost",
		"group.id": "notification",
	})

	if err != nil {
		panic(err)
	}

	err = consumer.SubscribeTopics([]string{"notifications"}, nil)

	if err != nil {
		panic(err)
	}

	for true {
		msg, err := consumer.ReadMessage(time.Second)
		if err == nil {
			Notification.UseNotification(Notification{}, "discord", msg)

			if err != nil {
				fmt.Println("Error to send notification: ", err)
			} else {
				fmt.Printf("Message on %s: %s", msg.TopicPartition, msg.Value)
			}

		} else if !err.(kafka.Error).IsTimeout() {
			fmt.Printf("Consumer error: %v\n", err)
		}
	}

	consumer.Close()
}
