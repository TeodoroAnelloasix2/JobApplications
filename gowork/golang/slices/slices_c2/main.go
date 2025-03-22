/*
Password Strength
As part of improving security, Textio wants to enforce a new password policy. A valid password must meet the following criteria:

At least 5 characters long but no more than 12 characters.
Contains at least one uppercase letter.
Contains at least one digit.

A string is really just a read-only slice of bytes. This means that you can use the same techniques you learned in the previous
lesson to iterate over the characters of a string.
Assignment
Implement the isValidPassword function by looping through each character in the password string.
Make sure the password is long enough and includes at least one uppercase letter and one digit.
Assume passwords consist of ASCII characters only.
*/
package main

import (
	"regexp"
)

func isValidPassword(password string) bool {
	is_ok := false
	length_passwrd := len(password)
	digit_match := match_digit(password)
	uppercase_match := match_uppercase(password)

	if length_passwrd >= 5 && length_passwrd <= 12 {
		if digit_match && uppercase_match {
			is_ok = true
			return is_ok
		}
		return is_ok
	}
	return is_ok

}

func match_uppercase(text string) bool {
	pattern := "[A-Z]"
	matched, _ := regexp.MatchString(pattern, text)
	return matched
}

func match_digit(text string) bool {

	pattern := "[0-9]"

	matched, _ := regexp.MatchString(pattern, text)

	return matched

}
