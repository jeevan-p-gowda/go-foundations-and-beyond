package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

var wg sync.WaitGroup

func main() {
	wg.Add(10)
	go f1()
	f2()
	wg.Wait() // block the execution of this function until the internal counter becomes 0 (default = 0)
}

func f1() {
	fmt.Println("f1 started")
	simulatedTime := rand.Intn(20)
	fmt.Printf("[f1] simulating a time consuming operation for %vs duration\n", simulatedTime)
	time.Sleep(time.Duration(simulatedTime) * time.Second)
	fmt.Println("f1 completed")
	wg.Done() // decrement the wg counter by 1
}

func f2() {
	fmt.Println("f2 invoked")
}
