/*
Assignment

We want to be able to send emails in batches.
A writing goroutine will write an entire batch of email messages to a buffered channel, and later,
once the channel is full, a reading goroutine will read all of the messages from the channel and send them out to our clients.

Complete the addEmailsToQueue function.
It should create a buffered channel with a buffer large enough to store all of the emails it's given.
It should then write the emails to the channel in order, and finally return the channel.
*/

package main

func addEmailsToQueue(emails []string) chan string {

	email_channel := make(chan string, len(emails))

	for _, email := range emails {
		email_channel <- email
	}
	return email_channel
}
