package main

import "fmt"

type Product struct {
	Id   int
	Name string
	Cost float64
}

func main() {
	/*
		var p Product // memory allocated and attributes are initialized
		p.Id = 100
		p.Name = "Pen"
		p.Cost = 10
	*/

	// var p Product = Product{100, "Pen", 10}
	// var p Product = Product{Id: 100, Name: "Pen", Cost: 10}
	var p Product = Product{
		Id:   100,
		Name: "Pen",
		Cost: 10,
	}
	// fmt.Println(p)
	// fmt.Printf("%#v\n", p)
	fmt.Printf("%+v\n", p)

	// create a copy
	/*
		p2 := p
		p2.Cost = 20
		fmt.Printf("p.Cost = %v, p2.Cost = %v\n", p.Cost, p2.Cost)
	*/

	// create a reference
	p2 := &p
	p2.Cost = 20
	fmt.Printf("p.Cost = %v, p2.Cost = %v\n", p.Cost, p2.Cost)

	fmt.Println("Before applying 10% discount:")
	fmt.Println(Format(p))
	ApplyDiscount(&p, 10) // apply 10% discount on the proeduct (p)
	fmt.Println("After applying 10% discount:")
	fmt.Println(Format(p))
	fmt.Println(Format(*p2))

}

func Format(p Product) string {
	// return a formatted string of the given product
	// ex : Id = 100, Name = "Pen", Cost = 10.00
	return fmt.Sprintf("Id = %d, Name = %q, Cost = %0.2f", p.Id, p.Name, p.Cost)
}

func ApplyDiscount(p *Product, discountPercentage float64) { // no return values
	// update the cost of the given product by applying the given discount
	p.Cost = p.Cost * ((100 - discountPercentage) / 100)
}
