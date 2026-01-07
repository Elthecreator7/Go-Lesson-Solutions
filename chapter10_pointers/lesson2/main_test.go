package main

import (
	"fmt"
	"testing"
)

func Test(t *testing.T) {
	type testCase struct {
		messageIn string
		expected  string
	}

	runCases := []testCase{
		{
			"English, motherfubber, do you speak it?",
			"English, mother****er, do you speak it?",
		},
		{
			"Oh man I've seen some crazy ass shiz in my time...",
			"Oh man I've seen some crazy ass **** in my time...",
		},
	}

	submitCases := append(runCases, []testCase{
		{
			"Does he look like a witch?",
			"Does he look like a *****?",
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
		original := test.messageIn
		removeProfanity(&test.messageIn)
		if test.messageIn != test.expected {
			failCount++
			t.Errorf(`---------------------------------
Test Failed:
  input:    %v
  expected: %v
  actual:   %v
`, original, test.expected, test.messageIn)
		} else {
			passCount++
			fmt.Printf(`---------------------------------
Test Passed:
  input:    %v
  expected: %v
  actual:   %v
`, original, test.expected, test.messageIn)
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
