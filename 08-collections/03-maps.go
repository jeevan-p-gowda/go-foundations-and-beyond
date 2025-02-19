package main

import (
	"fmt"
)

func main() {
	/*
		var productRanks map[string]int
		productRanks = make(map[string]int)
		productRanks["pen"] = 3
		productRanks["pencil"] = 1
		productRanks["marker"] = 5
	*/

	// var productRanks map[string]int = map[string]int{"pen": 3, "pencil": 1, "marker": 5}
	// var productRanks  = map[string]int{"pen": 3, "pencil": 1, "marker": 5}
	// productRanks := map[string]int{"pen": 3, "pencil": 1, "marker": 5}
	productRanks := map[string]int{
		"pen":    3,
		"pencil": 1,
		"marker": 5,
	}
	productRanks["notepad"] = 2
	fmt.Println(productRanks)

	fmt.Println("Iterating a map")
	for key, val := range productRanks {
		fmt.Printf("productRanks[%q] = %d\n", key, val)
	}

	// check for the existance of a key
	fmt.Println("check for the existance of a key")
	// keyToCheck := "marker"
	keyToCheck := "scribble pad"
	if rank, exists := productRanks[keyToCheck]; exists {
		fmt.Printf("Rank of %q is %d\n", keyToCheck, rank)
	} else {
		fmt.Printf("Key %q does not exist\n", keyToCheck)
	}

	fmt.Println("Removing an item")
	// keyToRemove := "pen"
	keyToRemove := "scribble pad"
	delete(productRanks, keyToRemove)
	fmt.Printf("After removing %q, productRanks = %v\n", keyToRemove, productRanks)
}
