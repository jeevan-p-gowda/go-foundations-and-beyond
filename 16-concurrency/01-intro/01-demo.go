package main

import (
	"fmt"
)

func main() {
	go f1() // scheduling the exec of this function through the scheduler
	f2()

	// DO NOT DO THIS (poor man's synchronization techniques)
	// time.Sleep(1 * time.Second)
	fmt.Scanln()
}

func f1() {
	fmt.Println("f1 invoked")
}

func f2() {
	fmt.Println("f2 invoked")
}
