package main

import "fmt"

func main() {
	var conferenceName = "Go Conference"
	const conferenceTickets = 50
	var remainingTickets uint = 50
	var bookings []string

	fmt.Printf("Welcome to %v booking application\n", conferenceName)
	fmt.Printf("We've a total of %v tickets but %v tickets are still available\n", conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets here to attend")

	for {

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

		remainingTickets = remainingTickets - userTickets
		bookings = append(bookings, userName)

		fmt.Printf("Thank you %v for booking %v tickets, you'll receive confirmation email at %v\n", userName, userTickets, email)
		fmt.Printf("%v tickets still remains for %v\n", remainingTickets, conferenceName)

		fmt.Printf("Booked names are %v\n", bookings)

		if remainingTickets == 0{
			fmt.Printf("availbele tickets are all sold out\n")
			break
		}
	}

}
