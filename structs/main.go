package main

import "fmt"

type person struct {
	firstName string
	lastName string
}

func main() {
	alex := person{"Alex", "Novak"}
	john := person{firstName: "John", lastName:"Novak"}

	fmt.Println(alex)
	fmt.Println(john)

}