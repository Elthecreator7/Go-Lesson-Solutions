package main

type cost struct {
	day   int
	value float64
}

func getDayCosts(costs []cost, day int) []float64 {
	dayCosts := make([]float64, 0)

	for i := 0; i < len(costs); i++ {
		if costs[i].day == day {
			dayCosts = append(dayCosts, costs[i].value)
		}
	}
	return dayCosts
}
