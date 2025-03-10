package main

import "fmt"

func main() {
	exec(f1) // f1
	exec(f2) // f2
	exec(func() {
		fmt.Println("anonymous func invoked")
	})
}

func exec(fn func()) {
	fn()
}

func f1() {
	fmt.Println("f1 invoked")
}

func f2() {
	fmt.Println("f2 invoked")
}
