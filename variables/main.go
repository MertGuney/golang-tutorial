package main

import "fmt"

func main() {
	/*var productName string
	var quantity int
	var discount float32
	var isInStock bool

	productName = "GoLang"
	quantity = 10
	discount = 0.15
	isInStock = true

	fmt.Println("Product Name:", productName, reflect.TypeOf(productName))
	fmt.Println("Quantity:", quantity, reflect.TypeOf(quantity))
	fmt.Println("Discount:", discount, reflect.TypeOf(discount))
	fmt.Println("Is In Stock:", isInStock, reflect.TypeOf(isInStock))*/

	/*productName := "GoLang"
	quantity := 10
	discount := 0.15
	isInStock := true

	fmt.Println("Product Name:", productName, reflect.TypeOf(productName))
	fmt.Println("Quantity:", quantity, reflect.TypeOf(quantity))
	fmt.Println("Discount:", discount, reflect.TypeOf(discount))
	fmt.Println("Is In Stock:", isInStock, reflect.TypeOf(isInStock))*/

	productName, quantity, discount, isInStock := "GoLang", 10, 0.15, true

	fmt.Println("Product Name:", productName, "Quantity:", quantity, "Discount:", discount, "Is In Stock:", isInStock)

	fmt.Printf("Product Name: %s, Quantity: %d, Discount: %.2f, Is In Stock: %t\n", productName, quantity, discount, isInStock)
}
