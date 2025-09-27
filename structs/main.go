package main

import "fmt"

func main() {

	var customer1 = Customer{ id: 1, name: "John", age: 30, address: Address{ street: "123 Main St", city: "New York" } }
	var customer2 = Customer{ 2, "Doe", 25, Address{ "456 Elm St", "Los Angeles" } }

	fmt.Println(customer1)
	
	fmt.Println(customer2)

	customer1.ToString()

	changeName(&customer1)
	
	customer1.ToString()

	customer1.changeName("Alice")
	
	customer1.ToString()
}

func changeName(customer *Customer) {
	customer.name = "Jane"
}

func (customer *Customer) changeName(newName string) {
	customer.name = newName
}

type Customer struct {
	id int64
	name string
	age int
	address Address
}

type Address struct {
	street string
	city string
}

func (customer *Customer) ToString(){
	fmt.Printf("Name: %s, Age: %d\n", customer.name, customer.age)
}