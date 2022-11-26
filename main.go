package main

import (
	"fmt"
)

func main() {
	const eventTickets int = 50
	eventName := "Go Event 22"
	var remainingTickets uint = 50
	// Slices Definition
	var bookings []string

	fmt.Println("Welcome to the", eventName, "application")
	// Print Formating
	fmt.Printf("Total tickets %v\n", eventTickets)
	fmt.Println("Get your tickets here to attend the event")

	for {
		var firstName string
		var lastName string
		var email string
		var userTickets uint

		fmt.Println("Enter you first name: ")
		fmt.Scan(&firstName)

		fmt.Println("Enter you lastname: ")
		fmt.Scan(&lastName)

		fmt.Println("Enter you email address: ")
		fmt.Scan(&email)

		fmt.Println("How many tickets required: ")
		fmt.Scan(&userTickets)

		remainingTickets = remainingTickets - userTickets

		fmt.Printf("Thank you %v %v for booking %v tickets.\nYou will receive booking details in you registered email %v\n", firstName, lastName, userTickets, email)
		fmt.Printf("\nTotal tickets remaining for the event %v\n", remainingTickets)
		fmt.Printf("\nTotal event bookings %v\n", bookings)
	}

}
