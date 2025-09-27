package main

import "fmt"

func main() {
	/*var names map[string]int

	names = make(map[string]int, 0)

	names["Mert"] = 1
	names["Mustafa"] = 2
	names["Melek"] = 3

	fmt.Println(names)
	fmt.Println(names["Mert"])*/

	names := map[string]int{
		"Mert":    1,
		"Mustafa": 2,
		"Melek":   3,
	}

	delete(names, "Mert")

	fmt.Println(names)

}
