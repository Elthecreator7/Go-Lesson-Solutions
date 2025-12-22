package main

import (
	"fmt"
	"testing"
)

func Test(t *testing.T) {
	type testCase struct {
		messages         []string
		expectedMessages [3]string
		expectedCosts    [3]int
	}

	runCases := []testCase{
		{
			[]string{
				"Hello sir/madam can I interest you in a yacht?",
				"Please I'll even give you an Amazon gift card?",
				"You're missing out big time",
			},
			[3]string{
				"Hello sir/madam can I interest you in a yacht?",
				"Please I'll even give you an Amazon gift card?",
				"You're missing out big time",
			},
			[3]int{46, 92, 119},
		},
		{
			[]string{"It's the spring fling sale!", "Don't miss this event!", "Last chance."},
			[3]string{"It's the spring fling sale!", "Don't miss this event!", "Last chance."},
			[3]int{27, 49, 61},
		},
	}

	submitCases := append(runCases, []testCase{
		{
			[]string{
				"Put that coffee down!",
				"Coffee is for closers",
				"Always be closing",
			},
			[3]string{
				"Put that coffee down!",
				"Coffee is for closers",
				"Always be closing",
			},
			[3]int{21, 42, 59},
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
		actualMessages, actualCosts := getMessageWithRetries(test.messages[0], test.messages[1], test.messages[2])
		if actualMessages[0] != test.expectedMessages[0] ||
			actualMessages[1] != test.expectedMessages[1] ||
			actualMessages[2] != test.expectedMessages[2] ||
			actualCosts[0] != test.expectedCosts[0] ||
			actualCosts[1] != test.expectedCosts[1] ||
			actualCosts[2] != test.expectedCosts[2] {
			failCount++
			t.Errorf(`---------------------------------
Test Failed:
Inputs:
%v
Expecting:
%v
%v
Actual:
%v
%v
Fail
`, sliceWithBullets(test.messages), sliceWithBullets(test.expectedMessages[:]), test.expectedCosts, sliceWithBullets(actualMessages[:]), actualCosts)
		} else {
			passCount++
			fmt.Printf(`---------------------------------
Test Passed:
Inputs:
%v
Expecting:
%v
%v
Actual:
%v
%v
Pass
`, sliceWithBullets(test.messages), sliceWithBullets(test.expectedMessages[:]), test.expectedCosts, sliceWithBullets(actualMessages[:]), actualCosts)
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
	output := ""
	for i, item := range slice {
		form := "  - %v\n"
		if i == (len(slice) - 1) {
			form = "  - %v"
		}
		output += fmt.Sprintf(form, item)
	}
	return output
}

// withSubmit is set at compile time depending
// on which button is used to run the tests
var withSubmit = true
