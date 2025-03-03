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

	greetUsers(conferenceName, conferenceTickets, remainingTickets)

	for remainingTickets > 0 && len(bookings) < 50 {
		firstName, lastName, email, userTickets := getUserInput()
		isNameValid, isEmailValid, isTicketsValid := validateUserInput(firstName, lastName, email, userTickets)

		if isNameValid && isEmailValid && isTicketsValid {
			bookTicket(remainingTickets, conferenceName, firstName, lastName, email, userTickets)
			bookings = append(bookings, firstName+" "+lastName)

			firstNames := getFirstNames(bookings)
			fmt.Printf("These are the first names %v\n", firstNames)
			// firstNames := []string{}
			// for _, booking := range bookings { // _ defines we are not using the entity which comes in that place, here its index
			// 	var names = strings.Fields(booking)
			// 	firstNames = append(firstNames, names[0])
			// }

			noRemainingTickets := remainingTickets == 0
			if noRemainingTickets {
				//end program
				fmt.Println("Our conference is booked")
				break
			}
		} else {
			fmt.Println("your input data is invalid try again")
			continue
		}

	}

}

func greetUsers(conferenceName string, conferenceTickets int, remainingTickets uint) {
	fmt.Printf("Welcome to %v booking application\n", conferenceName)
	fmt.Printf("We have total of %v tickets and %v are still valid\n", conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets to attend")
}

func getFirstNames(bookings []string) []string {
	firstNames := []string{}
	for _, booking := range bookings { // _ defines we are not using the entity which comes in that place, here its index
		var names = strings.Fields(booking)
		firstNames = append(firstNames, names[0])
	}
	return firstNames

}

func validateUserInput(firstName string, lastName string, email string, userTickets uint) (bool, bool, bool) {
	isNameValid := len(firstName) >= 2 && len(lastName) >= 2
	isEmailValid := strings.Contains(email, "@")
	isTicketsValid := userTickets > 0 && userTickets <= 50
	return isNameValid, isEmailValid, isTicketsValid
}

func getUserInput() (string, string, string, uint) {
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

	return firstName, lastName, email, userTickets
}

func bookTicket(remainingTickets uint, conferenceName, firstName, lastName, email string, userTickets uint) {

	remainingTickets = remainingTickets - userTickets

	fmt.Printf("Thank you %v %v for booking %v tickets. You will receive a confirmation email at %v\n", firstName, lastName, userTickets, email)
	fmt.Printf("%v tickets remaining for %v\n", remainingTickets, conferenceName)
}
