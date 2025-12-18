package main

import (
	"fmt"
	"testing"
)

func TestDivide(t *testing.T) {
	type testCase struct {
		x, y, expected float64
		expectedErr    string
	}

	runCases := []testCase{
		{10, 0, 0, "no dividing by 0"},
		{10, 2, 5, ""},
		{15, 30, 0.5, ""},
		{6, 3, 2, ""},
	}

	submitCases := append(runCases, []testCase{
		{0, 10, 0, ""},
		{100, 0, 0, "no dividing by 0"},
		{-10, -2, 5, ""},
		{-10, 2, -5, ""},
	}...)

	testCases := runCases
	if withSubmit {
		testCases = submitCases
	}

	skipped := len(submitCases) - len(testCases)

	passCount := 0
	failCount := 0

	for _, test := range testCases {
		result, err := divide(test.x, test.y)
		errString := ""
		if err != nil {
			errString = err.Error()
		}
		if result != test.expected || errString != test.expectedErr {
			failCount++
			t.Errorf(`---------------------------------
Inputs:     (%v, %v)
Expecting:  (%v, %v)
Actual:     (%v, %v)
Fail
`, test.x, test.y, test.expected, test.expectedErr, result, errString)
		} else {
			passCount++
			fmt.Printf(`---------------------------------
Inputs:     (%v, %v)
Expecting:  (%v, %v)
Actual:     (%v, %v)
Pass
`, test.x, test.y, test.expected, test.expectedErr, result, errString)
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
