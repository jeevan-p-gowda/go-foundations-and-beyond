package main

import (
	"fmt"

	"github.com/jeevan-p-gowda/go-foundations-and-beyond/11-modularity-app/calculator"
	/* "github.com/tkmagesh/Cisco-GoFoundation-Feb-2025/11-modularity-app/math/utils" */
	"github.com/fatih/color"
	ut "github.com/jeevan-p-gowda/go-foundations-and-beyond/11-modularity-app/calculator/utils"
)

func main() {
	color.Red("Hello, World!")
	fmt.Println(calculator.Add(100, 200))
	fmt.Println(calculator.Subtract(100, 200))
	fmt.Println("Operation Count :", calculator.OpCount())
	// fmt.Println("is 21 even? :", utils.IsEven(21))
	fmt.Println("is 21 even? :", ut.IsEven(21))
}
