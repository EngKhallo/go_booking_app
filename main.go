package main

import (
	"fmt"
	"strings"
)

func main() {
	var conferenceName = "Go Conference"
	const conferenceTickets = 50
	var remainingTickets uint = 50
	var bookings []string

	greetUser(conferenceName, conferenceTickets, remainingTickets)

	for remainingTickets > 0 && len(bookings) < 50 { // conditions with loop

		userName, email, userTickets := getUserInputs()
		isValidName, isValidEmail, isValidTicketBooking := validateUserInputs(userName, email, userTickets, remainingTickets)

		if isValidName && isValidEmail && isValidTicketBooking {

			bookTickets(remainingTickets, userTickets, bookings, userName, email, conferenceName)

			if remainingTickets == 0 {
				fmt.Printf("availbele tickets are all sold out\n")
				break
			}
		} else {
			if !isValidName {
				fmt.Println("Invalid username")
			}
			if !isValidEmail {
				fmt.Println("Invalid email")
			}
			if !isValidTicketBooking {
				fmt.Println("Invalid booking ticket number")
			}

			fmt.Println("Your input is invalid, try again")
		}
	}

}

func greetUser(confName string, confTickets int, remainTickets uint) {
	fmt.Printf("Hello, welcome to %v booking application\n", confName)
	fmt.Printf("We've a total of %v tickets but %v tickets are still available\n", confTickets, remainTickets)
	fmt.Println("Get your tickets here to attend")
}

func validateUserInputs(userName string, email string, userTickets uint, remainingTickets uint) (bool, bool, bool) {
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

func bookTickets(remainingTickets uint, userTickets uint, bookings []string, userName, email string, conferenceName string) {
	remainingTickets = remainingTickets - userTickets
	bookings = append(bookings, userName)

	fmt.Printf("Thank you %v for booking %v tickets, you'll receive confirmation email at %v\n", userName, userTickets, email)
	fmt.Printf("%v tickets still remains for %v\n", remainingTickets, conferenceName)
	fmt.Printf("Booked names are %v\n", bookings)
}
