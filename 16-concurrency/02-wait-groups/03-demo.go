package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	wg := &sync.WaitGroup{}
	for range 20 {
		wg.Add(1) //increment the counter by 1
		go f1(wg) // scheduling the exec of this function through the scheduler
	}
	f2()

	wg.Wait() // block the execution of this function until the wg counter becomes 0 (default = 0)

}

func f1(wg *sync.WaitGroup) {
	defer wg.Done() //decrement the counter by 1
	fmt.Println("f1 started")
	time.Sleep(5 * time.Second)
	fmt.Println("f1 completed")
}

func f2() {
	fmt.Println("f2 invoked")
}
