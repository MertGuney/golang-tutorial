package main

import "fmt"

func main() {
	/*index := 1

	for index <= 10 {
		fmt.Println("Index", index)
		index++
	}*/

	/*total := 0

	for index := 1; index <= 10; index++ {
		total += index
	}

	fmt.Println("Total", total)*/

	/*index := 0

	var names = [3]string{"Mert", "Mustafa", "Melek"}

	for index < len(names) {
		fmt.Println(names[index])
		index++
	}*/

	for index := 0; index <= 10; index++ {
		if index == 3 || index == 5 {
			continue
		} else if index == 7 {
			break
		}
		fmt.Println("Index", index)
	}

}
