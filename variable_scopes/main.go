package main

import "fmt"

var z = 10

func main() {

	// var condition = true
	// if condition {
	// 	var x = 10
	// 	fmt.Println(x)
	// }
	// fmt.Println(x)

	// var condition = true
	// if condition {
	// 	fmt.Println(condition)
	// }
	// fmt.Println(condition)

	// test()
	// fmt.Println(x, y)

	// fmt.Println(z)

	fmt.Println(z)
	test2()
	fmt.Println(z)
	test3()
	fmt.Println(z)
}

func test() {
	var x = 10
	var y = 20

	fmt.Println(x, y)
}

func test2() {
	var z = 20
	fmt.Println(z)
}

func test3() {
	z = 30
	fmt.Println(z)
}
