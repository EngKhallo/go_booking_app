package main

import (
	"fmt"
	"strings"
)

func validateUserInputs(userName string, email string, userTickets uint) (bool, bool, bool) {
	isValidName := len(userName) >= 2
	isValidEmail := strings.Contains(email, "@")
	isValidTicketBooking := userTickets > 0 && userTickets <= remainingTickets

	return isValidName, isValidEmail, isValidTicketBooking
}

func getUserInputs() (string, string, uint) {
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
