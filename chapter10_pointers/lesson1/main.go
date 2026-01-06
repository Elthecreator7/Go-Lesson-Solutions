package main

import "fmt"

type Message struct {
	Recipient string
	Text      string
}

func getMessageText(m Message) string {
	return fmt.Sprintf(`
To: %v
Message: %v
`, m.Recipient, m.Text)
}

func main() {
	// I added this part so the code actually runs
	m := Message{
		Recipient: "Caleb",
		Text:      "Checking my pointers",
	}

	output := getMessageText(m)
	fmt.Println(output)
}
