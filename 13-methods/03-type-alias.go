package main

import "fmt"

// MyString is an alias for "string"
type MyString string

func (s MyString) Length() int {
	return len(s)
}

func main() {
	var str string = "Labore qui ex consequat do ut pariatur fugiat nulla magna veniam aute."
	s := MyString(str)
	fmt.Println(s.Length())
}
