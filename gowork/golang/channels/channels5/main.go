/*
Assignment

At Mailio we're all about keeping track of what our systems are up to with great logging and telemetry.

The sendReports function sends out a batch of reports to our clients and reports back how many were sent across a channel.
It closes the channel when it's done.

Complete the countReports function. It should:

	Use an infinite for loop to read from the channel:
	If the channel is closed, break out of the loop
	Otherwise, keep a running total of the number of reports sent
	Return the total number of reports sent
*/
package main

func countReports(numSentCh chan int) int {
	total := 0
	for {
		value, ok := <-numSentCh
		if !ok {
			break
		}
		total += value
	}
	return total
}

// don't touch below this line

func sendReports(numBatches int, ch chan int) {
	for i := 0; i < numBatches; i++ {
		numReports := i*23 + 32%17
		ch <- numReports
	}
	close(ch)
}
