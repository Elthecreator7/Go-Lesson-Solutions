package main

import (
	"fmt"
	"testing"
)

func TestUpdateBalance(t *testing.T) {
	type testCase struct {
		name            string
		initialCustomer customer
		transaction     transaction
		expectedBalance float64
		expectError     bool
		errorMessage    string
	}

	runCases := []testCase{
		{
			name:            "Deposit operation",
			initialCustomer: customer{id: 1, balance: 100.0},
			transaction:     transaction{customerID: 1, amount: 50.0, transactionType: transactionDeposit},
			expectedBalance: 150.0,
			expectError:     false,
		},
		{
			name:            "Withdrawal operation",
			initialCustomer: customer{id: 2, balance: 200.0},
			transaction:     transaction{customerID: 2, amount: 100.0, transactionType: transactionWithdrawal},
			expectedBalance: 100.0,
			expectError:     false,
		},
	}

	submitCases := append(runCases, []testCase{
		{
			name:            "insufficient funds for withdrawal",
			initialCustomer: customer{id: 3, balance: 50.0},
			transaction:     transaction{customerID: 3, amount: 100.0, transactionType: transactionWithdrawal},
			expectedBalance: 50.0,
			expectError:     true,
			errorMessage:    "insufficient funds",
		},
		{
			name:            "unknown transaction type",
			initialCustomer: customer{id: 4, balance: 100.0},
			transaction:     transaction{customerID: 4, amount: 50.0, transactionType: "unknown"},
			expectedBalance: 100.0,
			expectError:     true,
			errorMessage:    "unknown transaction type",
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
		t.Run(test.name, func(t *testing.T) {
			err := updateBalance(&test.initialCustomer, test.transaction)
			failureMessage := ""

			if (err != nil) != test.expectError {
				failureMessage += "Unexpected error presence: expected an error but didn't get one, or vice versa.\n"
			} else if err != nil && err.Error() != test.errorMessage {
				failureMessage += "Incorrect error message.\n"
			}

			if test.initialCustomer.balance != test.expectedBalance {
				failureMessage += "Balance not updated as expected.\n"
			}

			if failureMessage != "" {
				failCount++
				failureMessage = "FAIL\n" + failureMessage +
					"Transaction: " + string(test.transaction.transactionType) +
					fmt.Sprintf(", Amount: %.2f\n", test.transaction.amount) +
					fmt.Sprintf("Expected balance: %.2f, Actual balance: %.2f", test.expectedBalance, test.initialCustomer.balance)
				t.Errorf(`---------------------------------
				%s
`, failureMessage)
			} else {
				passCount++
				successMessage := "PASSED\n" +
					"Transaction: " + string(test.transaction.transactionType) +
					fmt.Sprintf(", Amount: %.2f\n", test.transaction.amount) +
					fmt.Sprintf("Expected balance: %.2f, Actual balance: %.2f", test.expectedBalance, test.initialCustomer.balance)
				fmt.Printf(`---------------------------------
%s
`, successMessage)
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
