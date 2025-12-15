package main

func reformat(message string, formatter func(string) string) string {
	firstMessage := formatter(message)
	secondMessage := formatter(firstMessage)
	thirdMessage := formatter(secondMessage)

	finalMessage := "TEXTIO: " + thirdMessage

	return finalMessage
}
