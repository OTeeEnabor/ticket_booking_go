package main

import (
	"fmt"
	"strings"
)
func main () {
	var conferenceName string = "Demystifying Blockchain Bootcap"
	const confTickets uint8 = 50
	var remainingTickets uint8 = 50
	// define slice
	var bookings []string // array of users who booked a ticket

	fmt.Printf("Welcome to the %v!\n", conferenceName)
	fmt.Println("Are you ready to start this amazing journey into blockchain?")
	fmt.Printf("There are %v tickets, and %v are still available.\n", confTickets, remainingTickets)

	for { // infinite loop syntaxt
		var firstName string
		var lastName string
		var email string 
		var userTickets uint8
		
		// getting user input
		fmt.Println("Enter your first name: ")
		fmt.Scan(&firstName)
	
		fmt.Println("Enter your last name: ")
		fmt.Scan(&lastName)
	
		fmt.Println("Enter your email address: ")
		fmt.Scan(&email)
	
		fmt.Println("Enter number of tickets to purchase: ")
		fmt.Scan(&userTickets)
	
	
		remainingTickets = remainingTickets - userTickets
		// bookings[0] = firstName + " " + lastName
		bookings = append(bookings, firstName + " " + lastName)
	
	
		fmt.Printf("Thank you %v %v for booking %v tickets. You will receive a confirmation email at %v\n", firstName, lastName, userTickets, email)
		fmt.Printf("%v tickets remaining for %v\n", remainingTickets, conferenceName)
		// declare a slice variable to store firstnames of bookings
		firstNames := []string{}

		for _, booking := range bookings {
			var names = strings.Fields(booking)
			firstNames = append(firstNames, names[0])

		}
		
		// print bookings
		fmt.Printf("Thes first names of bookings are: %v\n", firstNames)
	}



}
