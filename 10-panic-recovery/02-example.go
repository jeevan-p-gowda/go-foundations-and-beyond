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
	for {
		var divisor int
		fmt.Println("Enter the divisor")
		fmt.Scanln(&divisor)
		q, r, err := divideAdapter(100, divisor)
		if err == ErrDivideByZero {
			fmt.Println("Do not attempt to divide by zero. Try again!")
			continue
		}
		if err != nil {
			fmt.Println("Unknown error :", err)
			break
		}
		fmt.Printf("[main] Dividing 100 by %d, quotient = %d and remainder = %d\n", divisor, q, r)
		break
	}
}

// adapter function to covert the panic into an error
func divideAdapter(x, y int) (quotient, remainder int, err error) {
	defer func() {
		if e := recover(); e != nil {
			err = e.(error)
		}
		return
	}()
	quotient, remainder = divide(x, y)
	return
}

// 3rd party OS api
func divide(x, y int) (quotient, remainder int) {
	if y == 0 {
		panic(ErrDivideByZero)
	}
	quotient = x / y
	remainder = x % y
	return
}
