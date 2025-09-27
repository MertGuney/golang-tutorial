package main

import (
	"fmt"
)

func main() {
	var numbers = []int{1, 2, 3, 4}

	for i := 0; i < len(numbers); i++ {
		fmt.Println(numbers[i])
	}

	for index, value := range numbers {
		fmt.Println(index, value)
	}

	for _, value := range numbers {
		fmt.Println(value)
	}

	var language = "GoLang"

	for _, value := range language {
		fmt.Printf("Character %c\n", value)
	}

	var names = map[string]int{
		"Mustafa": 10,
		"Mert":    5,
		"Melek":   15,
	}

	for key, value := range names {
		fmt.Println(key, value)
	}
}
