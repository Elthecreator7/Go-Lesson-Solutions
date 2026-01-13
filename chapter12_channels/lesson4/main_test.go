package main

import (
	"fmt"
	"testing"
)

func Test(t *testing.T) {
	type testCase struct {
		emails   []string
		expected int
	}

	runCases := []testCase{
		{
			[]string{
				"To boldly go where no man has gone before.",
				"Live long and prosper.",
			},
			2,
		},
		{
			[]string{
				"The needs of the many outweigh the needs of the few, or the one.",
				"Change is the essential process of all existence.",
				"Resistance is futile.",
			},
			3,
		},
	}

	submitCases := append(runCases, []testCase{
		{
			[]string{
				"It's life, Jim, but not as we know it.",
				"Infinite diversity in infinite combinations.",
				"Make it so.",
				"Engage!",
			},
			4,
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
		ch := addEmailsToQueue(test.emails)
		actual := len(ch)
		if actual != test.expected {
			failCount++
			t.Errorf(`---------------------------------
Test Failed:
  emails:
%v
  expected channel length: %v
  actual channel length:   %v
`,
				sliceWithBullets(test.emails),
				test.expected,
				actual)
		} else {
			passCount++
			fmt.Printf(`---------------------------------
Test Passed:
  emails:
%v
  expected channel length: %v
  actual channel length:   %v
`,
				sliceWithBullets(test.emails),
				test.expected,
				actual)
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
