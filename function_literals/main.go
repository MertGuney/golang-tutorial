package main

import (
	"fmt"
	"reflect"
)

func main() {
	func() {
		fmt.Println("f1")
	}()

	func(x int, y int) {
		fmt.Println(x + y)
	}(4, 5)

	add := func(x int, y int) int {
		return x + y
	}

	fmt.Println(reflect.TypeOf(add))

	fmt.Println(add(3, 5))

	multiply := func(x int, y int) int {
		return x * y
	}

	a, b := calculator(3, 5, add, multiply)

	fmt.Println(a, b)
}

func calculator(x int, y int, f1 func(int, int) int, f2 func(int, int) int) (int, int) {
	return f1(x, y), f2(x, y)
}
