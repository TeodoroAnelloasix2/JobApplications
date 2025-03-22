/*
Assignment

There is a bug in the getIssueData function!
It's returning the entire http.Response instead of the data from the body (a slice of bytes). Fix it so that it returns []byte.

	Use io.ReadAll to read the .Body of the response.
	Return the resulting []byte
*/
package main

import (
	"fmt"
	"io"
	"net/http"
)

func getIssueData(url string) ([]byte, error) {

	res, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("error creating request: %w", err)
	}
	defer res.Body.Close()

	data, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, fmt.Errorf("error creating request: %w", err)
	}

	return data, nil
}
