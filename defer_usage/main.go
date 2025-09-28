package main

import (
	"fmt"
)

func main() {

	defer fmt.Println("First deferred statement")
	defer fmt.Println("Second deferred statement")
	defer fmt.Println("Third deferred statement")

	fmt.Println("Main function started")
	/*--------------------------------------------------*/
	var condition = false

	defer fmt.Println("1. Deferred statement after condition check")

	if condition {
		return
	}

	defer fmt.Println("2. Deferred statement after return")
	/*--------------------------------------------------*/
	x := 10
	y := 20

	defer fmt.Println("X:", x)
	defer fmt.Println("Y:", y)

	x += 5
	y += 10

	fmt.Println("Updated X:", x)
	fmt.Println("Updated Y:", y)
	/*--------------------------------------------------*/
	condition = true

	defer cleanup()

	if condition {
		defer panic("A panic occurred!")
	}

}

func cleanup() {
	fmt.Println("Cleanup worked...")
}
