/*
Assignment

Define an emailStatus type that uses iota syntax to represent the following states:

	emailBounced: 0
	emailInvalid: 1
	emailDelivered: 2
	emailOpened: 3

Go has a language feature, that when used with a type definition (and if you squint really hard), kinda looks like an enum (but it's not).
It's called iota.

type sendingChannel int

const (

	Email sendingChannel = iota
	SMS
	Phone

)
*/
package main

type emailStatus int

const (
	emailBounced = iota
	emailInvalid
	emailDelivered
	emailOpened
)
