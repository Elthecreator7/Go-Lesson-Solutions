package main

import (
	"fmt"
	"testing"
)

func Test(t *testing.T) {
	type testCase struct {
		mToSend  messageToSend
		expected bool
	}

	runCases := []testCase{
		{messageToSend{
			message:   "you have an appointment tomorrow",
			sender:    user{name: "Brenda Halafax", number: 16545550987},
			recipient: user{name: "Sally Sue", number: 19035558973},
		}, true},
		{messageToSend{
			message:   "you have an event tomorrow",
			sender:    user{number: 16545550987},
			recipient: user{name: "Suzie Sall", number: 19035558973},
		}, false},
	}

	submitCases := append(runCases, []testCase{
		{messageToSend{
			message:   "you have an birthday tomorrow",
			sender:    user{name: "Jason Bjorn", number: 16545550987},
			recipient: user{name: "Jim Bond"},
		}, false},
		{messageToSend{
			message:   "you have a party tomorrow",
			sender:    user{name: "Njorn Halafax"},
			recipient: user{name: "Becky Sue", number: 19035558973},
		}, false},
		{messageToSend{
			message:   "you have a birthday tomorrow",
			sender:    user{name: "Eli Halafax", number: 16545550987},
			recipient: user{number: 19035558973},
		}, false},
	}...)

	testCases := runCases
	if withSubmit {
		testCases = submitCases
	}

	skipped := len(submitCases) - len(testCases)
	passCount := 0
	failCount := 0

	for _, test := range testCases {
		output := canSendMessage(test.mToSend)
		if output != test.expected {
			failCount++
			t.Errorf(`---------------------------------
Inputs:
  * message:          %s
  * sender.name:      %s
  * sender.number:    %d
  * recipient.name:   %s
  * recipient.number: %d
  Expected:           %v
  Actual:             %v
Fail
`,
				test.mToSend.message,
				test.mToSend.sender.name,
				test.mToSend.sender.number,
				test.mToSend.recipient.name,
				test.mToSend.recipient.number,
				test.expected,
				output)
		} else {
			passCount++
			fmt.Printf(`---------------------------------
Inputs:
  * message:          %s
  * sender.name:      %s
  * sender.number:    %d
  * recipient.name:   %s
  * recipient.number: %d
  Expected:           %v
  Actual:             %v
Pass
`,
				test.mToSend.message,
				test.mToSend.sender.name,
				test.mToSend.sender.number,
				test.mToSend.recipient.name,
				test.mToSend.recipient.number,
				test.expected,
				output)
		}
	}

	fmt.Println("---------------------------------")
	if skipped > 0 {
		fmt.Printf("%d passed, %d failed, %d skipped\n", passCount, failCount, skipped)
	} else {
		fmt.Printf("%d passed, %d failed\n", passCount, failCount)
	}

}

// withSubmit is set at compile time depending
// on which button is used to run the tests
var withSubmit = true
