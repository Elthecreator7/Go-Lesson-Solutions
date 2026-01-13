package main

func countReports(numSentCh chan int) int {
	numReports := 0
	for {
		val, ok := <-numSentCh
		if !ok {
			break
		} else {
			numReports += val
		}
	}
	return numReports
}

// don't touch below this line

func sendReports(numBatches int, ch chan int) {
	for i := 0; i < numBatches; i++ {
		numReports := i*23 + 32%17
		ch <- numReports
	}
	close(ch)
}
