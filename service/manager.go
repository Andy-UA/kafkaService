package service

import (
	"encoding/json"
	"fmt"
	"github.com/segmentio/kafka-go"
	"messageService/domain"
)

var status = make(chan string)

func ManageReceivers(flatters []domain.Flatters) {
	for _, flatter := range flatters {
		go invokeConsumer(flatter.InputTopic, flatter.DestinationTopic)
	}

	for {
		if v, ok := <-status; ok {
			fmt.Println(v)
		}
	}
}

func receiveMessage(message []byte, inputTopic string, destinationTopic string) {
	status <- fmt.Sprintf("Message has been received from topic '%s' and preparing to be sent into '%s' topic",
		inputTopic, destinationTopic)

	var inputMessage domain.InputMessage

	if err := json.Unmarshal(message, &inputMessage); err != nil {
		fmt.Println(err)
		return
	}

	convertToOutputMessage(inputMessage, destinationTopic)
}

func convertToOutputMessage(inputMessage domain.InputMessage, destinationTopic string) {
	var outputMessages = make([]domain.OutputMessage, len(inputMessage.Message.Partitions))

	for i, partition := range inputMessage.Message.Partitions {
		outputMessages[i].PayloadData.Name = partition.Name
		outputMessages[i].PayloadData.DriveType = partition.DriveType
		outputMessages[i].PayloadData.UsedSpaceBytes = partition.Metric.UsedSpaceBytes
		outputMessages[i].PayloadData.TotalSpaceBytes = partition.Metric.TotalSpaceBytes
		outputMessages[i].PayloadData.CreateAtTimeUTC = inputMessage.Message.CreateAtTimeUTC
	}

	sendMessage(outputMessages, destinationTopic)
}

func marshalMessage(c chan<- kafka.Message, outputMessages []domain.OutputMessage) {
	for _, outputMessage := range outputMessages {
		message, err := json.Marshal(outputMessage)
		if err != nil {
			fmt.Println(err)
			return
		}

		c <- kafka.Message{
			Value: message,
		}
	}
	close(c)
}
