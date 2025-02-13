package main

import "fmt"

func main() {
	conferenceName := "Go Conference" //this is equal to var conferenceName = "Go Conference", := cannot be used for const
	const conferenceTickets = 50

	// int16 in GO is equal to int in Java,
	// uint shall be used only for +ve integer, int shall be used for both -ve and +ve integer
	// use the byte size wisely like int8, uint32
	var remainingTickets uint = 50

	fmt.Printf("conference name is %T, conference tickets is %T, remaining tickets is %T\n", conferenceName, conferenceTickets, remainingTickets)

	fmt.Printf("Welcome to %v booking application\n", conferenceName)
	fmt.Printf("We have total of %v tickets and %v are still valid\n", conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets to attend")

	var userName string
	var userTickets int

	userName = "Tom"
	userTickets = 2

	fmt.Printf("User %v booked %v tickets.\n", userName, userTickets)
}
