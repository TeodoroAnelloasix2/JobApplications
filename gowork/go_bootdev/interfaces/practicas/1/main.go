/*
Assignment
Implement the Formatter interface with a method Format that returns a formatted string.
Define structs that satisfy the Formatter interface: PlainText, Bold, Code.
The structs must all have a message field of type string
PlainText should return the message as is.
Bold should wrap the message in two asterisks (**) to simulate bold text (e.g., message).
Code should wrap the message in a single backtick (`) to simulate code block (e.g., message)
*/
package main

import (
	"fmt"
)

func SendMessage(formatter Formatter) string {

	return formatter.Format() // Adjusted to call Format without an argument
}

type Formatter interface {
	Format() string
}

type PlainText struct {
	message string
}

func (p PlainText) Format() string {
	return p.message
}

type Bold struct {
	message string
}

func (b Bold) Format() string {
	return fmt.Sprintf("**%s**", b.message)
}

type Code struct {
	message string
}

func (c Code) Format() string {
	return fmt.Sprintf("`%s`", c.message)
}
