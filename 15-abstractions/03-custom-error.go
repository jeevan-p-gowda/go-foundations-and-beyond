package main

import (
	"fmt"
)

func main() {
	divisor := 0
	if q, r, e := divide(100, divisor); e != nil {
		fmt.Println(e)
		return
	} else {
		fmt.Println(q, r)
	}
}

func divide(x, divisor int) (int, int, error) {
	if divisor != 0 {
		quotient, remainder := x/divisor, x%divisor
		return quotient, remainder, nil
	}
	return 0, 0, NewDivideError(x, divisor)
}

// custom error type
type DivideError struct {
	multiplier int
	divisor    int
}

// error interface implementation
func (de DivideError) Error() string {
	return fmt.Sprintf("divide error : multiplier = %d, divisor = %d", de.multiplier, de.divisor)
}

func NewDivideError(m, d int) *DivideError {
	return &DivideError{
		multiplier: m,
		divisor:    d,
	}
}
