package main

func getMessageCosts(messages []string) []float64 {
	msgCosts := make([]float64, len(messages))
	for i := 0; i < len(messages); i++ {
		cost := float64(len(messages[i])) * 0.01
		msgCosts[i] = cost
	}
	return msgCosts
}
