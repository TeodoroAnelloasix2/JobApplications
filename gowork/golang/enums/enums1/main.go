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

//import "fmt"

func (a *analytics) handleEmailBounce(em email) error {
	var err error
	err = em.recipient.updateStatus(em.status)
	if err != nil {
		return err
	}
	err = a.track(em.status)
	if err != nil {
		return err
	}
	return err
}
