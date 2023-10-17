package main

import (
	"booking-app/helper"
	"fmt"
	"time"
)

// package level variables: needs to be accessed from everywhere outside the Main
var conferenceName = "Go Conference"

const conferenceTickets = 50

var remainingTickets uint = 50

var bookings = make([]userData, 0) // initiates an empty list of userData struct

type userData struct {
	userName        string
	email           string
	numberOfTickets uint
}

func main() {

	greetUser()

	for remainingTickets > 0 && len(bookings) < 50 { // conditions with loop

		userName, email, userTickets := helper.GetUserInputs()
		isValidName, isValidEmail, isValidTicketBooking := helper.ValidateUserInputs(userName, email, userTickets, remainingTickets)

		if isValidName && isValidEmail && isValidTicketBooking {

			bookTickets(userTickets, userName, email)
			sendTicket(userTickets, userName, email)

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

func greetUser() {
	fmt.Printf("Hello, welcome to %v booking application\n", conferenceName)
	fmt.Printf("We've a total of %v tickets but %v tickets are still available\n", conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets here to attend")
}

func bookTickets(userTickets uint, userName string, email string) {
	remainingTickets = remainingTickets - userTickets

	var userData = userData{
		userName:        userName,
		email:           email,
		numberOfTickets: userTickets,
	}
	bookings = append(bookings, userData)

	fmt.Printf("Thank you %v for booking %v tickets, you'll receive confirmation email at %v\n", userName, userTickets, email)
	fmt.Printf("%v tickets still remains for %v\n", remainingTickets, conferenceName)
	fmt.Printf("List of booked people are %v\n", bookings)
}

func sendTicket(userTickets uint, userName string, email string) {
	time.Sleep(10 * time.Second)
	var ticket = fmt.Sprintf("%v tickets for %v", userTickets, userName)
	fmt.Println("########")
	fmt.Printf("Sending %v\n to email address: %v\n", ticket, email)
	fmt.Println("########")
}
