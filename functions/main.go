package main

func main() {
	/*result := add(3, 5)
	fmt.Println("The result is:", result)*/

	/*total, difference := calculation(10, 4)
	fmt.Println("The total is:", total)
	fmt.Println("The difference is:", difference)*/

	/*var numbers = []int{1, 2, 3, 4, 5}
	total := sum(numbers)
	fmt.Println("The sum of the numbers is:", total)*/

	total := sum(1, 2, 3, 4, 5, 6, 7)
	println("The sum of the numbers is:", total)
}

/*func add(a, b int) int {
	return a + b
}*/

/*func calculation(a, b int) (int, int) {
	return a + b, a - b
}*/

/*func sum(numbers []int) int {
	total := 0
	for _, number := range numbers {
		total += number
	}
	return total
}*/

func sum(numbers ...int) int {
	total := 0
	for _, number := range numbers {
		total += number
	}
	return total
}
