package main

import "fmt"

func main() {
	var conferenceName = "Go Conference"
	const conferenceTickets = 50
	const remainingTickets = 50

	fmt.Println("Welcome to", conferenceName, "booking application")
	fmt.Println("We've a total of:", conferenceTickets, "tickets but", remainingTickets, "tickets are still available")
	fmt.Println("Get your tickets here to attend")

}
