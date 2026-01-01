package main

import (
	"fmt"
	"testing"
)

func TestDeleteIfNecessary(t *testing.T) {
	type testCase struct {
		users             map[string]user
		name              string
		expectedErrString string
		expectedUsers     map[string]user
		expectedDeleted   bool
	}

	runCases := []testCase{
		{
			map[string]user{"Erwin": {"Erwin", 14355550987, true}, "Levi": {"Levi", 98765550987, true}, "Hanji": {"Hanji", 18265554567, true}},
			"Erwin",
			"",
			map[string]user{"Levi": {"Levi", 98765550987, true}, "Hanji": {"Hanji", 18265554567, true}},
			true,
		},
		{
			map[string]user{"Erwin": {"Erwin", 14355550987, false}, "Levi": {"Levi", 98765550987, false}, "Hanji": {"Hanji", 18265554567, false}},
			"Erwin",
			"",
			map[string]user{"Erwin": {"Erwin", 14355550987, false}, "Levi": {"Levi", 98765550987, false}, "Hanji": {"Hanji", 18265554567, false}},
			false,
		},
	}

	submitCases := append(runCases, []testCase{
		{
			map[string]user{"Erwin": {"Erwin", 14355550987, true}, "Levi": {"Levi", 98765550987, true}, "Hanji": {"Hanji", 18265554567, true}},
			"Eren",
			"not found",
			map[string]user{"Erwin": {"Erwin", 14355550987, true}, "Levi": {"Levi", 98765550987, true}, "Hanji": {"Hanji", 18265554567, true}},
			false,
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
		originalUsers := getMapCopy(test.users)
		deleted, err := deleteIfNecessary(test.users, test.name)
		if test.expectedErrString != "" {
			if err == nil {
				failCount++
				t.Errorf(`---------------------------------
Test Failed:
  users:
%v
  name: %v
  expected error: %v
  actual error: none
`, formatMap(originalUsers), test.name, test.expectedErrString)
			} else if err.Error() != test.expectedErrString {
				failCount++
				t.Errorf(`---------------------------------
Test Failed:
  users:
%v
  name: %v
  expected error: %v
  actual error: %v
`, formatMap(originalUsers), test.name, test.expectedErrString, err)
			} else {
				passCount++
				fmt.Printf(`---------------------------------
Test Passed:
  users:
%v
  name: %v
  expected error: %v
  actual error: %v
`, formatMap(originalUsers), test.name, test.expectedErrString, err)
			}
		} else if err != nil {
			failCount++
			t.Errorf(`---------------------------------
Test Failed:
  users:
%v
  name: %v
  expected error: none
  actual error: %v
`, formatMap(originalUsers), test.name, err)
		} else if !compareMaps(test.users, test.expectedUsers) {
			failCount++
			t.Errorf(`---------------------------------
Test Failed:
  users:
%v
  name: %v
  expected users:
%v
  actual users:
%v
`, formatMap(originalUsers), test.name, formatMap(test.expectedUsers), formatMap(test.users))
		} else if deleted != test.expectedDeleted {
			failCount++
			t.Errorf(`---------------------------------
Test Failed:
  users:
%v
  name: %v
  expected deleted: %v
  actual deleted: %v
`, formatMap(originalUsers), test.name, test.expectedDeleted, deleted)
		} else {
			passCount++
			fmt.Printf(`---------------------------------
Test Passed:
  users:
%v
  name: %v
  expected users:
%v
  actual users:
%v
  expected deleted: %v
  actual deleted: %v
`, formatMap(originalUsers), test.name, formatMap(test.expectedUsers), formatMap(test.users), test.expectedDeleted, deleted)
		}
	}

	fmt.Println("---------------------------------")
	if skipped > 0 {
		fmt.Printf("%d passed, %d failed, %d skipped\n", passCount, failCount, skipped)
	} else {
		fmt.Printf("%d passed, %d failed\n", passCount, failCount)
	}

}

func getMapCopy(m map[string]user) map[string]user {
	copy := make(map[string]user)
	for key, value := range m {
		copy[key] = value
	}
	return copy
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
