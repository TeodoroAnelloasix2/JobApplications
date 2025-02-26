/*
Assignment

We've been asked to add a feature that extracts costs for a given day.
Complete the getDayCosts() function using the append() function.
It accepts a slice of cost structs and a day int, and it returns a float64 slice containing that day's costs.
If there are no costs for that day, return an empty slice.
*/
package main

type cost struct {
	day   int
	value float64
}

func getDayCosts(costs []cost, day int) []float64 {

	costs_of_given_day := make([]float64, 0, len(costs))

	if len(costs) <= 0 {
		return costs_of_given_day
	}
	for i := 0; i < len(costs); i++ {
		if costs[i].day == day {
			costs_of_given_day = append(costs_of_given_day, costs[i].value)
		}

	}
	return costs_of_given_day

}
