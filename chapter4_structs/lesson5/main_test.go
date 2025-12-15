package main

import (
	"fmt"
	"reflect"
	"testing"
)

func Test(t *testing.T) {
	type testCase struct {
		name     string
		expected uintptr
	}

	runCases := []testCase{
		{"contact", uintptr(24)},
		{"perms", uintptr(16)},
	}

	submitCases := append(runCases, []testCase{}...)

	skipped := len(submitCases) - len(submitCases)
	passCount := 0
	failCount := 0

	for _, test := range submitCases {
		var typ reflect.Type
		if test.name == "contact" {
			typ = reflect.TypeOf(contact{})
		} else if test.name == "perms" {
			typ = reflect.TypeOf(perms{})
		}

		size := typ.Size()

		if size != test.expected {
			failCount++
			t.Errorf(`---------------------------------
Inputs:     (%v)
Expecting:  %v bytes
Actual:     %v bytes
Fail
`, test.name, test.expected, size)
		} else {
			passCount++
			fmt.Printf(`---------------------------------
Inputs:     (%v)
Expecting:  %v bytes
Actual:     %v bytes
Pass
`, test.name, test.expected, size)
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
