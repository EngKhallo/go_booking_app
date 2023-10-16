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

	fmt.Printf("Welcome to %v booking application\n", conferenceName)
	fmt.Printf("We've a total of %v tickets but %v tickets are still available\n", conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets here to attend")

	for remainingTickets > 0 && len(bookings) < 50 { // conditions with loop
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

		isValidName := len(userName) >= 2
		isValidEmail := strings.Contains(email, "@")
		isValidTicketBooking := userTickets > 0 && userTickets <= remainingTickets

		if isValidName && isValidEmail && isValidTicketBooking {

			remainingTickets = remainingTickets - userTickets
			bookings = append(bookings, userName)

			fmt.Printf("Thank you %v for booking %v tickets, you'll receive confirmation email at %v\n", userName, userTickets, email)
			fmt.Printf("%v tickets still remains for %v\n", remainingTickets, conferenceName)

			fmt.Printf("Booked names are %v\n", bookings)

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
