package main

import "fmt"

func main(){
	/*Array*/
	/*var names [3]string
	
	names[0] = "Mert"
	names[1] = "Melek"
	names[2] = "Mustafa"

	fmt.Println(names)*/

	/*var names = [3] string {"Mert", "Mustafa", "Melek"}

	fmt.Println(names)

	fmt.Println(names[0:2])*/

	/*Slice*/
	var names = []string {"Mert", "Mustafa", "Melek"}

	names = append(names, "Özlem")

	fmt.Println(names)
}
