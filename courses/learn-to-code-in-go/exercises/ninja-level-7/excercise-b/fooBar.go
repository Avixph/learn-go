package main

type person struct {
	name string
	age  int
}

func (personPointer *person) changMe(newName string) {
	(*personPointer).name = newName
}
