package main

func adder() func(int) int {
	sum := 0

	return func(input int) int {
		sum += input

		return sum
	}
	// ?
}
