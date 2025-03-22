/*

Assignment

I've updated the getIssueData() function to be a bit more flexible.
It now takes a URL as a parameter.

Try running the code in its current state. You should notice an error because the URL we're using is invalid.

Fix the code so that the call to getIssueData function uses the provided issueURL.

*/

package main

import (
	"fmt"
	"log"
)

const issueURL = "https://api.boot.dev/v1/courses_rest_api/learn-http/issues"

func main() {

	issues, err := getIssueData(issueURL)
	if err != nil {
		log.Fatalf("Error getting issue data: %v", err)
	}
	prettyData, err := prettify(string(issues))

	if err != nil {
		log.Fatalf("error prettifying data: %v", err)

	}
	fmt.Println(prettyData)
}
