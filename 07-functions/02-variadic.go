package main

import "fmt"

func main() {
	fmt.Println(sum(10))
	fmt.Println(sum(10, 20))
	fmt.Println(sum(10, 20, 30, 40))
	fmt.Println(sum())
}

func sum(nos ...int) int {
	fmt.Println("[sum] # of operands :", len(nos))
	var result int
	for i := 0; i < len(nos); i++ {
		result += nos[i]
	}
	return result
}
