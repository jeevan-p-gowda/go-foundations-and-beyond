/*
Accept a range (start & end) from the user and print all the prime numbers between the given range
*/
package main

import "fmt"

func main() {
	var start, end int
	fmt.Println("Enter the start and end :")
	fmt.Scanln(&start, &end)

OUTER_LOOP:
	for no := start; no <= end; no++ {
		if no < 2 { // Skip numbers less than 2, as they are not prime
			continue
		}

		for i := 2; i <= (no / 2); i++ {
			if no%i == 0 {
				continue OUTER_LOOP
			}
		}
		fmt.Println("Prime No :", no)
	}
}
