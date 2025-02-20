package main

import (
	"errors"
	"fmt"
	"math/rand"
	"time"
)

func main() {
	for range 20 {
		time.Sleep(500 * time.Millisecond)
		if err := doSomething(); err == nil {
			fmt.Println("Job done!")
		} else {
			fmt.Println("Error :", err)
		}
	}
}

func doSomething() error {
	if rand.Intn(200)%2 == 0 {
		return errors.New("something went wrong")
	}
	return nil
}
