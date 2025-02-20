package main

import (
	"fmt"
	"time"
)

// Share memory by communicating (using channels)

func main() {
	/*
		var ch chan int
		ch = make(chan int)
	*/
	ch := make(chan int)
	go add(100, 200, ch)
	result := <-ch
	fmt.Println("Result :", result)
}

func add(x, y int, ch chan int) {
	time.Sleep(2 * time.Second)
	result := x + y
	ch <- result
}
