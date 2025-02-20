package main

import (
	"fmt"
	"sort"
	"strings"
)

type Product struct {
	Id       int
	Name     string
	Cost     float32
	Units    int
	Category string
}

func (product Product) String() string {
	return fmt.Sprintf("Id = %d, Name = %q, Cost = %.2f, Units = %d, Category = %q", product.Id, product.Name, product.Cost, product.Units, product.Category)
}

/*
Create a "Sort" api to sort the products by any attribute
	IMPORTANT: DO NOT write your own logic for sorting. instead use sort.Sort() function
*/

type Products []Product

func (products Products) String() string {
	var sb strings.Builder
	for _, p := range products {
		sb.WriteString(fmt.Sprintf("%s\n", p))
	}
	return sb.String()
}

// sort.Interface implmentation
func (products Products) Len() int {
	return len(products)
}

// Default sort (by id)
func (products Products) Less(i, j int) bool {
	return products[i].Id < products[j].Id
}

func (products Products) Swap(i, j int) {
	products[i], products[j] = products[j], products[i]
}

// Sort by Name
type ByName struct {
	Products
}

// override Less() method to compare by Name
func (byName ByName) Less(i, j int) bool {
	return byName.Products[i].Name < byName.Products[j].Name
}

// Sort by Cost
type ByCost struct {
	Products
}

// override Less() method to compare by Cost
func (byCost ByCost) Less(i, j int) bool {
	return byCost.Products[i].Cost < byCost.Products[j].Cost
}

// Sort by Units
type ByUnits struct {
	Products
}

// override Less() method to compare by Units
func (byUnits ByUnits) Less(i, j int) bool {
	return byUnits.Products[i].Units < byUnits.Products[j].Units
}

// Sort by Name
type ByCategory struct {
	Products
}

// override Less() method to compare by Category
func (byCategory ByCategory) Less(i, j int) bool {
	return byCategory.Products[i].Category < byCategory.Products[j].Category
}

// Utility method for sorting
func (products Products) Sort(by string) {
	switch by {
	case "Id":
		sort.Sort(products)
	case "Name":
		sort.Sort(ByName{Products: products})
	case "Cost":
		sort.Sort(ByCost{Products: products})
	case "Units":
		sort.Sort(ByUnits{Products: products})
	case "Category":
		sort.Sort(ByCategory{Products: products})
	}
}

func main() {
	products := Products{
		Product{105, "Pen", 5, 50, "Stationary"},
		Product{107, "Pencil", 2, 100, "Stationary"},
		Product{103, "Marker", 50, 20, "Utencil"},
		Product{102, "Stove", 5000, 5, "Utencil"},
		Product{101, "Kettle", 2500, 10, "Utencil"},
		Product{104, "Scribble Pad", 20, 20, "Stationary"},
		Product{109, "Golden Pen", 2000, 20, "Stationary"},
	}

	fmt.Println("Initial List")
	fmt.Println(products)

	// sort by Id & print
	fmt.Println("Products sorted by Id")
	// sort.Sort(products)
	products.Sort("Id")
	fmt.Println(products)

	// sort by Name & print
	fmt.Println("Products sorted by Name")
	// sort.Sort(ByName{Products: products})
	products.Sort("Name")
	fmt.Println(products)

	// sort by Cost & print
	fmt.Println("Products sorted by Cost")
	// sort.Sort(ByCost{Products: products})
	products.Sort("Cost")
	fmt.Println(products)

	// sort by Units & print
	fmt.Println("Products sorted by Units")
	// sort.Sort(ByUnits{Products: products})
	products.Sort("Units")
	fmt.Println(products)

	// sort by Category & print
	fmt.Println("Products sorted by Category")
	// sort.Sort(ByCategory{Products: products})
	products.Sort("Category")
	fmt.Println(products)
}
