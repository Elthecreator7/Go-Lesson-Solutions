package main

import (
	"fmt"
	"testing"
)

func TestHandleEmailBounce(t *testing.T) {
	type testCase struct {
		email           email
		expectedError   string
		expectedStatus  string
		expectedBounces int
	}

	runCases := []testCase{
		{
			email: email{
				status:    "email_bounced",
				recipient: &user{email: "bugs@acme.inc"},
			},
			expectedError:   "<nil>",
			expectedStatus:  "email_bounced",
			expectedBounces: 1,
		},
		{
			email: email{
				status:    "email_failed",
				recipient: &user{email: "elmer@acme.inc"},
			},
			expectedError:   "error tracking user bounce: invalid event: email_failed",
			expectedStatus:  "email_failed",
			expectedBounces: 0,
		},
		{
			email: email{
				status:    "email_sent",
				recipient: &user{email: "daffy@acme.inc"},
			},
			expectedError:   "error updating user status: invalid status: email_sent",
			expectedStatus:  "",
			expectedBounces: 0,
		},
	}

	submitCases := append(runCases, []testCase{
		{
			email: email{
				status:    "email_failed",
				recipient: &user{email: "porky@acme.inc"},
			},
			expectedError:   "error tracking user bounce: invalid event: email_failed",
			expectedStatus:  "email_failed",
			expectedBounces: 0,
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
		a := &analytics{}
		err := a.handleEmailBounce(test.email)
		actualError := fmt.Sprintf("%v", err)
		if actualError != test.expectedError {
			failCount++
			t.Errorf(`---------------------------------
Test Failed:
  status:    %v
  recipient: %v
  expected error:   %v
  actual error:     %v
`, test.email.status, test.email.recipient.email, test.expectedError, actualError)
		} else {
			passCount++
			fmt.Printf(`---------------------------------
Test Passed:
  status:    %v
  recipient: %v
  expected error:   %v
  actual error:     %v
`, test.email.status, test.email.recipient.email, test.expectedError, actualError)
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
