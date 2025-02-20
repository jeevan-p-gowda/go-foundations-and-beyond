package math

import "log"

/*
var opCount int

func OpCount() int {
	return opCount
}
*/

var opCount map[string]int

func init() {
	log.Println("[math/math.go] init")
	opCount = make(map[string]int, 0)
}

func OpCount() map[string]int {
	return opCount
}
