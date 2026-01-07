package main

type Analytics struct {
	MessagesTotal     int
	MessagesFailed    int
	MessagesSucceeded int
}

type Message struct {
	Recipient string
	Success   bool
}

func analyzeMessage(analytics *Analytics, message Message) {
	if message.Success {
		analytics.MessagesSucceeded++
		analytics.MessagesTotal++
	} else {
		analytics.MessagesFailed++
		analytics.MessagesTotal++
	}
}
