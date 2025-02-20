package main

import "fmt"

type Product struct {
	Id   int
	Name string
	Cost float64
}

// Method = function with a receiver
func (p Product) Format() string {
	return fmt.Sprintf("Id = %d, Name = %q, Cost = %0.2f", p.Id, p.Name, p.Cost)
}

// fmt.Stringer interface implementation
func (p Product) String() string {
	// fmt.Println("Product.String() invoked")
	return fmt.Sprintf("Id = %d, Name = %q, Cost = %0.2f", p.Id, p.Name, p.Cost)
}

func (p *Product) ApplyDiscount(discountPercentage float64) {
	p.Cost = p.Cost * ((100 - discountPercentage) / 100)
}

/* ********** */
type PerishableProduct struct {
	Product //composition
	Expiry  string
}

// overriding the Product.Format() method
func (pp PerishableProduct) Format() string {
	return fmt.Sprintf("%s, Expiry = %q", pp.Product.Format(), pp.Expiry)
}

// fmt.Stringer interface implementation
func (pp PerishableProduct) String() string {
	// fmt.Println("PerishableProduct.String() invoked")
	return fmt.Sprintf("%s, Expiry = %q", pp.Product, pp.Expiry)
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

	var milk = NewPerishableProduct(100, "Milk", 50, "2 Days")

	// Use the Format() and ApplyDiscount() for milk (PerishableProduct)
	fmt.Println("Before applying 10% discount:")
	// fmt.Println(milk.Format())
	fmt.Println(milk)
	// fmt.Println(FormatPerishableProduct(*milk))

	milk.ApplyDiscount(10)

	fmt.Println("After applying 10% discount:")
	// fmt.Println(Format(milk.Product))
	// fmt.Println(milk.Format())
	fmt.Println(milk)
}
