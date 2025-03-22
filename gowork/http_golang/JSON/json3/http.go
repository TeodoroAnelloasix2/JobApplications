/*
Assignment

Update the getIssueData function in http.go.

	Change the return signature to return []Issue instead of []byte.
	Because the function will now return a decoded slice of issues, change the name from getIssueData to getIssues.
	Get the data from the response body using io.ReadAll, creating a slice of bytes []byte.
	Create a nil slice of issues []Issue.
	Use json.Unmarshal on the data to get the JSON data.
	Return the issues.
*/
package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

var item []Issue

func getIssues(url string) ([]Issue, error) {
	res, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("error creating request: %w", err)
	}
	defer res.Body.Close()

	data, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	if err = json.Unmarshal(data, &item); err != nil {
		return nil, err
	}

	return item, nil
}

func main() {

}
