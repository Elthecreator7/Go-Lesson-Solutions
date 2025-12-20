package main

func countConnections(groupSize int) int {
	var numOfConn int
	for i := 1; i < groupSize; i++ {
		numOfConn += i
	}
	return numOfConn
}

// Alternative solution with math formula
// func countConnections(groupSize int) int {
// 	numOfConn := (groupSize * (groupSize - 1)) / 2
// 	return numOfConn
// }
