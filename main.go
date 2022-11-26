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

	// Array Defination
	// var bookings = [50]string{}
	// Alternate Syntax for Array
	var bookingsOne [50]string

	// Adding elemets to Array
	bookingsOne[0] = "Sam Mathew"
	fmt.Println("Array Print Out: ", bookingsOne)

	// Slices Definition
	var bookingsTwo []string
	bookingsTwo = append(bookingsTwo, "Tharayathu Sea Palace")

	fmt.Println("Slice Print Out: ", bookingsTwo)

	// fmt.Println("Memory Location of the variable eventName is:", &eventName)
	// fmt.Println(eventName)

	fmt.Println("Welcome to the", eventName, "application")
	// Print Formating
	fmt.Printf("Total tickets %v\n", eventTickets)
	fmt.Println("Get your tickets here to attend the event")

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
	// fmt.Printf("firstName type: %T and type of userTickets: %T\n", firstName, userTickets)

}
