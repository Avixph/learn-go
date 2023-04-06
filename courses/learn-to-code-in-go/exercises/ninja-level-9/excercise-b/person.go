package main

import "fmt"

type person struct {
	firstName string
	lastName  string
}

type human interface {
	speak() string
}

func (p *person) speak() string {
	greeting := "Hello my name is " + p.firstName + " " + p.lastName + "!"
	return greeting
}

func saySomething(h human) {
	fmt.Println(h.speak())
}
