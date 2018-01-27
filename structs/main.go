package main

import "fmt"

type contactInfo struct {
	email string
	zip int
}

type person struct {
	firstName string
	lastName string
	contactInfo
}

func main() {
	// 1. example
	alex := person{firstName: "Alex",lastName: "Novak", contactInfo: contactInfo{email: "asd@asd.cz",zip: 14111}}
	john := person{
		firstName: 	"John", 
		lastName:	"Novak",
		contactInfo: contactInfo{
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
	fmt.Printf("%+v\n", martin)

	fmt.Println("Print martin")
	martin.updateName("John")
	martin.print()


}

func (p person) print() {
	fmt.Printf("%+v\n", p)
}

func (p *person) updateName(newFirstName string){
	p.firstName = newFirstName
}