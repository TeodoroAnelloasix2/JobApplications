/*

Assignment
We need to sum up the costs of all individual messages so we can send an end-of-month bill to our customers.
Complete the sum function to return the sum of all inputs.
Take note of how the variadic inputs and the spread operator are used in the test suite.

*/

package main

func sum(nums ...int) int {

	var bill_cost float64
	nums_copy := nums[:]

	for i := 0; i < len(nums_copy); i++ {
		bill_cost += float64(nums_copy[i])
	}
	return int(bill_cost)
}
