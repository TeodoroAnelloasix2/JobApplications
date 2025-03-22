/*
Assignment

Retries are a premium feature now!
Textio's free users only get 1 retry message, while pro members get an unlimited amount.

Complete the getMessageWithRetriesForPlan function.

It takes a plan variable as input as well as an array of 3 messages.
You've been provided with constants representing the plan types at the top of the file.
If the plan is a pro plan, return all the strings from the messages input in a slice.
If the plan is a free plan, return the first 2 strings from the messages input in a slice.
If the plan isn't either of those, return a nil slice and an error that says unsupported plan.

primes := [6]int{2, 3, 5, 7, 11, 13}
mySlice := primes[1:4]
// mySlice = {3, 5, 7}

The syntax is:

arrayname[lowIndex:highIndex]
arrayname[lowIndex:]
arrayname[:highIndex]
arrayname[:]

# Alternativa para resolver ejercicio

	if plan == planPro {
	    msg_list = append(msg_list, messages[:]...)
	} else if plan == planFree {

	    msg_list = append(msg_list, messages[:2]...)
	}
*/
package main

import "errors"

const (
	planFree = "free"
	planPro  = "pro"
)

func getMessageWithRetriesForPlan(plan string, messages [3]string) ([]string, error) {

	var msg_list []string

	if plan != planFree && plan != planPro {
		err_to_send := errors.New("unsupported plan")
		return nil, err_to_send
	}

	if plan == planPro {

		for i := 0; i < 3; i++ {

			msg_list = append(msg_list, messages[i])
		}
	} else if plan == planFree {

		for i := 0; i < 2; i++ {

			msg_list = append(msg_list, messages[i])
		}

	}

	return msg_list, nil
}
