package main

import (
	"fmt"
	"testing"
)

func Test(t *testing.T) {
	type testCase struct {
		expense      expense
		expectedTo   string
		expectedCost float64
	}

	runCases := []testCase{
		{
			email{isSubscribed: true, body: "Whoa there!", toAddress: "soldier@monty.com"},
			"soldier@monty.com",
			0.11,
		},
		{
			sms{isSubscribed: false, body: "Halt! Who goes there?", toPhoneNumber: "+155555509832"},
			"+155555509832",
			2.1,
		},
	}

	submitCases := append(runCases, []testCase{
		{
			email{
				isSubscribed: false,
				body:         "It is I, Arthur, son of Uther Pendragon, from the castle of Camelot. King of the Britons, defeator of the Saxons, sovereign of all England!",
				toAddress:    "soldier@monty.com",
			},
			"soldier@monty.com",
			6.95,
		},
		{
			email{
				isSubscribed: true,
				body:         "Pull the other one!",
				toAddress:    "arthur@monty.com",
			},
			"arthur@monty.com",
			0.19,
		},
		{
			sms{
				isSubscribed:  true,
				body:          "I am. And this my trusty servant Patsy.",
				toPhoneNumber: "+155555509832",
			},
			"+155555509832",
			1.17,
		},
		{
			invalid{},
			"",
			0.0,
		},
	}...)

	testCases := runCases
	if withSubmit {
		testCases = submitCases
	}

	passCount := 0
	failCount := 0
	skipped := len(submitCases) - len(testCases)

	for _, test := range testCases {
		to, cost := getExpenseReport(test.expense)
		if to != test.expectedTo || cost != test.expectedCost {
			failCount++
			t.Errorf(`---------------------------------
Inputs:     %+v
Expecting:  (%v, %v)
Actual:     (%v, %v)
Fail
`, test.expense, test.expectedTo, test.expectedCost, to, cost)
		} else {
			passCount++
			fmt.Printf(`---------------------------------
Inputs:     %+v
Expecting:  (%v, %v)
Actual:     (%v, %v)
Pass
`, test.expense, test.expectedTo, test.expectedCost, to, cost)
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
