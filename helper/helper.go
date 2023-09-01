package helper

import "strings"

func ValidateUserInput(firstName, lastName, email string, userTickets, remainingTickets uint8) (bool, bool, bool, bool) {
	// input validation
	//  name validation
	isValidName := len(firstName) >= 2 && len(lastName) >= 2
	// email validtion - mail contains @
	isValidEmail := strings.Contains(email, "@")
	// ticket validation: more than 0 less than or equal to remaining tickets
	isValidTicketNumber := userTickets > 0 && userTickets <= remainingTickets
	//  valid if all are true
	isValidInput := isValidName && isValidEmail && isValidTicketNumber

	return isValidInput, isValidEmail, isValidEmail, isValidTicketNumber

}