package main

import (
	"fmt"
	"testing"
)

func TestValidateStatus(t *testing.T) {
	type testCase struct {
		status      string
		expectedErr string
	}

	runCases := []testCase{
		{"", "status cannot be empty"},
		{"This is a valid status update that is well within the character limit.", ""},
		{"This status update is way too long. Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam, quis nostrud exercitation ullamco.", "status exceeds 140 characters"},
	}

	submitCases := append(runCases, []testCase{
		{"Another valid status.", ""},
		{"This status update, while derivative, contains exactly one hundred and forty-one characters, which is over the status update character limit.", "status exceeds 140 characters"},
	}...)

	testCases := runCases
	if withSubmit {
		testCases = submitCases
	}

	skipped := len(submitCases) - len(testCases)
	passCount := 0
	failCount := 0

	for _, test := range testCases {
		err := validateStatus(test.status)
		errString := ""
		if err != nil {
			errString = err.Error()
		}
		if errString != test.expectedErr {
			failCount++
			t.Errorf(`---------------------------------
Inputs:     "%v"
Expecting:  "%v"
Actual:     "%v"
Fail
`, test.status, test.expectedErr, errString)
		} else {
			passCount++
			fmt.Printf(`---------------------------------
Inputs:     "%v"
Expecting:  "%v"
Actual:     "%v"
Pass
`, test.status, test.expectedErr, errString)
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
