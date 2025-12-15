package main

import (
	"fmt"
	"testing"
)

func getSenderLog(s sender) string {
	return fmt.Sprintf(`
====================================
Sender name: %v
Sender number: %v
Sender rateLimit: %v
====================================
`, s.name, s.number, s.rateLimit)
}

func Test(t *testing.T) {
	type testCase struct {
		rateLimit int
		name      string
		number    int
		expected  string
	}

	runCases := []testCase{
		{
			10000,
			"Deborah",
			18055558790,
			`
====================================
Sender name: Deborah
Sender number: 18055558790
Sender rateLimit: 10000
====================================
`,
		},
		{
			5000,
			"Jason",
			18055558791,
			`
====================================
Sender name: Jason
Sender number: 18055558791
Sender rateLimit: 5000
====================================
`,
		},
	}
	submitCases := append(runCases, []testCase{
		{
			1000,
			"Jill",
			18055558792,
			`
====================================
Sender name: Jill
Sender number: 18055558792
Sender rateLimit: 1000
====================================
`,
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
		output := getSenderLog(sender{
			rateLimit: test.rateLimit,
			user: user{
				name:   test.name,
				number: test.number,
			},
		})
		if output != test.expected {
			failCount++
			t.Errorf(`---------------------------------
Inputs:     (%v, %v, %v)
Expecting:  %v
Actual:     %v
Fail
`, test.rateLimit, test.name, test.number, test.expected, output)
		} else {
			passCount++
			fmt.Printf(`---------------------------------
Inputs:     (%v, %v, %v)
Expecting:  %v
Actual:     %v
Pass
`, test.rateLimit, test.name, test.number, test.expected, output)
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
