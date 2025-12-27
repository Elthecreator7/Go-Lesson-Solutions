package main

import (
	"fmt"
	"reflect"
	"testing"
)

func Test(t *testing.T) {
	type testCase struct {
		costs    []cost
		day      int
		expected []float64
	}

	runCases := []testCase{
		{
			costs: []cost{
				{0, 1.0},
				{1, 2.0},
				{1, 3.1},
				{5, 2.5},
				{2, 3.6},
				{1, 2.7},
				{1, 3.3},
			},
			day: 1,
			expected: []float64{
				2.0,
				3.1,
				2.7,
				3.3,
			},
		},
	}

	submitCases := append(runCases, []testCase{
		{
			costs: []cost{
				{0, 1.0},
				{1, 2.0},
				{1, 3.1},
				{2, 2.5},
				{3, 3.1},
				{3, 2.6},
				{4, 3.34},
			},
			day: 4,
			expected: []float64{
				3.34,
			},
		},
		{
			costs: []cost{
				{0, 1.0},
				{10, 2.0},
				{3, 3.1},
				{2, 2.5},
				{1, 3.6},
				{2, 2.7},
				{4, 56.34},
				{13, 2.34},
				{28, 1.34},
				{25, 2.34},
				{30, 4.34},
			},
			day:      5,
			expected: []float64{},
		},
	}...)

	testCases := runCases
	if withSubmit {
		testCases = submitCases
	}

	passCount := 0
	failCount := 0
	skipped := len(submitCases) - len(testCases)

	for _, test := range testCases {
		output := getDayCosts(test.costs, test.day)
		if !reflect.DeepEqual(output, test.expected) {
			failCount++
			t.Errorf(`---------------------------------
Inputs:
%v
Expecting:
%v
Actual:
%v
Fail
`, sliceWithBullets(test.costs), sliceWithBullets(test.expected), sliceWithBullets(output))
		} else {
			passCount++
			fmt.Printf(`---------------------------------
Inputs:     %v
Expecting:
%v
Actual:
%v
Pass
`, sliceWithBullets(test.costs), sliceWithBullets(test.expected), sliceWithBullets(output))
		}
	}

	fmt.Println("---------------------------------")
	if skipped > 0 {
		fmt.Printf("%d passed, %d failed, %d skipped\n", passCount, failCount, skipped)
	} else {
		fmt.Printf("%d passed, %d failed\n", passCount, failCount)
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
		form := "  - %v\n"
		if i == (len(slice) - 1) {
			form = "  - %v"
		}
		output += fmt.Sprintf(form, item)
	}
	return output
}

// withSubmit is set at compile time depending
// on which button is used to run the tests
var withSubmit = true
