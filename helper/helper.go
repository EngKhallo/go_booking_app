package helper

import (
	"fmt"
	"strings"
)

func ValidateUserInputs(userName string, email string, userTickets uint, remainingTickets uint) (bool, bool, bool) {
	isValidName := len(userName) >= 2
	isValidEmail := strings.Contains(email, "@")
	isValidTicketBooking := userTickets > 0 && userTickets <= remainingTickets

	return isValidName, isValidEmail, isValidTicketBooking
}

func GetUserInputs() (string, string, uint) {
	var userName string
	var email string
	var userTickets uint
	// ask user for their name: user input
	fmt.Println("Please enter your name: ")
	fmt.Scan(&userName)

	fmt.Println("Please enter your email: ")
	fmt.Scan(&email)

	fmt.Println("enter number of tickets: ")
	fmt.Scan(&userTickets)

	return userName, email, userTickets
}
