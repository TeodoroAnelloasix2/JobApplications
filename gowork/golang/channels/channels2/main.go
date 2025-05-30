/*
Assignment

Run the program.
You'll see that it deadlocks and never exits.
The sendIsOld function is trying to send on a channel, but no other goroutines are running that can accept the value from the channel.

Fix the deadlock by spawning a goroutine to send the "is old" values.

*/

package main

import (
	"time"
)

type email struct {
	body string
	date time.Time
}

func checkEmailAge(emails [3]email) [3]bool {
	isOldChan := make(chan bool)

	go sendIsOld(isOldChan, emails) // Creamos una go routine para que los canales puedan inter-comunicar entre ellos

	isOld := [3]bool{}
	isOld[0] = <-isOldChan
	isOld[1] = <-isOldChan
	isOld[2] = <-isOldChan
	return isOld
}

// don't touch below this line

func sendIsOld(isOldChan chan<- bool, emails [3]email) {
	defer close(isOldChan)
	for _, e := range emails {
		if e.date.Before(time.Date(2020, 0, 0, 0, 0, 0, 0, time.UTC)) {
			isOldChan <- true
			continue
		}
		isOldChan <- false
	}

}
