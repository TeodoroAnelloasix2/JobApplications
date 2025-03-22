/*
Assignment

A lazy Go programmer wrote the handleEmailBounce function...
just because the compiler doesn't force us to check errors doesn't mean we shouldn't!

Take a look at the updateStatus and track methods in the user.go file.
Handle their errors properly, and use the fmt.Errorf function and the %w formatting verb to add useful context to the errors.

    If updateStatus fails, return an error saying "error updating user status: ERR"
    If track fails, return an error saying "error tracking user bounce: ERR"

Where ERR is the error returned by the method.
*/

package main

import "fmt"

type email struct {
	status    string
	recipient *user
}

type user struct {
	email  string
	status string
}

type analytics struct {
	totalBounces int
}

func (u *user) updateStatus(status string) error {
	var err error
	if status != "email_bounced" {

		err = fmt.Errorf("error updating user status: invalid status: %v", status)
	}
	u.status = status
	return err
}

func (a *analytics) track(event string) error {
	var err error
	if event != "email_bounced" {
		err = fmt.Errorf("error tracking user bounce: %v", event)
	}
	a.totalBounces++
	return err
}
