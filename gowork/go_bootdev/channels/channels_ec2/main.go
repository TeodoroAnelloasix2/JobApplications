/*
Process Messages

Textio needs to speed up message processing using concurrency.
Assignment

Implement the processMessages function, it takes a slice of messages.
It should process each message concurrently within a goroutine.
Use the process function for this to simulate the benefit of using goroutines.
Use a channel to ensure that all messages are processed and collected correctly, then return the slice of processed messages.

Example output:
messages := []string{"Here are some messages", "Here is another", "and another"}
processedMessages := processMessages(messages)
fmt.Println(processedMessages)
// prints ["Here are some messages-processed", "Here is another-processed", "and another-processed"]
*/
package main

import "time"

func processMessages(messages []string) []string {
	var processed []string

	proces_msg_ch := make(chan string, len(messages))

	for _, msg := range messages {

		go func() {

			p := process(msg)
			proces_msg_ch <- p

		}()

		processed = append(processed, <-proces_msg_ch)

	}
	return processed
}

// don't touch below this line

func process(message string) string {
	time.Sleep(1 * time.Second)
	return message + "-processed"
}
