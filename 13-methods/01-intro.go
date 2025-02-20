package main

import "fmt"

type Product struct {
	Id   int
	Name string
	Cost float64
}

func main() {
	var p Product = Product{
		Id:   100,
		Name: "Pen",
		Cost: 10,
	}

	fmt.Println("Before applying 10% discount:")

	// fmt.Println(Format(p))
	fmt.Println(p.Format())

	// ApplyDiscount(&p, 10) // apply 10% discount on the proeduct (p)
	p.ApplyDiscount(10)

	fmt.Println("After applying 10% discount:")

	// fmt.Println(Format(p))
	fmt.Println(p.Format())

}

// Method = function with a receiver

func (p Product) Format() string {
	return fmt.Sprintf("Id = %d, Name = %q, Cost = %0.2f", p.Id, p.Name, p.Cost)
}

func (p *Product) ApplyDiscount(discountPercentage float64) {
	p.Cost = p.Cost * ((100 - discountPercentage) / 100)
}
