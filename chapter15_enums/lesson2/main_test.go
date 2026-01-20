package main

import (
	"fmt"
	"testing"
)

func TestEmailStatus(t *testing.T) {
	type testCase struct {
		status   emailStatus
		expected string
	}

	runCases := []testCase{
		{EmailBounced, "EmailBounced"},
		{EmailInvalid, "EmailInvalid"},
		{EmailDelivered, "EmailDelivered"},
	}

	submitCases := append(runCases, []testCase{
		{EmailOpened, "EmailOpened"},
		{17, "Unknown"},
	}...)

	testCases := runCases
	if withSubmit {
		testCases = submitCases
	}

	skipped := len(submitCases) - len(testCases)

	passCount := 0
	failCount := 0

	for _, test := range testCases {
		output := getEmailStatusName(test.status)
		if output != test.expected {
			failCount++
			t.Errorf(`---------------------------------
Test Failed:
  status:   %v
  expected: %v
  actual:   %v
`, test.status, test.expected, output)
		} else {
			passCount++
			fmt.Printf(`---------------------------------
Test Passed:
  status:   %v
  expected: %v
  actual:   %v
`, test.status, test.expected, output)
		}
	}

	fmt.Println("---------------------------------")
	if skipped > 0 {
		fmt.Printf("%d passed, %d failed, %d skipped\n", passCount, failCount, skipped)
	} else {
		fmt.Printf("%d passed, %d failed\n", passCount, failCount)
	}

}

func getEmailStatusName(status emailStatus) string {
	switch status {
	case EmailBounced:
		return "EmailBounced"
	case EmailInvalid:
		return "EmailInvalid"
	case EmailDelivered:
		return "EmailDelivered"
	case EmailOpened:
		return "EmailOpened"
	default:
		return "Unknown"
	}
}

// withSubmit is set at compile time depending
// on which button is used to run the tests
var withSubmit = true
