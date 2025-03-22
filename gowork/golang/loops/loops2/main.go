/*
Assignment
Complete the maxMessages function. Given a cost threshold, it should calculate the maximum number of messages that can be sent.
Each message costs 100 pennies, plus an additional fee. The fee structure is:

	1st message: 100 + 0
	2nd message: 100 + 1
	3rd message: 100 + 2
	4th message: 100 + 3
*/

package main

func maxMessages(thresh int) int {
	max_msg := 0

	for i := 0; ; i++ {

		coste_mensaje := 100 + i

		if thresh < coste_mensaje {
			break
		}

		max_msg = max_msg + 1

		thresh = thresh - coste_mensaje

	}

	return max_msg

}
