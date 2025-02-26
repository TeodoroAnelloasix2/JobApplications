/*
Mailio

Textio is launching a new email messaging product, "Mailio"!
Assignment

Fix the compile-time bug in the getFormattedMessages function.
The function body is correct, but the function signature is not.

package main

													the problem is here!
													   |
													   |
													  \ /
	func getFormattedMessages(messages []string, formatter func) []string {
		formattedMessages := []string{}
		for _, message := range messages {
			formattedMessages = append(formattedMessages, formatter(message))
		}
		return formattedMessages
	}
*/
package main

func getFormattedMessages(messages []string, formatter func(string) string) []string {
	formattedMessages := []string{}
	for _, message := range messages {
		formattedMessages = append(formattedMessages, formatter(message))
	}
	return formattedMessages
}
