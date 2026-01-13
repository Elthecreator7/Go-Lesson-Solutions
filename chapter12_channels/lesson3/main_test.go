package main

import (
	"fmt"
	"testing"
)

func Test(t *testing.T) {
	type testCase struct {
		numDBs int
	}

	runCases := []testCase{
		{1},
		{3},
		{4},
	}

	submitCases := append(runCases, []testCase{
		{0},
		{13},
	}...)

	testCases := runCases
	if withSubmit {
		testCases = submitCases
	}
	skipped := len(submitCases) - len(testCases)

	passed, failed := 0, 0

	for _, test := range testCases {
		fmt.Printf(`---------------------------------`)
		fmt.Printf("\nTesting %v Databases...\n\n", test.numDBs)
		dbChan, count := getDBsChannel(test.numDBs)
		waitForDBs(test.numDBs, dbChan)
		for *count != test.numDBs {
			fmt.Println("...")
		}
		if len(dbChan) == 0 && *count == test.numDBs {
			passed++
			fmt.Printf(`
expected length: 0, count: %v
actual length:   %v, count: %v
PASS
`,
				test.numDBs, len(dbChan), *count)
		} else {
			failed++
			fmt.Printf(`
expected length: 0, count: %v
actual length:   %v, count: %v
FAIL
`,
				test.numDBs, len(dbChan), *count)
		}
	}

	fmt.Println("---------------------------------")
	if skipped > 0 {
		fmt.Printf("%d passed, %d failed, %d skipped\n", passed, failed, skipped)
	} else {
		fmt.Printf("%d passed, %d failed\n", passed, failed)
	}

}

// withSubmit is set at compile time depending
// on which button is used to run the tests
var withSubmit = true
