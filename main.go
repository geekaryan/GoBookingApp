package main

import (
	"fmt"
)

func main() {

	conferenceName := "Go Conference"
	const conferenceTickets int = 50
	var remainingTickets int = 50

	fmt.Printf("ConferenceTickets is %T, remainingTickets is %T, conferenceName is %T. \n", conferenceTickets, remainingTickets, conferenceName)

	fmt.Printf("Welcome to our %v booking applicaton.\n", conferenceName)
	fmt.Printf("We have total of %v tickets and out of which %v are remaining. \n", conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets here to attend")

	var firstName string
	var lastName string
	var email string
	var userTickets int
	//ask user for their name
	fmt.Println("Enter your first name: ")
	fmt.Scan(&firstName)

	fmt.Println("Enter your last name: ")
	fmt.Scan(&lastName)

	fmt.Println("Enter your email address: ")
	fmt.Scan(&email)

	fmt.Println("Enter number of tickets: ")
	fmt.Scan(&userTickets)

	remainingTickets = remainingTickets - userTickets

	var bookings []string
	bookings = append(bookings, firstName+" "+lastName)

	// fmt.Printf("The whole Slice: %v\n", bookings)
	// fmt.Printf("The first value: %v\n", bookings[0])
	// fmt.Printf("Slice type: %T\n", bookings) //captial T is used to print the type of the array
	// fmt.Printf("Slice length: %v\n", len(bookings))

	fmt.Printf("Thank you %v %v for booking %v tickets. You will recieve a confirmation email at %v.\n", firstName, lastName, userTickets, email)
	fmt.Printf("The remaing tickets left are %v", remainingTickets)

	fmt.Printf("These are all the bookings: %v\n", bookings)

}
