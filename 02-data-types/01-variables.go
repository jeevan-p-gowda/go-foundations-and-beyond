package main

import "fmt"

// package scoped variables
// var app_version string = "1.1.0"
var app_version = "1.1.0"

// app_version := "1.1.0" //==> will not work

// unused variable at package scope
var p_unused_var int = 100

func main() {
	// var userName string
	// userName = "Jeevan"

	// var userName string = "Jeevan"

	// type inference
	// var userName = "Jeevan"

	userName := "Jeevan"
	fmt.Printf("Hello, %s!\n", userName)

	// using package scoped variables
	fmt.Printf("App Version = %q\n", app_version)

	// unused variable
	var f_unused_var int = 100

	// reading the value from the variable makes it "used"
	f_unused_var = f_unused_var + 200

	// Multiple variables
	/*
		var x int
		var y int
		var str string
		var result int
		x = 100
		y = 200
		str = "sum of %d and %d is %d\n"
		result = x + y
		fmt.Printf(str, x, y, result)
	*/

	/*
		var x, y, result int
		var str string
		x = 100
		y = 200
		str = "sum of %d and %d is %d\n"
		result = x + y
		fmt.Printf(str, x, y, result)
	*/

	/*
		var (
			x, y, result int
			str          string
		)
		x = 100
		y = 200
		str = "sum of %d and %d is %d\n"
		result = x + y
		fmt.Printf(str, x, y, result)
	*/

	/*
		var x int = 100
		var y int = 200
		var str string = "sum of %d and %d is %d\n"
		var result int = x + y
		fmt.Printf(str, x, y, result)
	*/

	// multi assignment
	/*
		var x,y int = 100, 200
		var str string = "sum of %d and %d is %d\n"
		var result int = x + y
		fmt.Printf(str, x, y, result)
	*/

	/*
		var (
			x, y   int    = 100, 200
			str    string = "sum of %d and %d is %d\n"
			result int    = x + y
		)
		fmt.Printf(str, x, y, result)
	*/

	// type inference
	/*
		var (
			x, y   = 100, 200
			str    = "sum of %d and %d is %d\n"
			result = x + y
		)
		fmt.Printf(str, x, y, result)
	*/

	/*
		var (
			x, y, str = 100, 200, "sum of %d and %d is %d\n"
			result    = x + y
		)
		fmt.Printf(str, x, y, result)
	*/

	// using ":="
	/*
		x, y, str := 100, 200, "sum of %d and %d is %d\n"
		result := x + y
		fmt.Printf(str, x, y, result)
	*/

}
