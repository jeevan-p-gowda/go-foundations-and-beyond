package math

import "log"

func init() {
	log.Println("[math/subtract.go] init")
}

func Subtract(x, y int) int {
	// opCount++
	opCount["Subtract"]++
	return x - y
}
