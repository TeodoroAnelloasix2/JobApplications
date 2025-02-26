package main

import (
	"fmt"
)

func sendSMSToCouple(msgToCustomer, msgToSpouse string) (int, error) {
	cost_total := 0

	cost, err := sendSMS(msgToCustomer)

	if err == nil {
		cost_total = cost_total + cost

		cost, err = sendSMS(msgToSpouse)

		if err == nil {
			cost_total = cost_total + cost
			return cost_total, err
		}
		return 0, err
	}
	return 0, err
}

func sendSMS(message string) (int, error) {
	const maxTextLen = 25
	const costPerChar = 2
	if len(message) > maxTextLen {
		return 0, fmt.Errorf("can't send texts over %v characters", maxTextLen)
	}
	return costPerChar * len(message), nil
}
