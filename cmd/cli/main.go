package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"messageService/cmd/handlers"
	"messageService/domain"
	"messageService/service"
	"os"
)

func main() {
	configFile := os.Args[1:]

	filePath, err := handlers.ErrorHandler(configFile)
	if err != nil {
		fmt.Println(err)
	}

	jsonFile, err := os.Open(filePath)
	if err != nil {
		fmt.Println(err)
	}
	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)
	var flatters []domain.Flatters

	if err := json.Unmarshal(byteValue, &flatters); err != nil {
		fmt.Println(err)
		return
	}

	service.ManageReceivers(flatters)
}
