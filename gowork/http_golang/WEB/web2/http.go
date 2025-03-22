package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func getIssueData(url string) ([]byte, error) {
	res, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("error creating request: %v", err)
	}
	defer res.Body.Close()
	data, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, fmt.Errorf("error creating request: %w", err)
	}
	return data, nil
}

func prettify(data string) (string, error) {

	var prettyJson bytes.Buffer
	err := json.Indent(&prettyJson, []byte(data), "", " ")
	if err != nil {
		return "", fmt.Errorf("error indenting JSON: %w", err)
	}
	return prettyJson.String(), nil
}
