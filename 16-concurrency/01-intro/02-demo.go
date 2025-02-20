package main

import (
	"fmt"
	"time"
)

func main() {
	go f1() // scheduling the exection of "f1" through the scheduler
	f2()
	/*  the scheduler looks for other goroutines that scheduled and execute them when
	1. the current function execution is blocked for any reason (cooperative multitasking)
	2.  the current function is being busy for more than 20 ms (preemptive multitasking)
	*/

	// simulating (1.)
	// DO NOT use the following "poor man's synchronization" techniques
	time.Sleep(4 * time.Second)
	// fmt.Scanln()
}

func f1() {
	fmt.Println("f1 started")
	time.Sleep(3 * time.Second)
	fmt.Println("f1 completed")
}

func f2() {
	fmt.Println("f2 invoked")
}
