/*
Message Tagger

Textio needs a way to tag messages based on specific criteria.
Assignment

	    Complete the tagMessages function. It should take a slice of sms messages, and a
		function (that takes a sms as input and returns a slice of strings) as inputs. And it should return a slice of sms messages.

	    It should loop through each message and set the tags to the result of the passed in function.
	    Be sure to modify the messages of the original slice.
	    Ensure the tags field contains an initialized slice. No nil slices.

	    Complete the tagger function. It should take an sms message and return a slice of strings.

	    For any message that contains "urgent" (regardless of casing) in the content, the Urgent tag should be applied first.
	    For any message that contains "sale" (regardless of casing), the Promo tag should be applied second.

Note: regardless of casing just means that even "uRGent" or "SALE" should trigger the tag.

Example usage:

	messages := []sms{
		{id: "001", content: "Urgent! Last chance to see!"},
		{id: "002", content: "Big sale on all items!"},
		// Additional messages...
	}

taggedMessages := tagMessages(messages, tagger)
// `taggedMessages` will now have tags based on the content.
// 001 = [Urgent]
// 002 = [Promo]

# Tip

The go strings package, specifically the Contains and ToLower functions, can be very helpful here!
*/

package main

import "strings"

type sms struct {
	id      string
	content string
	tags    []string
}

func tagMessages(messages []sms, tagger func(sms) []string) []sms {

	for i, current_msg := range messages {

		messages[i].tags = tagger(current_msg)

	}

	return messages
}

func tagger(msg sms) []string {

	urg_tag := "Urgent"
	sale_tag := "Promo"

	urg_text_to_find := "urgent"
	sale_text_to_find := "sale"

	tags := []string{}

	is_urgent := strings.Contains(strings.ToLower(msg.content), urg_text_to_find)
	is_sale := strings.Contains(strings.ToLower(msg.content), sale_text_to_find)

	if is_urgent {
		tags = append(tags, urg_tag)

	}
	if is_sale {

		tags = append(tags, sale_tag)
	}
	return tags
}
