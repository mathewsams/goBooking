package main

import (
	"fmt"
)

func main() {
	const eventTickets int = 50
	// var eventName string = "Go Event 22"
	// Alternate Syntax to declare variables - Syntatic Sugar :)
	eventName := "Go Event 22"
	var remainingTickets uint = 50

	// fmt.Println(eventName)

	fmt.Println("Welcome to the", eventName, "application")
	fmt.Printf("Total tickets %v\n", eventTickets)
	fmt.Println("Get your tickets here to attend the event")

	// Print Formating
	fmt.Printf("Total tickets Remaining are %v\n", remainingTickets)

	var userName string
	var userTickets int

	userName = "Sam"
	userTickets = 20

	fmt.Printf("User %v booked %v tickets\n", userName, userTickets)
	fmt.Printf("UserName type: %T and type of userTickets: %T\n", userName, userTickets)

}
