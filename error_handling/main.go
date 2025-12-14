package main

import (
	"errors"
	"fmt"
	"os"
)

func main() {
	// var x int
	// var y float32

	// fmt.Println(x, y)

	// var pointer1 *int
	// fmt.Println(pointer1)

	// var pointer2 *int
	// if pointer2 == nil {
	// 	fmt.Println("Nil pointer")
	// }

	fileData, err := os.ReadFile("sample.txt")

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(fileData)
	}

	result, err := divide(10, 0)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(result)
	}

	productService := ProductService{}

	er := productService.Add(Product{id: 1, name: "", price: 100})

	if er != nil {
		fmt.Println(er)
	}
}

type Product struct {
	id    int
	name  string
	price int
}
type ProductService struct {
}

type ValidationError struct {
	code string
	text string
}

func (validationError ValidationError) Error() string {
	return fmt.Sprintf("Error: %s, Code: %s", validationError.text, validationError.text)
}

// func (productService *ProductService) Add(product Product) error {
// 	if len(product.name) == 0 {
// 		return errors.New("Product name required!")
// 	}
// 	if product.price < 0 {
// 		return errors.New("Product price must be greater then 0")
// 	}
// 	fmt.Println("Product added")
// 	return nil
// }

func (productService *ProductService) Add(product Product) error {
	if len(product.name) == 0 {
		return ValidationError{text: "Product name required!", code: "1"}
	}
	if product.price < 0 {
		return ValidationError{text: "Product price must be greater then 0", code: "2"}
	}
	fmt.Println("Product added")
	return nil
}

func divide(x int, y int) (int, error) {
	if y == 0 {
		return 0, errors.New("Error: 0")
	}
	return x / y, nil
}
