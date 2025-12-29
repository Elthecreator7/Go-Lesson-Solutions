package main

import (
	"fmt"
	"testing"
)

func Test(t *testing.T) {
	type testCase struct {
		msg      []string
		badWords []string
		expected int
	}

	runCases := []testCase{
		{[]string{"hey", "there", "john"}, []string{"crap", "shoot", "frick", "dang"}, -1},
		{[]string{"ugh", "oh", "my", "frick"}, []string{"crap", "shoot", "frick", "dang"}, 3},
	}

	submitCases := append(runCases, []testCase{
		{[]string{"what", "the", "shoot", "I", "hate", "that", "crap"}, []string{"crap", "shoot", "frick", "dang"}, 2},
		{[]string{"crap", "shoot", "frick", "dang"}, []string{""}, -1},
		{[]string{""}, nil, -1},
	}...)

	testCases := runCases
	if withSubmit {
		testCases = submitCases
	}

	skipped := len(submitCases) - len(testCases)

	passCount := 0
	failCount := 0

	for _, test := range testCases {
		output := indexOfFirstBadWord(test.msg, test.badWords)
		if output != test.expected {
			failCount++
			t.Errorf(`---------------------------------
Test Failed:
message:
%v
bad words:
%v
Expecting:  %v
Actual:     %v
Fail
`, sliceWithBullets(test.msg), sliceWithBullets(test.badWords), test.expected, output)
		} else {
			passCount++
			fmt.Printf(`---------------------------------
Test Passed:
message:
%v
bad words:
%v
Expecting:  %v
Actual:     %v
Pass
`, sliceWithBullets(test.msg), sliceWithBullets(test.badWords), test.expected, output)
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
