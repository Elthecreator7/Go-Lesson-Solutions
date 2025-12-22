package main

import (
	"errors"
)

const (
	planFree = "free"
	planPro  = "pro"
)

func getMessageWithRetriesForPlan(plan string, messages [3]string) ([]string, error) {
	switch plan {
	case planPro:
		return messages[:], nil
	case planFree:
		return messages[0:2], nil
	default:
		return nil, errors.New("unsupported plan")
	}

	// Solution with if statement
	// if plan == planPro {
	// 	return messages[:], nil
	// } else if plan == planFree {
	// 	return messages[0:2], nil
	// } else {
	// 	return nil, errors.New("unsupported plan")
	// }
}
