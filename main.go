package main

import (
	"fmt"
	"ticket-booking/helper"
	"time"
	"sync"
)

//  define package level variables
const conferenceTickets uint8 = 50
var conferenceName string = "Demystifying Blockchain Bootcamp"
var remainingTickets uint8 = 50
// define slice
var bookings = make([]UserData,0) // create a list of maps
var wg = sync.WaitGroup{}

type UserData struct {
	firstName string
	lastName string
	email string
	numberOfTickets uint
}

func main () {
	greetUsers()
	
	for { // infinite loop syntax

		// get user input
		firstName, lastName, email, userTickets := getUserInput()
		
		// validate user input
		isValidInput, isValidName, isValidEmail, isValidTicketNumber := helper.ValidateUserInput(firstName, lastName, email, userTickets, remainingTickets)
		

		if isValidInput {
			// book tickets
			bookTickets(userTickets, firstName, lastName, email)
			wg.Add(1)
			// send ticket
			go sendTicket(userTickets, firstName, lastName, email)
			
			
			// declare a slice variable to store firstnames of bookings
			firstNames := getFirstNames()
			
			// print first name bookings
			fmt.Printf("The first names of bookings are: %v\n", firstNames)

			var noTicketsRemaining bool = remainingTickets == 0 //bit of overkill because this check is only used once in the application

			if noTicketsRemaining {
				// end the application
				fmt.Println("There are no tickets available!")
				// end the for loop
				break
			}
			
		}else {
			if !isValidName {
				fmt.Printf("%v is an invalid name! \n", firstName)
			}
			if !isValidEmail{
				fmt.Printf("%v is an invalid email, it doesnt contain an @ sign!\n", email)
			}
			if !isValidTicketNumber{
				fmt.Printf("%v is an invalid number of tickets! \n ", userTickets)
			}
			
		}
	}
	wg.Wait()


}

func greetUsers () {
	fmt.Printf("Welcome to the %v!\n", conferenceName)

	fmt.Println("")

	fmt.Println("Are you ready to start this amazing journey into blockchain?")

	fmt.Println("")
	fmt.Printf("There are %v tickets, and %v are still available.\n", conferenceTickets, remainingTickets)

	fmt.Println("")

}

func getFirstNames() []string {
	var firstNameList = []string{}
	for _,booking := range bookings{
		firstNameList = append(firstNameList, booking.firstName)
	}
	return firstNameList
}



func getUserInput () (string, string, string,uint8){
	var firstName string
		var lastName string
		var email string 
		var userTickets uint8
		
		// getting user input
		fmt.Println("Enter your first name: ")
		fmt.Scan(&firstName)
		
		// fmt.Println("")
		
		fmt.Println("Enter your last name: ")
		fmt.Scan(&lastName)
	
		fmt.Println("Enter your email address: ")
		fmt.Scan(&email)
	
		fmt.Println("Enter number of tickets to purchase: ")
		fmt.Scan(&userTickets)

		return firstName, lastName, email, userTickets
}

func bookTickets(userTickets uint8, firstName, lastName, email string)  {

	remainingTickets = remainingTickets - userTickets
	//  create a ma for a user 
	var userData = UserData{
		firstName: firstName,
		lastName: lastName,
		email: email,
		numberOfTickets: uint(userTickets),
	}
	

	// bookings list of userData maps
	bookings = append(bookings, userData)
	// print list of bookings
	fmt.Printf("List of bookings is %v\n", bookings)


	fmt.Printf("Thank you %v %v for booking %v tickets. You will receive a confirmation email at %v\n", firstName, lastName, userTickets, email)
	fmt.Printf("%v tickets remaining for %v\n", remainingTickets, conferenceName)

}

func sendTicket(userTickets uint8, firstName,lastName, email string) {
	time.Sleep(10 * time.Second)
	var ticket = fmt.Sprintf("%v tickets for %v %v", userTickets, firstName, lastName)
	fmt.Println("###############")
	fmt.Printf("Sending ticket:\n %v to email address %v\n", ticket, email)
	fmt.Println("###############")
	wg.Done()
}