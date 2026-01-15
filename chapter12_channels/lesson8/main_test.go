package main

import (
	"fmt"
	"slices"
	"testing"
	"time"
)

func TestSaveBackups(t *testing.T) {
	type testCase struct {
		expectedLogs []string
	}

	runCases := []testCase{
		{
			expectedLogs: []string{
				"Nothing to do, waiting...",
				"Nothing to do, waiting...",
				"Taking a backup snapshot...",
				"Nothing to do, waiting...",
				"Nothing to do, waiting...",
				"Taking a backup snapshot...",
				"Nothing to do, waiting...",
				"Taking a backup snapshot...",
				"Nothing to do, waiting...",
				"All backups saved!",
			},
		},
	}

	submitCases := append(runCases, []testCase{}...)

	testCases := runCases
	if withSubmit {
		testCases = submitCases
	}

	skipped := len(submitCases) - len(testCases)

	passed, failed := 0, 0
	for _, test := range testCases {
		expectedLogs := test.expectedLogs

		snapshotTicker := time.Tick(800 * time.Millisecond)
		saveAfter := time.After(2800 * time.Millisecond)
		logChan := make(chan string)
		go saveBackups(snapshotTicker, saveAfter, logChan)
		actualLogs := []string{}
		for actualLog := range logChan {
			fmt.Println(actualLog)
			actualLogs = append(actualLogs, actualLog)
		}

		if !slices.Equal(expectedLogs, actualLogs) {
			t.Errorf(`---------------------------------
Test Failed:
expected:
%v
actual:
%v
`, sliceWithBullets(expectedLogs), sliceWithBullets(actualLogs))
			failed++
		} else {
			fmt.Printf(`---------------------------------
Test Passed:
expected:
%v
actual:
%v
`, sliceWithBullets(expectedLogs), sliceWithBullets(actualLogs))
			passed++
		}
	}

	fmt.Println("---------------------------------")
	if skipped > 0 {
		fmt.Printf("\n%d passed, %d failed, %d skipped\n", passed, failed, skipped)
	} else {
		fmt.Printf("\n%d passed, %d failed\n", passed, failed)
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
