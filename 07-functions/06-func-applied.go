package main

import (
	"fmt"
	"log"
)

func main() {
	// v1.0
	/*
		add(100, 200)
		subtract(100, 200)
	*/

	// v2.0 (log a message before and after the operations)
	/*
		log.Println("Operation started...")
		add(100, 200)
		log.Println("Operation completed!")

		log.Println("Operation started...")
		subtract(100, 200)
		log.Println("Operation completed!")
	*/
	logOperation(add, 100, 200)
	logOperation(subtract, 100, 200)
	logOperation(func(i1, i2 int) {
		fmt.Println("Multiply Result :", i1*i2)
	}, 100, 200)

	add := getLogOperation(add)
	subtract := getLogOperation(subtract)

	add(100, 200)
	subtract(100, 200)

}

func logOperation(oper func(int, int), x, y int) {
	log.Println("Operation started...")
	oper(x, y)
	log.Println("Operation completed!")
}

func getLogOperation(oper func(int, int)) func(int, int) {
	return func(x, y int) {
		log.Println("Operation started...")
		oper(x, y)
		log.Println("Operation completed!")
	}
}

func add(x, y int) {
	fmt.Println("Add Result :", x+y)
}

func subtract(x, y int) {
	fmt.Println("Subtract Result :", x-y)
}
