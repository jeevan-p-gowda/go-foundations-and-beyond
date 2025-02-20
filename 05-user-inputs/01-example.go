package main

import "fmt"

/*
Note:

	cannot use fmt.Scan() if the input data has spaces
*/
func main() {
	/*
		var userName string
		fmt.Println("Enter the username :")
		fmt.Scanln(&userName)
		fmt.Printf("Hi %s, Have a nice day!\n", userName)
	*/

	var x, y int
	fmt.Scanln(&x, &y)
	fmt.Println(x + y)

}
