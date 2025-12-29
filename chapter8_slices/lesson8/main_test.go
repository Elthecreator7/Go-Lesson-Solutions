package main

import (
	"fmt"
	"testing"
)

func TestFilterMessages(t *testing.T) {
	messages := []Message{
		TextMessage{"Alice", "Hello, World!"},
		MediaMessage{"Bob", "image", "A beautiful sunset"},
		LinkMessage{"Charlie", "http://example.com", "Example Domain"},
		TextMessage{"Dave", "Another text message"},
		MediaMessage{"Eve", "video", "Cute cat video"},
		LinkMessage{"Frank", "https://boot.dev", "Learn Coding Online"},
	}
	type testCase struct {
		filterType    string
		expectedCount int
		expectedType  string
	}

	runCases := []testCase{
		{"text", 2, "text"},
		{"media", 2, "media"},
		{"link", 2, "link"},
	}

	submitCases := append(runCases, []testCase{
		{"media", 2, "media"},
		{"text", 2, "text"},
	}...)

	testCases := runCases
	if withSubmit {
		testCases = submitCases
	}

	skipped := len(submitCases) - len(testCases)

	passCount := 0
	failCount := 0

	for i, test := range testCases {
		t.Run(fmt.Sprintf("TestCase%d", i+1), func(t *testing.T) {
			filtered := filterMessages(messages, test.filterType)
			if len(filtered) != test.expectedCount {
				failCount++
				t.Errorf(`---------------------------------
Test Case %d - Filtering for %s
Expecting:  %d messages
Actual:     %d messages
Fail
`, i+1, test.filterType, test.expectedCount, len(filtered))
			} else {
				passCount++
				fmt.Printf(`---------------------------------
Test Case %d - Filtering for %s
Expecting:  %d messages
Actual:     %d messages
Pass
`, i+1, test.filterType, test.expectedCount, len(filtered))
			}

			for _, m := range filtered {
				if m.Type() != test.expectedType {
					failCount++
					t.Errorf(`---------------------------------
Test Case %d - Message Type Check
Expecting:  %s message
Actual:     %s message
Fail
`, i+1, test.expectedType, m.Type())
				} else {
					passCount++
					fmt.Printf(`---------------------------------
Test Case %d - Message Type Check
Expecting:  %s message
Actual:     %s message
Pass
`, i+1, test.expectedType, m.Type())
				}
			}
		})
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
