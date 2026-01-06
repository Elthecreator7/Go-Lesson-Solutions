package main

import "strings"

func countDistinctWords(messages []string) int {
	distinctWordStruct := make(map[string]struct{})
	for i := 0; i < len(messages); i++ {
		mWithoutWhiteSpaces := strings.Fields(messages[i])
		for j := 0; j < len(mWithoutWhiteSpaces); j++ {
			word := strings.ToLower(mWithoutWhiteSpaces[j])
			distinctWordStruct[word] = struct{}{}
		}
	}
	distinctWordCount := len(distinctWordStruct)
	return distinctWordCount
}
