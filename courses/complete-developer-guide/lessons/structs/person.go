package main

import "fmt"

type contactInfo struct {
	email   string
	address string
	zipcode int
}

type person struct {
	firstName string
	lastName  string
	contact   contactInfo
}

func (p person) print() {
	fmt.Printf("%+v", p)
}

// "*person" is a type description which indicates we're working with a pointer to a person
func (personPointer *person) updateName(newFirstName string, newLastName string) {
	// THe "*" in "*personPointer" tells the code to give you the value this memory adress is pointing at
	// This indicates that we want to manipulate the value the pointer is refrencing to
	(*personPointer).firstName = newFirstName
	(*personPointer).lastName = newLastName
}
