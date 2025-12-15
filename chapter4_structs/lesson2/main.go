package main

type messageToSend struct {
	message   string
	sender    user
	recipient user
}

type user struct {
	name   string
	number int
}

func canSendMessage(mToSend messageToSend) bool {
	if mToSend.sender.name != "" && mToSend.recipient.name != "" && mToSend.sender.number != 0 && mToSend.recipient.number != 0 {
		return true
	}
	return false
	// ?
}
