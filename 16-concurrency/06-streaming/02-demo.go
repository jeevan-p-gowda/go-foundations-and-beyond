package main

import (
	"fmt"
	"time"
)

// consumer
func main() {
	ch := make(chan int)
	go genNos(ch)
	for range 5 {
		fmt.Println(<-ch)
	}
	fmt.Println("Done")
}

// producer
func genNos(ch chan int) {
	for i := range 5 {
		ch <- (i + 1) * 10
		time.Sleep(1 * time.Second)
	}

}
