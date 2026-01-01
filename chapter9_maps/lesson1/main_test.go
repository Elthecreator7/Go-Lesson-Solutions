package main

import (
	"fmt"
	"testing"
)

func Test(t *testing.T) {
	type testCase struct {
		names        []string
		phoneNumbers []int
		expected     map[string]user
		errString    string
	}
	runCases := []testCase{
		{
			[]string{"Eren", "Armin", "Mikasa"},
			[]int{14355550987, 98765550987, 18265554567},
			map[string]user{"Eren": {"Eren", 14355550987}, "Armin": {"Armin", 98765550987}, "Mikasa": {"Mikasa", 18265554567}},
			"",
		},
		{
			[]string{"Eren", "Armin"},
			[]int{14355550987, 98765550987, 18265554567},
			nil,
			"invalid sizes",
		},
	}
	submitCases := append(runCases, []testCase{
		{
			[]string{"George", "Annie", "Reiner", "Sasha"},
			[]int{20955559812, 38385550982, 48265554567, 16045559873},
			map[string]user{"George": {"George", 20955559812}, "Annie": {"Annie", 38385550982}, "Reiner": {"Reiner", 48265554567}, "Sasha": {"Sasha", 16045559873}},
			"",
		},
		{
			[]string{"George", "Annie", "Reiner"},
			[]int{20955559812, 38385550982, 48265554567, 16045559873},
			nil,
			"invalid sizes",
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
		output, err := getUserMap(test.names, test.phoneNumbers)
		if test.errString != "" && err == nil {
			failCount++
			t.Errorf(`---------------------------------
Test Failed:
  names: %v
  phoneNumbers: %v
  expected err: %v
  actual err: none
`, test.names, test.phoneNumbers, test.errString)
		} else if test.errString == "" && err != nil {
			failCount++
			t.Errorf(`---------------------------------
Test Failed:
  names: %v
  phoneNumbers: %v
  expected err: none
  actual err: %v
`, test.names, test.phoneNumbers, err)
		} else if test.errString != "" && err != nil && err.Error() != test.errString {
			failCount++
			t.Errorf(`---------------------------------
Test Failed:
  names: %v
  phoneNumbers: %v
  expected err: %v
  actual err: %v
`, test.names, test.phoneNumbers, test.errString, err)
		} else if !compareMaps(output, test.expected) {
			failCount++
			t.Errorf(`---------------------------------
Test Failed:
  names: %v
  phoneNumbers: %v
  expected:
%v
  actual:
%v
`, test.names, test.phoneNumbers, formatMap(test.expected), formatMap(output))
		} else {
			passCount++
			fmt.Printf(`---------------------------------
Test Passed:
  names: %v
  phoneNumbers: %v
  expected:
%v
  actual:
%v
`, test.names, test.phoneNumbers, formatMap(test.expected), formatMap(output))
		}
	}
	fmt.Println("---------------------------------")
	if skipped > 0 {
		fmt.Printf("%d passed, %d failed, %d skipped\n", passCount, failCount, skipped)
	} else {
		fmt.Printf("%d passed, %d failed\n", passCount, failCount)
	}
}

func formatMap(m map[string]user) string {
	var str string
	for key, value := range m {
		str += fmt.Sprintf("  * %s: %v\n", key, value)
	}
	return str
}

func compareMaps(map1, map2 map[string]user) bool {
	if len(map1) != len(map2) {
		return false
	}
	for key, value1 := range map1 {
		if value2, exist := map2[key]; !exist || value1 != value2 {
			return false
		}
	}
	return true
}

// withSubmit is set at compile time depending
// on which button is used to run the tests
var withSubmit = true
