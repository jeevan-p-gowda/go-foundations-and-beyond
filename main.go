package main

import (
	"fmt"
	"strings"
)

func main() {
	conferenceName := "Go Conference" //this is equal to var conferenceName = "Go Conference", := cannot be used for const
	const conferenceTickets = 50

	// int16 in GO is equal to int in Java,
	// uint shall be used only for +ve integer, int shall be used for both -ve and +ve integer
	// use the byte size wisely like int8, uint32
	var remainingTickets uint = 50

	//slices are abstraction of array, which adds dynamic attribute
	var bookings []string //alt: var bookings = []string{} or bookings := []string{}

	fmt.Printf("conference name is %T, conference tickets is %T, remaining tickets is %T\n", conferenceName, conferenceTickets, remainingTickets)

	fmt.Printf("Welcome to %v booking application\n", conferenceName)
	fmt.Printf("We have total of %v tickets and %v are still valid\n", conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets to attend")

	for {
		var firstName string
		var lastName string
		var email string
		var userTickets uint

		fmt.Println("Enter your first name: ")
		//& is a pointer, pointer is address which has the address of another varaible
		fmt.Scan(&firstName)

		fmt.Println("Enter your last name: ")
		fmt.Scan(&lastName)

		fmt.Println("Enter your email: ")
		fmt.Scan(&email)

		fmt.Println("Enter no. of tickets: ")
		fmt.Scan(&userTickets)

		bookings = append(bookings, firstName+" "+lastName)

		remainingTickets = remainingTickets - userTickets

		fmt.Printf("Thank you %v %v for booking %v tickets. You will receive a confirmation email at %v\n", firstName, lastName, userTickets, email)
		fmt.Printf("%v tickets remaining for %v\n", remainingTickets, conferenceName)

		firstNames := []string{}
		for _, booking := range bookings {
			var names = strings.Fields(booking)
			firstNames = append(firstNames, names[0])
		}

		fmt.Printf("The first name of the bookings are: %v\n", firstNames)

	}

}
