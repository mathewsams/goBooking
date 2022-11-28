package main

import (
	"fmt"
	"strconv"
	"strings"
)

const eventTickets int = 50

var eventName = "Go Event 22"
var remainingTickets uint = 50

// Slices Definition
// var bookings []string

// Creating an empty List of Maps
var bookings = make([]map[string]string, 0) // --> Initial Size of list

func main() {

	greetUser()

	for {
		firstName, lastName, email, userTickets := getUserInput()
		isValidName, isValidEmail, isValidTicketCount := validateUserInput(firstName, lastName, email, userTickets)

		if isValidName && isValidEmail && isValidTicketCount {
			bookTickets(firstName, lastName, email, userTickets)
			// fmt.Printf("\nTotal event bookings %v\n", bookings)

			firstNames := getFirstNames()
			fmt.Printf("\nFirst Names of users with bookings %v\n", firstNames)

			ticketsAvailable := remainingTickets == 0
			if ticketsAvailable {
				fmt.Println("Event is booked Out. Come back next year")
				break
			}
		} else {
			// fmt.Printf("We only have %v tickets remaining. Try again.\n", remainingTickets)
			fmt.Println("Input values are invalid")
			if !isValidName {
				fmt.Println("First and last name must be more than 2 charecter long")
			}
			if !isValidEmail {
				fmt.Println("Email address must contain @ sign")
			}
			if !isValidTicketCount {
				fmt.Println("Ticket count must be greater than 0 and less than remaining")
			}

			// If valdation must stop at the first found error then else-if statments can be added
		}
	}
}

func greetUser() {
	fmt.Println("Welcome to the", eventName, "application")
	// Print Formating
	fmt.Printf("Total tickets %v\n", eventTickets)
	fmt.Printf("Remaining tickets %v\n", remainingTickets)
	fmt.Println("Get your tickets here to attend the event")
}

func getFirstNames() []string {
	firstNames := []string{}
	for _, booking := range bookings {
		// var names = strings.Fields(booking)
		// firstNames = append(firstNames, names[0])

		// Accessing a MAP value using the key
		firstNames = append(firstNames, booking["firstName"])

	}
	return firstNames
}

func getUserInput() (string, string, string, uint) {
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

	return firstName, lastName, email, userTickets
}

func bookTickets(firstName string, lastName string, email string, userTickets uint) {
	remainingTickets = remainingTickets - userTickets

	// MAP Declaration and Assignment
	var userData = make(map[string]string)
	userData["firstName"] = firstName
	userData["lastName"] = lastName
	userData["email"] = email
	userData["userTickets"] = strconv.FormatUint(uint64(userTickets), 10)

	// bookings = append(bookings, firstName+" "+lastName)
	bookings = append(bookings, userData) // append map to the list

	fmt.Printf("List of Bookings %v\n", bookings)
	fmt.Printf("Thank you %v %v for booking %v tickets.\nYou will receive booking details in you registered email %v\n", firstName, lastName, userTickets, email)
	fmt.Printf("\nTotal tickets remaining for the event %v\n", remainingTickets)

	// return remainingTickets, bookings
}

func validateUserInput(firstName string, lastName string, email string, userTickets uint) (bool, bool, bool) {
	isValidName := len(firstName) >= 2 && len(lastName) >= 2
	isValidEmail := strings.Contains(email, "@")
	isValidTicketCount := userTickets > 0 && userTickets <= remainingTickets

	return isValidName, isValidEmail, isValidTicketCount
}
