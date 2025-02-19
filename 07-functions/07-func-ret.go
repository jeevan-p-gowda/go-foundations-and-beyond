package main

import (
	"fmt"
	"math/rand"
)

func main() {
	fn := getFn()
	fn()
}

func getFn() func() {
	if r := rand.Intn(20); r%2 == 0 {
		return f1
	}
	return f2
}

func f1() {
	fmt.Println("f1 invoked")
}

func f2() {
	fmt.Println("f2 invoked")
}
