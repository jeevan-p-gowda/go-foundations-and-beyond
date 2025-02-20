/*
Refactor this as follows

 1. IsPrime(no int) => responsible for checking if the given number is prime number
 2. GenPrimes(start, end)
    => Gen prime numbers CONCURRENTLY
 3. main()
    => Printing the generated prime numbers
    => Execute the GenPrimes() as a goroutine
*/
package main

import (
	"fmt"
	"sync"
)

func main() {
	var start, end int
	ch := make(chan int)

	fmt.Println("Enter the start and end :")
	fmt.Scanln(&start, &end)

	go GenPrimes(start, end, ch)

	for primeNo := range ch {
		fmt.Println("Prime No :", primeNo)
	}
}

func GenPrimes(start, end int, ch chan int) {
	wg2 := &sync.WaitGroup{}
	for no := start; no <= end; no++ {
		wg2.Add(1)
		go func() {
			defer wg2.Done()
			if IsPrime(no) {
				ch <- no
			}
		}()
	}
	wg2.Wait()
	close(ch)
}

func IsPrime(no int) bool {
	for i := 2; i <= (no / 2); i++ {
		if no%i == 0 {
			return false
		}
	}
	return true
}
