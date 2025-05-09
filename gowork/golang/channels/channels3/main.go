/*

Assignment

Our Mailio server isn't able to boot up until it receives the signal that its databases are all online,
and it learns about them being online by waiting for tokens (empty structs) on a channel.

Run the code.
It never exits! The channel passed to waitForDBs stays blocked, because it's only popping the first value off the channel.

Fix the waitForDBs function.
It should pause execution until it receives a token for every database from the dbChan channel.
Each time waitForDBs reads a token, the getDBsChannel goroutine will print a message to the console for you.
The succinctly named numDBs input is the total number of databases.
Look at the test code to see how these functions are used so you can understand the control flow.

*/

package main

import "fmt"

func waitForDBs(numDBs int, dbChan chan struct{}) {

	for i := 0; i < numDBs; i++ {
		<-dbChan
	}

}

// don't touch below this line

func getDBsChannel(numDBs int) (chan struct{}, *int) {
	count := 0
	ch := make(chan struct{})

	go func() {
		for i := 0; i < numDBs; i++ {
			ch <- struct{}{}
			fmt.Printf("Database %v is online\n", i+1)
			count++
		}
	}()

	return ch, &count
}
