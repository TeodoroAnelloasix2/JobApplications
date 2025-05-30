/*

Assignment

At Mailio we send a lot of network requests.
Each email we send must go out over the internet.
To serve our millions of customers, we need a single Go program to be capable of sending thousands of emails at once.

Edit the sendEmail() function to execute its anonymous function concurrently so that the "received" message prints after the "sent" message.

*/

package main

import (
	"fmt"
	"time"
)

func sendEmail(message string) {

	go func() { //Como que ahora recibed espera el tiempo estableccido en time.sleep y se ejecuta a la vez que email sent se imprime después!
		time.Sleep(time.Millisecond * 300)
		fmt.Printf("Email received: '%s'\n", message)
	}()
	fmt.Printf("Email sent: '%s'\n", message)
}

// Don't touch below this line

func test(message string) {
	sendEmail(message)
	time.Sleep(time.Millisecond * 500)
	fmt.Println("========================")
}

func main() {
	test("Hello there Kaladin!")
	test("Hi there Shallan!")
	test("Hey there Dalinar!")
}
