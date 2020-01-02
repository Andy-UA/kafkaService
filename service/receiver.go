package service

import (
	"context"
	"fmt"
	"github.com/segmentio/kafka-go"
	"messageService/domain"
)

func invokeConsumer(inputTopic string, destinationTopic string) {
	r := kafka.NewReader(kafka.ReaderConfig{
		Brokers:   []string{domain.KafkaServerAddress},
		GroupID:   domain.GroupID,
		Topic:     inputTopic,
		MinBytes:  domain.MinByteMessage,
		MaxBytes:  domain.MaxByteMessage,
	})

	status <- fmt.Sprintf("Consumer of topic '%s' has been started", inputTopic)

	for {
		m, err := r.ReadMessage(context.Background())
		if err != nil {
			status <- fmt.Sprintf("Consumer of topic '%s' failed with error message: '%s'", inputTopic, err.Error())
			break
		}
		go receiveMessage(m.Value, m.Topic, destinationTopic)
	}

	r.Close()
}