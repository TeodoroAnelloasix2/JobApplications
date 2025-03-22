/*
Assignment

A junior developer's deadline is looming.
They've been tasked with creating functions to retrieve and log various resources from Jello's API.
Rather than making multiple structs and type-safe functions, the junior developer has decided to make two flexible functions.
Help them complete their task before their boss, Wayne Lagner, asks them to refactor it.

Complete the getResources and logResources functions.
getResources()

getResources takes a url string and returns a slice of maps []map[string]interface{} and an error.

	It's a slice of maps because the API response is an array of JSON objects.

	Decode the response body into a slice of maps []map[string]interface{} and return it.

logResources()

logResources takes a slice of maps []map[string]interface{} and prints its keys and values to the console.
Because maps are unsorted we will be adding formatted strings to a slice of strings []string which is then* sorted.

	    Iterate over the slice of map[string]interface{}
	    For each map[string]interface{} get its keys and values using range and append it to formattedStrings as Key: %s - Value: %v,

		where %s is the key and %v is the value.

You do not need to account for nested JSON objects.
Assume they are neat key-value pairs, one level deep.

The output from one issue should look like this:

Issue:
Key: estimate - Value: 19
Key: id - Value: 0194fdc2-fa2f-4cc0-81d3-ff12045b73c8
Key: status - Value: TODO
Key: title - Value: Fix that one bug nobody understands
*/
package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"sort"
)

func getResources(url string) ([]map[string]any, error) {

	var resources []map[string]any

	res, err := http.Get(url)
	if err != nil {
		return resources, nil
	}

	defer res.Body.Close()

	data, err := io.ReadAll(res.Body)
	if err != nil {
		return resources, err
	}

	if err = json.Unmarshal(data, &resources); err != nil {
		return resources, err
	}

	return resources, nil
}

func logResources(resources []map[string]any) {

	var formattedStrings []string

	for _, res := range resources {

		for key, value := range res {
			txt := fmt.Sprintf("Key: %s - Value: %v", key, value)
			formattedStrings = append(formattedStrings, txt)
		}

	}

	sort.Strings(formattedStrings)

	for _, str := range formattedStrings {
		fmt.Println(str)
	}
}
