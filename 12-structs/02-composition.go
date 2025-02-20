package main

import "fmt"

type Product struct {
	Id   int
	Name string
	Cost float64
}

type PerishableProduct struct {
	Product //composition
	Expiry  string
}

func NewPerishableProduct(id int, name string, cost float64, expiry string) *PerishableProduct {
	return &PerishableProduct{
		Product: Product{
			Id:   id,
			Name: name,
			Cost: cost,
		},
		Expiry: expiry,
	}
}

func main() {
	/*
		var milk PerishableProduct
		milk.Expiry = "2 Days"
		milk.Product.Id = 100
		milk.Product.Name = "Milk"
		milk.Product.Cost = 50
	*/

	/*
		var milk PerishableProduct = PerishableProduct{
			Product: Product{
				Id:   100,
				Name: "Milk",
				Cost: 50,
			},
			Expiry: "2 Days",
		}
	*/

	// Use the factory function to create an instance of PerishableProduct
	var milk = NewPerishableProduct(100, "Milk", 50, "2 Days")
	// fmt.Println(milk.Product.Id, milk.Product.Name, milk.Product.Cost, milk.Expiry)
	fmt.Println(milk.Id, milk.Name, milk.Cost, milk.Expiry)
	// fmt.Printf("%#v\n", milk)

	// Use the Format() and ApplyDiscount() for milk (PerishableProduct)
	fmt.Println("Before applying 10% discount:")
	// fmt.Println(Format(milk.Product))
	fmt.Println(FormatPerishableProduct(*milk))
	ApplyDiscount(&milk.Product, 10) // apply 10% discount on the proeduct (p)
	fmt.Println("After applying 10% discount:")
	// fmt.Println(Format(milk.Product))
	fmt.Println(FormatPerishableProduct(*milk))
}

// DO NOT modify the below functions
func Format(p Product) string {
	// return a formatted string of the given product
	// ex : Id = 100, Name = "Pen", Cost = 10.00
	return fmt.Sprintf("Id = %d, Name = %q, Cost = %0.2f", p.Id, p.Name, p.Cost)
}

func ApplyDiscount(p *Product, discountPercentage float64) { // no return values
	// update the cost of the given product by applying the given discount
	p.Cost = p.Cost * ((100 - discountPercentage) / 100)
}

func FormatPerishableProduct(pp PerishableProduct) string {
	return fmt.Sprintf("%s, Expiry = %q", Format(pp.Product), pp.Expiry)
}
