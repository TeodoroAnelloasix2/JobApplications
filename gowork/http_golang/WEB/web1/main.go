/*
Jello
Throughout this course, we'll be building parts of an online issue tracking app called "Jello".
It's a typical issue tracker with one key difference: this time it's good, actually.

# Assignment

Take a look at the getIssueData function that I've provided in http.go. It retrieves information about issues from Jello's servers via HTTP as a slice of bytes []byte.

In main.go do the following:

	Convert the slice of bytes to a string with the built-in string() type conversion.
	Print the string representation of the bytes to the console.

It should look like a long, unformatted string of text representing raw issue data.
Tip

Notice how none of the data that is logged to the console was generated within our code!
That's because the data we retrieved is being sent over the internet from our servers via HTTP.
*/
package main

import (
	"fmt"
	"log"
)

func main() {
	issues, err := getIssueData()
	if err != nil {
		log.Fatalf("error getting issue data: %v", err)
	}
	fmt.Println(string(issues))

}
