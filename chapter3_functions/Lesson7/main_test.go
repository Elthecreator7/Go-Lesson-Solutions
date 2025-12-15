package main

import (
	"fmt"
	"testing"
)

func Test(t *testing.T) {
	type testCase struct {
		age                   int
		exYearsUntilAdult     int
		exYearsUntilDrinking  int
		exYearsUntilCarRental int
	}

	runCases := []testCase{
		{4, 14, 17, 21},
		{18, 0, 3, 7},
		{22, 0, 0, 3},
	}

	submitCases := append(runCases, []testCase{
		{25, 0, 0, 0},
	}...)

	testCases := runCases
	if withSubmit {
		testCases = submitCases
	}

	skipped := len(submitCases) - len(testCases)

	passCount := 0
	failCount := 0

	for _, test := range testCases {
		yearsUntilAdult, yearsUntilDrinking, yearsUntilCarRental := yearsUntilEvents(test.age)
		if yearsUntilAdult != test.exYearsUntilAdult ||
			yearsUntilDrinking != test.exYearsUntilDrinking ||
			yearsUntilCarRental != test.exYearsUntilCarRental {
			failCount++
			t.Errorf(`---------------------------------
Inputs:     (%v)
Expecting:  (%v, %v, %v)
Actual:     (%v, %v, %v)
Fail
`, test.age, test.exYearsUntilAdult, test.exYearsUntilDrinking, test.exYearsUntilCarRental, yearsUntilAdult, yearsUntilDrinking, yearsUntilCarRental)
		} else {
			passCount++
			fmt.Printf(`---------------------------------
Inputs:     (%v)
Expecting:  (%v, %v, %v)
Actual:     (%v, %v, %v)
Pass
`, test.age, test.exYearsUntilAdult, test.exYearsUntilDrinking, test.exYearsUntilCarRental, yearsUntilAdult, yearsUntilDrinking, yearsUntilCarRental)
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
