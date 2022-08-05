package main

import (
	"fmt"
)

type person struct {
	firstName string
	lastName  string
	address struct {
		street string
		num int
	}
	contactInfo 
}

type contactInfo struct {
	email string
	zipCode int
}

func main() {
	// juan := person{"Juan", "Arean"}

	// var juan person
	// juan.firstName = "Juan"
	// juan.lastName = "Arean"
	
	juan := person{firstName: "Juan", lastName: "Arean", contactInfo: contactInfo{email: "juan@gmail.com", zipCode: 1644}}
	
	fmt.Println(juan.firstName)

	// juanPointer := &juan
	// juanPointer.updateName("Alberto")

	juan.updateName("Alberto") // go entiende que le queres mandar el puntero a juan. es un "shortcut"

	juan.print()
}

func (pointerToPerson *person) updateName(newFirstName string) {
	(*pointerToPerson).firstName = newFirstName
}

func (p person) print() {
	fmt.Printf("%+v", p)
}