package main

import (
	"fmt"
)

func main() {
	sayHi()
	sayHello("Jeevan")
	greet("Jevan", "Gowda")
	fmt.Println(getGreet("Ravi", "Kiran"))

	// fmt.Println(divide(100, 7))

	q, r := divide(100, 7)
	fmt.Printf("Dividing 100 by 7, quotient = %d and remainder = %d\n", q, r)

	// Using ONLY the quotient
	/*
		q, _ := divide(100, 7)
		fmt.Printf("Dividing 100 by 7, quotient = %d\n", q)
	*/

	// Using ONLY the remainder
	/*
		_, r := divide(100, 7)
		fmt.Printf("Dividing 100 by 7, remainder = %d\n", r)
	*/
}

func sayHi() {
	fmt.Println("Hi!")
}

func sayHello(userName string) {
	fmt.Printf("Hello %q!\n", userName)
}

/*
func greet(firstName string, lastName string) {
	fmt.Printf("Hi %s %s, Have a nice day!\n", firstName, lastName)
}
*/

func greet(firstName, lastName string) {
	fmt.Printf("Hi %s %s, Have a nice day!\n", firstName, lastName)
}

func getGreet(firstName, lastName string) string {
	return fmt.Sprintf("Hi %s %s, Have a good day!", firstName, lastName)
}

/*
func divide(x, y int) (int, int) {
	quotient := x / y
	remainder := x % y
	return quotient, remainder
}
*/

// Using NAMED results
/*
func divide(x, y int) (quotient int, remainder int) {
	quotient = x / y
	remainder = x % y
	return
}
*/

func divide(x, y int) (quotient, remainder int) { //memory allocated & initialized for quotient & remainder
	quotient = x / y
	remainder = x % y
	return
}
