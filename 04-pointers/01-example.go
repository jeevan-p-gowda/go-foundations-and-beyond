package main

import "fmt"

func main() {
	var x int
	x = 100

	var xPtr *int
	xPtr = &x // value => address
	fmt.Println("[main], &x =", xPtr)

	// dereferencing : address => value
	fmt.Println(*xPtr)

	// in other words
	fmt.Println(x == *(&x))

	fmt.Println("Before incrementing, x =", x)
	increment(&x)
	fmt.Println("After incrementing, x =", x)

	n1, n2 := 100, 200
	fmt.Printf("Before swapping, n1 = %d & n2 = %d\n", n1, n2)
	swap(&n1, &n2)
	fmt.Printf("After swapping, n1 = %d & n2 = %d\n", n1, n2)

}

func increment(no *int) {
	fmt.Println("[increment], &no =", no)
	*no = *no + 1
}

func swap(x *int, y *int) {
	*x, *y = *y, *x
}
