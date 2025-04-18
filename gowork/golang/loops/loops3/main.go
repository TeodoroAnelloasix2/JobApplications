/*
Assignment

We have an interesting new cost structure from our SMS vendor.
They charge exponentially more money for each consecutive text we send! Let's write a function that calculates how many

	messages we can send in a given batch given a costMultiplier and a maxCostInPennies.

In a nutshell, the first message costs a penny, and each message after that costs the same as the previous
message multiplied by the costMultiplier.

There is a bug in the code! Add a condition to the for loop to fix the bug.
The loop should stop when balance is equal to or less than 0. So what condition should the for loop evaluate?
*/
package main

func getMaxMessagesToSend(costMultiplier float64, maxCostInPennies int) int {
	actualCostInPennies := 1.0
	maxMessagesToSend := 1
	balance := float64(maxCostInPennies) - actualCostInPennies

	for balance >= actualCostInPennies {
		actualCostInPennies *= costMultiplier
		balance -= actualCostInPennies

		if balance < 0 {
			maxMessagesToSend--
			break
		}

		maxMessagesToSend++

	}

	return maxMessagesToSend
}
