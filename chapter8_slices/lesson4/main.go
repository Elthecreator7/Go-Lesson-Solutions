package main

func sum(nums ...int) int {
	var costs int
	for i := 0; i < len(nums); i++ {
		costs += nums[i]
	}
	return costs
}
