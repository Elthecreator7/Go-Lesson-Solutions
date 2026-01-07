package main

import (
	"strings"
)

func removeProfanity(message *string) {
	messageVal := *message
	if message == nil {
		return
	}
	messageVal = strings.ReplaceAll(messageVal, "fubb", "****")
	messageVal = strings.ReplaceAll(messageVal, "shiz", "****")
	messageVal = strings.ReplaceAll(messageVal, "witch", "*****")
	*message = messageVal
}
