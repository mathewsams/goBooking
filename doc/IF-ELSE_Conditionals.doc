package main

import (
	"fmt"
	"strings"
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

	// Conditionals in for loop
	// for remainingTickets > 0 {}
	// for len(bookings) == 50 {}

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

		if userTickets <= remainingTickets {
			remainingTickets = remainingTickets - userTickets
			bookings = append(bookings, firstName+" "+lastName)

			firstNames := []string{}
			for _, booking := range bookings {
				var names = strings.Fields(booking)
				firstNames = append(firstNames, names[0])
			}
			fmt.Printf("Thank you %v %v for booking %v tickets.\nYou will receive booking details in you registered email %v\n", firstName, lastName, userTickets, email)
			fmt.Printf("\nTotal tickets remaining for the event %v\n", remainingTickets)
			// fmt.Printf("\nTotal event bookings %v\n", bookings)
			fmt.Printf("\nFirst Names of users with bookings %v\n", firstNames)

			// Boolean variable Declaration
			// var ticketsAvailable bool = remainingTickets == 0
			// Alternate Syntax
			ticketsAvailable := remainingTickets == 0
			if ticketsAvailable {
				fmt.Println("Event is booked Out. Come back next year")
				break
			}
		} else {
			fmt.Printf("We only have %v tickets remaining. Try again.\n", remainingTickets)
			// continue  --> No need of the continue as this is the last line of the code
		}
	}

}
