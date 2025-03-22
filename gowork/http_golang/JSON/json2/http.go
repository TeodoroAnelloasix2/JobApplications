/*
Assignment

In previous lessons, we've converted response into slices of bytes, and then strings.
Now, decode the response data directly into a slice of issues and return that instead.

	Import the json package from "encoding/json".
	Create a nil slice of items []Issue.
	Create a new *json.Decoder using json.NewDecoder.
	Decode the response body using the decoder's Decode method.
	Return the slice of issues.

If any error occurs return a nil slice and the error.
*/
package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

var items []Issue

func getIssues(url string) ([]Issue, error) {

	res, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("error creating request: %w", err)
	}
	defer res.Body.Close()

	decoder := json.NewDecoder(res.Body)

	if err := decoder.Decode(&items); err != nil {

		return nil, fmt.Errorf("Error during process: %w", err)

	}
	return items, nil

}
