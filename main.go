package main

import "fmt"

func main() {
	conferenceName := "Go Conference" //this is equal to var conferenceName = "Go Conference", := cannot be used for const
	const conferenceTickets = 50

	// int16 in GO is equal to int in Java,
	// uint shall be used only for +ve integer, int shall be used for both -ve and +ve integer
	// use the byte size wisely like int8, uint32
	var remainingTickets uint = 50
	var bookings [50]string

	fmt.Printf("conference name is %T, conference tickets is %T, remaining tickets is %T\n", conferenceName, conferenceTickets, remainingTickets)

	fmt.Printf("Welcome to %v booking application\n", conferenceName)
	fmt.Printf("We have total of %v tickets and %v are still valid\n", conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets to attend")

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

	bookings[0] = firstName + " " + lastName
	fmt.Printf("The whole array: %v\n", bookings)
	fmt.Printf("The first value: %v\n", bookings[0])
	fmt.Printf("Array type: %T\n", bookings)
	fmt.Printf("Array length: %v\n", len(bookings))

	remainingTickets = remainingTickets - userTickets

	fmt.Printf("Thank you %v %v for booking %v tickets. You will receive a confirmation email at %v\n", firstName, lastName, userTickets, email)
	fmt.Printf("%v tickets remaining for %v\n", remainingTickets, conferenceName)
}
