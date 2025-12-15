package main

import (
	"fmt"
	"slices"
	"testing"
)

func Test(t *testing.T) {
	type testCase struct {
		input    []int
		expected []int
	}

	runCases := []testCase{
		{
			input:    []int{1, 2, 3},
			expected: []int{1, 3, 6},
		},
		{
			input:    []int{1, 2, 3, 4, 5},
			expected: []int{1, 3, 6, 10, 15},
		},
	}
	submitCases := append(runCases, []testCase{
		{
			input:    []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
			expected: []int{1, 3, 6, 10, 15, 21, 28, 36, 45, 55},
		},
		{
			input:    []int{0, 0, 0, 0},
			expected: []int{0, 0, 0, 0},
		},
		{
			input:    []int{5, -3, -1},
			expected: []int{5, 2, 1},
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
		f := adder()
		results := make([]int, len(test.input))
		for i, v := range test.input {
			results[i] = f(v)
		}
		if !slices.Equal(results, test.expected) {
			failCount++
			t.Errorf(`---------------------------------
Inputs:     %v
Expecting:  %v
Actual:     %v
Fail
`, test.input, test.expected, results)
		} else {
			passCount++
			fmt.Printf(`---------------------------------
Inputs:     %v
Expecting:  %v
Actual:     %v
Pass
`, test.input, test.expected, results)
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
