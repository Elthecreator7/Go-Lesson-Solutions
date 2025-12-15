package main

import (
	"fmt"
	"testing"
	"time"
)

func Test(t *testing.T) {
	type testCase struct {
		msg          message
		expectedText string
		expectedCost int
	}

	runCases := []testCase{
		{birthdayMessage{time.Date(1994, 03, 21, 0, 0, 0, 0, time.UTC), "John Doe"},
			"Hi John Doe, it is your birthday on 1994-03-21T00:00:00Z",
			168,
		},
		{sendingReport{"First Report", 10},
			`Your "First Report" report is ready. You've sent 10 messages.`,
			183,
		},
	}

	submitCases := append(runCases, []testCase{
		{birthdayMessage{time.Date(1934, 05, 01, 0, 0, 0, 0, time.UTC), "Bill Deer"},
			"Hi Bill Deer, it is your birthday on 1934-05-01T00:00:00Z",
			171,
		},
		{sendingReport{"Second Report", 20},
			`Your "Second Report" report is ready. You've sent 20 messages.`,
			186,
		},
	}...)

	testCases := runCases
	if withSubmit {
		testCases = submitCases
	}

	skipped := len(submitCases) - len(testCases)

	passCount := 0
	failCount := 0

	for _, test := range testCases {
		text, cost := sendMessage(test.msg)
		if text != test.expectedText || cost != test.expectedCost {
			failCount++
			t.Errorf(`---------------------------------
Inputs:     %+v
Expecting:  (%v, %v)
Actual:     (%v, %v)
Fail
`, test.msg, test.expectedText, test.expectedCost, text, cost)
		} else {
			passCount++
			fmt.Printf(`---------------------------------
Inputs:     %+v
Expecting:  (%v, %v)
Actual:     (%v, %v)
Pass
`, test.msg, test.expectedText, test.expectedCost, text, cost)
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
