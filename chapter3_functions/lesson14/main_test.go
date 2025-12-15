package main

import (
	"fmt"
	"testing"
)

func TestSplitEmail(t *testing.T) {
	type testCase struct {
		email    string
		username string
		domain   string
	}

	runCases := []testCase{
		{"drogon@dragonstone.com", "drogon", "dragonstone.com"},
		{"rhaenyra@targaryen.com", "rhaenyra", "targaryen.com"},
	}

	submitCases := append(runCases, []testCase{
		{"viserys@kingslanding.com", "viserys", "kingslanding.com"},
		{"aegon@stormsend.com", "aegon", "stormsend.com"},
	}...)

	testCases := runCases
	if withSubmit {
		testCases = submitCases
	}

	passCount := 0
	failCount := 0
	skipped := len(submitCases) - len(testCases)

	for _, test := range testCases {
		username, domain := splitEmail(test.email)
		if username != test.username || domain != test.domain {
			failCount++
			t.Errorf(`---------------------------------
Inputs:     (%v)
Expecting:  (%v, %v)
Actual:     (%v, %v)
Fail
`, test.email, test.username, test.domain, username, domain)
		} else {
			passCount++
			fmt.Printf(`---------------------------------
Inputs:     (%v)
Expecting:  (%v, %v)
Actual:     (%v, %v)
Pass
`, test.email, test.username, test.domain, username, domain)
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
