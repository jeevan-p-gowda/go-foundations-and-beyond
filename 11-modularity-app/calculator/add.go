package math

import "log"

func init() {
	log.Println("[math/add.go] init")
}

func Add(x, y int) int {
	// opCount++
	opCount["Add"]++
	return x + y
}
