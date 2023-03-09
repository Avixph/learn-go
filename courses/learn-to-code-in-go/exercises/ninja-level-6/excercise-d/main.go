package main

func main() {
	p1 := person{firstName: "angel", lastName: "fernandez", age: 32}
	p2 := person{firstName: "jennifer", lastName: "smith", age: 27}
	p3 := person{firstName: "xavier", lastName: "fernandez", age: 25}
	p1.speak()
	p2.speak()
	p3.speak()
}
