package main

import (
	"fmt"
	"math/rand"
	"time"
)

// consumer
func main() {
	ch := make(chan int)
	go genNos(ch)
	for data := range ch {
		fmt.Println(data)
	}
	fmt.Println("Done")
}

// producer
func genNos(ch chan int) {
	for i := range rand.Intn(20) {
		ch <- (i + 1) * 10
		time.Sleep(500 * time.Millisecond)
	}
	close(ch)
}
