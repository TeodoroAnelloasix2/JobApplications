/*
Assignment

We send a lot of text messages at Textio, and our API is getting slow and unresponsive.
If we know the rough size of a slice before we fill it up,
we can make our program faster by creating the slice with that size ahead of time so that
the Go runtime doesn't need to continuously allocate new underlying arrays of larger and larger sizes.
By setting the length, the slice can still be resized later, but it means we can avoid all the expensive resizing since we know what we'll need.

Complete the getMessageCosts() function.

It takes a slice of messages and returns a slice of message costs.

Preallocate a slice for the message costs of the same length as the messages slice.
Fill the costs slice with costs for each message.
The cost in the cost slice should correspond to the message in the messages slice at the same index.
The cost of a message is the length of the message multiplied by 0.01.

// func make([]T, len, cap) []T
mySlice := make([]int, 5, 10)

	func main() {
		slice_test := []string{"Hello", "Hi", "Hey"}

		t := getMessageCosts(slice_test)
		fmt.Println(t)
		fmt.Println(len(t))
		fmt.Println(cap(t))
	}
*/
package main

func getMessageCosts(messages []string) []float64 {

	lentgh_slices_of_msg := len(messages)
	slice_with_costs := make([]float64, 0, lentgh_slices_of_msg) //Crear un slice variable_nombre:=make([]tipo,desde,hasta)

	for i := 0; i < lentgh_slices_of_msg; i++ {

		lenght_msg := len(messages[i])
		cost_msg := float64(lenght_msg) / 100
		slice_with_costs = append(slice_with_costs, cost_msg)
	}

	return slice_with_costs

}
