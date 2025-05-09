package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
)

const issueURL = "https://api.boot.dev/v1/courses_rest_api/learn-http/issues"

func main() {
	issues, err := getIssueData(issueURL)

	if err != nil {
		log.Fatalf("error getting issue data: %v", err)
	}
	prettyData, err := prettify(string(issues))
	if err != nil {
		log.Fatalf("error prettifying data: %v", err)
	}
	fmt.Println(prettyData)

}

func prettify(data string) (string, error) {

	var prettyJson bytes.Buffer
	err := json.Indent(&prettyJson, []byte(data), "", " ")
	if err != nil {
		return "", fmt.Errorf("error indenting JSON: %w", err)
	}
	return prettyJson.String(), nil
}
