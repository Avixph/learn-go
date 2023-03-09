package main

import "fmt"

type person struct {
	firstName string
	lastName  string
	age       int
}

func (p person) speak() {
	fmt.Printf("Hi my name is %v %v and I'm %v years old.\n", p.firstName, p.lastName, p.age)
}
