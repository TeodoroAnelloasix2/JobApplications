/*
Assignment

Complete the marshalAll function. It accepts a slice of "items", which can be of any type. The expectation is that they are structs of various forms.

    It should return a slice of slices of bytes (I didn't stutter),
	where each resulting slice of bytes represents the JSON representation of an item in the input slice.
    If an item cannot be marshaled, the function should immediately return an error.
    Return the marshalled data in the same order as the input items.
*/

package main

import (
	"encoding/json"
)

func marshalAll[T any](items []T) ([][]byte, error) {

	var marshaledjson [][]byte
	for _, j_son := range items {
		data, err := json.Marshal(j_son)
		if err != nil {
			return marshaledjson, err
		}

		marshaledjson = append(marshaledjson, data)
	}

	return marshaledjson, nil
}

func main() {

}
