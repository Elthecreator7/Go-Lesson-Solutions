package main

func updateCounts(messagedUsers []string, validUsers map[string]int) {
	for i := 0; i < len(messagedUsers); i++ {
		userName := messagedUsers[i]
		if _, ok := validUsers[userName]; !ok {
			continue
		} else {
			validUsers[userName]++
		}
	}
}
