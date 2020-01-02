package service

import (
	"context"
	"fmt"
	"github.com/segmentio/kafka-go"
	"messageService/domain"
)

func sendMessage(outputMessages []domain.OutputMessage, topic string) {
	w := kafka.NewWriter(kafka.WriterConfig{
		Brokers: []string{domain.KafkaServerAddress},
		Topic:   topic,
		Balancer: &kafka.LeastBytes{},
	})

	kafkaMessages := make([]kafka.Message, len(outputMessages))
	msgch := make(chan kafka.Message)

	go marshalMessage(msgch, outputMessages)

	for message := range msgch {
		kafkaMessages = append(kafkaMessages, message)
	}

	w.WriteMessages(context.Background(), kafkaMessages...)
	w.Close()

	status <- fmt.Sprintf("Message has been sent to a topic '%s'", topic)
}
