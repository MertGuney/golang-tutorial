package main

import "fmt"

func main() {

	/*var a int
	var p *int

	a = 5
	p = &a

	fmt.Println("Value of a:", a)
	fmt.Println("Address of a:", p)
	fmt.Println("Value at address p:", *p)

	*p = 10

	fmt.Println("Updated value of a:", a)
	fmt.Println("Address of a:", p)
	fmt.Println("Value at address p:", *p)*/

	/*var a = 10
	var b int
	var p *int

	b = a
	p = &a
	*p = 20

	println("Value of a:", a)
	println("Value of b:", b)*/

	/*var a = 10

	fmt.Println("Value of a:", a)

	add12Ptr(&a)

	fmt.Println("Value of a after add12:", a)*/

	var numbers = []int{1, 2, 3, 4, 5}

	fmt.Println("Original slice:", numbers)

	changeValues(numbers)

	fmt.Println("Modified slice:", numbers)
}

/*func add12(a int) int {
	return a + 12
}

func add12Ptr(a *int) {
	*a = *a + 12
}*/

func changeValues(numbers []int) {
	numbers[0] = 100
}
