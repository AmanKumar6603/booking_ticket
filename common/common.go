package common

import "strings"

// to validate input parameters
func ValidateInputParameters(firstname string, Lastname string, useremail string, userTickets uint, remainingTickets uint) (bool, bool, bool) {
	isValidName := len(firstname) >= 2 && len(Lastname) >= 2
	isValidEmail := strings.Contains(useremail, "@")
	isvalidticketNo := userTickets > 0 && userTickets <= remainingTickets
	return isValidName, isValidEmail, isvalidticketNo
}
