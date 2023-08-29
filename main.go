package main

import (
	"fmt"
)
func main () {
	var conferenceName string = "Demystifying Blockchain Bootcap"
	const confTickets int8 = 50
	var remainingTickets = 50

	fmt.Printf("Welcome to the %v!\n", conferenceName)
	fmt.Println("Are you ready to start this amazing journey into blockchain?")
	fmt.Printf("There are %v tickets, and %v are still available", confTickets, remainingTickets)


}
