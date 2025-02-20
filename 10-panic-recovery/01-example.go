package main

import (
	"errors"
	"fmt"
)

var ErrDivideByZero error = errors.New("divide by zero error")

func main() {
	defer func() {
		fmt.Println("	[main] - deferred")
		if err := recover(); err != nil {
			fmt.Println("Unknown error (panic) :", err)
			return
		}
		fmt.Println("Thank you!")
	}()
	var divisor int
	fmt.Println("Enter the divisor")
	fmt.Scanln(&divisor)
	q, r := divide(100, divisor)
	fmt.Printf("[main] Dividing 100 by %d, quotient = %d and remainder = %d\n", divisor, q, r)
}

func divide(x, y int) (quotient, remainder int) { //memory allocated & initialized for quotient & remainder
	defer func() {
		fmt.Println("	[divide] - deferred")
	}()
	fmt.Println("[divide] - calculating quotient")
	if y == 0 {
		panic(ErrDivideByZero)
	}
	quotient = x / y
	fmt.Println("[divide] - calculating remainder")
	remainder = x % y
	return
}
