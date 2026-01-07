package main

import (
	"strings"
)

func removeProfanity(message *string) {
	msgPtr := message
	firstWord := strings.ReplaceAll(*msgPtr, "fubb", "****")
	secondWord := strings.ReplaceAll(firstWord, "shiz", "****")
	finalWord := strings.ReplaceAll(secondWord, "witch", "*****")
	*msgPtr = finalWord

}
