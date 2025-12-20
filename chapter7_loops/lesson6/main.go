package main

func countConnections(groupSize int) int {
	var numOfConn int
	for i := 1; i < groupSize; i++ {
		numOfConn += i
	}
	return numOfConn
}
