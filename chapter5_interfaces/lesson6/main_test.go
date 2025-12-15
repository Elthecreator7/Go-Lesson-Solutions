package main

import (
	"fmt"
	"strconv"
	"testing"
)

func TestSendMessage(t *testing.T) {
	type testCase struct {
		format   formatter
		expected string
	}

	runCases := []testCase{
		{plainText{message: "Hello, World!"}, "Hello, World!"},
		{bold{message: "Bold Message"}, "**Bold Message**"},
		{code{message: "Code Message"}, "`Code Message`"},
	}

	submitCases := append(runCases, []testCase{
		{code{message: ""}, "``"},
		{bold{message: ""}, "****"},
		{plainText{message: ""}, ""},
	}...)

	testCases := runCases
	if withSubmit {
		testCases = submitCases
	}
	skipped := len(submitCases) - len(testCases)

	passCount := 0
	failCount := 0

	for i, test := range testCases {
		testName := "Test Case " + strconv.Itoa(i+1)
		t.Run(testName, func(t *testing.T) {
			formattedMessage := sendMessage(test.format)
			if formattedMessage != test.expected {
				failCount++
				t.Errorf(`---------------------------------
%s
Inputs:     (%v)
Expecting:  %v
Actual:     %v
Fail
`, testName, test.format, test.expected, formattedMessage)
			} else {
				passCount++
				fmt.Printf(`---------------------------------
%s
Inputs:     (%v)
Expecting:  %v
Actual:     %v
Pass
`, testName, test.format, test.expected, formattedMessage)
			}
		})
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
