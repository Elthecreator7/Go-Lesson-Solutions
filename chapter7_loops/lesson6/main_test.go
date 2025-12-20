package main

import (
	"fmt"
	"testing"
)

func TestCountGroupConnections(t *testing.T) {
	type testCase struct {
		groupSize int
		expected  int
	}

	runCases := []testCase{
		{1, 0},
		{2, 1},
		{3, 3},
		{4, 6},
	}

	submitCases := append(runCases, []testCase{
		{0, 0},
		{10, 45},
		{100, 4950},
		{1000, 499500},
	}...)

	testCases := runCases
	if withSubmit {
		testCases = submitCases
	}

	skipped := len(submitCases) - len(testCases)

	passCount := 0
	failCount := 0

	for _, test := range testCases {
		result := countConnections(test.groupSize)
		if result != test.expected {
			failCount++
			t.Errorf(`---------------------------------
Group Size: %v
Expecting: %v
Actual:    %v
Fail
`, test.groupSize, test.expected, result)
		} else {
			passCount++
			fmt.Printf(`---------------------------------
Group Size: %v
Expecting: %v
Actual:    %v
Pass
`, test.groupSize, test.expected, result)
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
