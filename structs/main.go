package main

import "fmt"

type person struct {
	firstName string
	lastName string
}

func main() {
	// 1. example
	alex := person{"Alex", "Novak"}
	john := person{firstName: "John", lastName:"Novak"}

	fmt.Println(alex)
	fmt.Println(john)

	// 2. example update
	var martin person

	martin.firstName = "Martin"
	martin.lastName = "Hlavacek"

	fmt.Println(martin)
	fmt.Printf("%+v", martin)



}