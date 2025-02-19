package main

import "fmt"

func main() {
	// var x interface{}
	var x any
	x = 100
	x = "Veniam sint minim magna proident adipisicing magna incididunt ex pariatur non commodo."
	x = true
	x = 99.99
	fmt.Println(x)

	x = 200
	// x = "Non commodo est id sint Lorem labore ea."
	// y := x.(int) * 2
	// type assertion (using if-else)
	if val, ok := x.(int); ok {
		y := val * 2
		fmt.Println(y)
	} else {
		fmt.Println("x is not an int")
	}

	// type asserting (using type-switch)
	// x = 100
	// x = "Nulla aliquip excepteur nulla aliquip consequat velit cupidatat aute."
	// x = true
	x = 99.99
	switch val := x.(type) {
	case int:
		fmt.Println("x is an int, x * 2 =", val*2)
	case string:
		fmt.Println("x is a string, len(x) =", len(val))
	case bool:
		fmt.Println("x is a bool, !x =", !val)
	default:
		fmt.Println("Unknown type")
	}

}
