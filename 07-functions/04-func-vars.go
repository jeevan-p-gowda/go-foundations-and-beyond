package main

import "fmt"

func main() {
	/*
		func(string)
		func(int, string)
		func(int, int) int
		func(int, int) (int, string)
		func(...int)int
	*/
	var sayHi func()
	sayHi = func() {
		fmt.Println("Hi!")
	}
	sayHi()

	var sayHello func(string)
	sayHello = func(userName string) {
		fmt.Printf("Hello %q!\n", userName)
	}
	sayHello("Jeevan")

	var greet func(string, string)
	greet = func(firstName, lastName string) {
		fmt.Printf("Hi %s %s, Have a nice day!\n", firstName, lastName)
	}
	greet("Jeevan", "Gowda")

	var getGreet func(string, string) string
	getGreet = func(firstName, lastName string) string {
		return fmt.Sprintf("Hi %s %s, Have a good day!", firstName, lastName)
	}
	msg := getGreet("Ravi", "Kiran")
	fmt.Println(msg)

	var divide func(int, int) (int, int)
	divide = func(x, y int) (quotient, remainder int) { //memory allocated & initialized for quotient & remainder
		quotient = x / y
		remainder = x % y
		return
	}
	q, r := divide(100, 7)
	fmt.Printf("Dividing 100 by 7, quotient = %d and remainder = %d\n", q, r)
}
