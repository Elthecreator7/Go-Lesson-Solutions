package main

import (
	"fmt"
	"testing"
)

func Test(t *testing.T) {
	type testCase struct {
		msgToCustomer string
		msgToSpouse   string
		expectedCost  int
		expectedErr   error
	}

	runCases := []testCase{
		{"Thanks for coming in to our flower shop today!", "We hope you enjoyed your gift.", 0, fmt.Errorf("can't send texts over 25 characters")},
		{"Thanks for joining us!", "Have a good day.", 76, nil},
	}

	submitCases := append(runCases, []testCase{
		{"Thank you.", "Enjoy!", 32, nil},
		{"We loved having you in!", "We hope the rest of your evening is fantastic.", 0, fmt.Errorf("can't send texts over 25 characters")},
	}...)

	testCases := runCases
	if withSubmit {
		testCases = submitCases
	}

	skipped := len(submitCases) - len(testCases)
	passCount := 0
	failCount := 0

	for _, test := range testCases {
		cost, err := sendSMSToCouple(test.msgToCustomer, test.msgToSpouse)
		errString := ""
		if err != nil {
			errString = err.Error()
		}
		expectedErrString := ""
		if test.expectedErr != nil {
			expectedErrString = test.expectedErr.Error()
		}
		if cost != test.expectedCost || errString != expectedErrString {
			failCount++
			t.Errorf(`---------------------------------
Inputs:     (%v, %v)
Expecting:  (%v, %v)
Actual:     (%v, %v)
Fail
`, test.msgToCustomer, test.msgToSpouse, test.expectedCost, test.expectedErr, cost, err)
		} else {
			passCount++
			fmt.Printf(`---------------------------------
Inputs:     (%v, %v)
Expecting:  (%v, %v)
Actual:     (%v, %v)
Pass
`, test.msgToCustomer, test.msgToSpouse, test.expectedCost, test.expectedErr, cost, err)
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
