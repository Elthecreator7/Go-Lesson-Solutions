package main

import (
	"fmt"
	"testing"
)

func Test(t *testing.T) {
	type testCase struct {
		nums     []int
		expected int
	}

	runCases := []testCase{
		{[]int{1, 2, 3}, 6},
		{[]int{1, 2, 3, 4, 5}, 15},
	}

	submitCases := append(runCases, []testCase{
		{[]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, 55},
		{[]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15}, 120},
		{[]int{}, 0},
		{[]int{5}, 5},
	}...)

	testCases := runCases
	if withSubmit {
		testCases = submitCases
	}

	skipped := len(submitCases) - len(testCases)

	passCount := 0
	failCount := 0

	for _, test := range testCases {
		output := sum(test.nums...)
		if output != test.expected {
			failCount++
			t.Errorf(`---------------------------------
Inputs:
%v
Expecting:  %v
Actual:     %v
Fail
`, sliceWithBullets(test.nums), test.expected, output)
		} else {
			passCount++
			fmt.Printf(`---------------------------------
Inputs:
%v
Expecting:  %v
Actual:     %v
Pass
`, sliceWithBullets(test.nums), test.expected, output)
		}
	}

	fmt.Println("---------------------------------")
	if skipped > 0 {
		fmt.Printf("%d passed, %d failed, %d skipped\n", passCount, failCount, skipped)
	} else {
		fmt.Printf("%d passed, %d failed\n", passCount, failCount)
	}
}

func sliceWithBullets[T any](slice []T) string {
	if slice == nil {
		return "  <nil>"
	}
	if len(slice) == 0 {
		return "  []"
	}
	output := ""
	for i, item := range slice {
		form := "  - %#v\n"
		if i == (len(slice) - 1) {
			form = "  - %#v"
		}
		output += fmt.Sprintf(form, item)
	}
	return output
}

// withSubmit is set at compile time depending
// on which button is used to run the tests
var withSubmit = true
