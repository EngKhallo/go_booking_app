package main

import "fmt"

func main() {
	var conferenceName = "Go Conference"
	const conferenceTickets = 50
	const remainingTickets = 50

	fmt.Printf("Welcome to %v booking application\n", conferenceName)
	fmt.Printf("We've a total of %v tickets but %v tickets are still available\n", conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets here to attend")

	var userName string
	var email string
	var userTickets int
	// ask user for their name: user input
	fmt.Println("Please enter your name: ")
	fmt.Scan(&userName)

	fmt.Println("Please enter your email: ")
	fmt.Scan(&email)
	
	fmt.Println("enter number of tickets: ")
	fmt.Scan(&userTickets)

	fmt.Printf("Thank you %v for booking %v tickets, you'll receive confirmation email at %v", userName, userTickets, email)
}
