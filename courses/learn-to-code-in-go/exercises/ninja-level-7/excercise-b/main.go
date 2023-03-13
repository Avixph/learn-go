package main

import "fmt"

func main() {
	p1 := person{name: "angle fernadez", age: 32}
	fmt.Println(p1.name)
	p1.changMe("angel fernandez")
	fmt.Println(p1.name)
}
