package main

import "fmt"

func main() {
	book := &Book{desi: 10}

	fmt.Println("Shipping cost:", book.ShippingCost())

	var product IShippable

	product = &Book{ desi: 12}
	fmt.Println("Product shipping cost:", product.ShippingCost())

	product = &Flower{ desi: 8}
	fmt.Println("Product shipping cost:", product.ShippingCost())

	product = &Electronic{ desi: 20}
	fmt.Println("Product shipping cost:", product.ShippingCost())

	var products []IShippable = []IShippable{
		&Book{desi: 10},
		&Electronic{desi: 15},
		&Book{desi: 5},
		&Flower{desi: 7},
	}
	
	totalCost := calculateTotalShippingCost(products)
	fmt.Println("Total shipping cost:", totalCost)
}

type IShippable interface {
	ShippingCost() int
}

type Book struct {
	desi int

}

type Flower struct {
	desi int
}

type Electronic struct {
	desi int
}

func (book *Book)  ShippingCost() int{ 
	return 5 +  book.desi * 2
}

func (flower *Flower)  ShippingCost() int{ 
	return 5 +  flower.desi * 3
}

func (electronic *Electronic)  ShippingCost() int{ 
	return 5 +  electronic.desi * 4
}

func calculateTotalShippingCost(products []IShippable) int {
	total := 0
	for _, product := range products {
		total += product.ShippingCost()
	}
	return total
}