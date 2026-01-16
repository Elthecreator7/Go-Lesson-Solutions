package main

func getLast[T any](s []T) T {
	var zeroVal T
	if len(s) == 0 {
		return zeroVal
	}
	lastElem := len(s) - 1
	return s[lastElem]

}
