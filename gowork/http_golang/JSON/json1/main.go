/*
Assignment

Complete the issueList and userObject string constant values.

	issueList is a JSON array containing two issues (object literals). It should have the following issues:

ISSUE ONE:

	id: 0 (number)
	name: "Fix the thing" (string)
	estimate: 0.5 (number)
	completed: false (boolean)

ISSUE TWO:

	id: 1 (number)
	name: "Unstick the widget" (string)
	estimate: 30 (number)
	completed: false (boolean)

	  userObject is a JSON object literal that represents a Jello user. It should have the following fields:

name: "Wayne Lagner" (string)
role: "Developer" (string)
remote: true (boolean)

	You can use a backtick (`) to create a multi-line string in Go. This will make it easier to write the JSON data.

The tests simply check that the JSON is valid
*/
package main

const issueList = `
	{	
		"issues":[
			{
			"id":0,
			"name": "Fix the string",
			"estimate":0.5,
			"completed":false
			},
			{
			"id":1,
			"name": "Unstick the wigget",
			"estimate":30,
			"completed":false
			}
		]
	}

`
const userObject = `
		{
	 		"name": "Wayne Lagner",
			"role": "Developer",
			"remote": true
		}
	 
`

func main() {

}
