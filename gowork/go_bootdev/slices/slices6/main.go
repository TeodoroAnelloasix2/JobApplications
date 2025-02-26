/*
Assignment
We need to be able to quickly detect bad words in the messages our system sends.
Complete the indexOfFirstBadWord function.
If it finds any bad words in the message it should return the index of the first bad word in the msg slice.
This will help us filter out naughty words from our messaging system. If no bad words are found, return -1 instead.
Use the range keyword.

Tip
Remember, you can ignore returned values with an underscore _ instead of creating an unused variable.

fruits := []string{"apple", "banana", "grape"}

	for i, fruit := range fruits {
	    fmt.Println(i, fruit)
	}

0 apple
1 banana
2 grape
*/
package main

func indexOfFirstBadWord(msg []string, badWords []string) int {

	for i := range msg {
		for a := range badWords {
			if msg[i] == badWords[a] {
				return i
			}
		}
	}

	return -1
}
