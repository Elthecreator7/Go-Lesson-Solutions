package main

import (
	"fmt"
	"sync"
	"testing"
)

func Test(t *testing.T) {
	type testCase struct {
		email string
		count int
	}

	runCases := []testCase{
		{"norman@bates.com", 23},
		{"marion@bates.com", 67},
	}

	submitCases := append(runCases, []testCase{
		{"lila@bates.com", 31},
		{"sam@bates.com", 453},
	}...)

	testCases := runCases
	if withSubmit {
		testCases = submitCases
	}

	skipped := len(submitCases) - len(testCases)

	passCount := 0
	failCount := 0

	for _, test := range testCases {
		sc := safeCounter{
			counts: make(map[string]int),
			mu:     &sync.Mutex{},
		}
		var wg sync.WaitGroup
		for i := 0; i < test.count; i++ {
			wg.Add(1)
			go func(email string) {
				sc.inc(email)
				wg.Done()
			}(test.email)
		}
		wg.Wait()

		if output := sc.val(test.email); output != test.count {
			failCount++
			t.Errorf(`---------------------------------
Test Failed:
  email: %v
  count: %v
  expected count: %v
  actual count:   %v
`, test.email, test.count, test.count, output)
		} else {
			passCount++
			fmt.Printf(`---------------------------------
Test Passed:
  email: %v
  count: %v
  expected count: %v
  actual count:   %v
`, test.email, test.count, test.count, output)
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
