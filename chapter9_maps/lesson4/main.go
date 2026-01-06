package main

func getNameCounts(names []string) map[rune]map[string]int {
	nameCounts := make(map[rune]map[string]int)
	for i := 0; i < len(names); i++ {
		runes := []rune(names[i])
		firstChar := runes[0]
		if _, ok := nameCounts[firstChar]; !ok {
			firstCharMap := make(map[string]int)
			nameCounts[firstChar] = firstCharMap
		}
		nameCounts[firstChar][names[i]]++
	}
	return nameCounts
}
