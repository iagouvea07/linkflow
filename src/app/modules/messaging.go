package modules

import (
	"fmt"

	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
)

type Message struct{
	Content string `json:"content"`
}

type UrlEncoded struct {
	ID uint64 		`json:"id"`
	URL string  	`json:"url"`
	ENCODE string 	`json:"encode"`	
}

func SendMessage(messageContent []byte){
	producer, err := kafka.NewProducer(&kafka.ConfigMap{
		"bootstrap.servers": "localhost",		
	})

	topic := "notifications"

	if err != nil {
		panic(err)
	}

	err = producer.Produce(&kafka.Message{
		TopicPartition: kafka.TopicPartition{Topic: &topic, Partition: kafka.PartitionAny},
		Value: messageContent,
	}, nil)

	if err != nil {
		panic(err)
	}

	fmt.Println("Message Sended: %s", string(messageContent))
	producer.Flush(1000)
	defer producer.Close()
}