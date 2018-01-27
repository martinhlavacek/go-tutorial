package main

import "fmt"

type contactInfo struct {
	email string
	zip int
}

type person struct {
	firstName string
	lastName string
	contact contactInfo
}

func main() {
	// 1. example
	alex := person{firstName: "Alex",lastName: "Novak", contact: contactInfo{email: "asd@asd.cz",zip: 14111}}
	john := person{
		firstName: 	"John", 
		lastName:	"Novak",
		contact: contactInfo{
			email: "ll$asd.com",
			zip: 15000,
		},
	}
	fmt.Println("Alex")
	fmt.Println(alex)
	fmt.Println("John")
	fmt.Printf("%+v\n", john)

	// 2. example update
	var martin person

	martin.firstName = "Martin"
	martin.lastName = "Hlavacek"

	fmt.Println("Martin")
	fmt.Println(martin)
	fmt.Printf("%+v", martin)



}