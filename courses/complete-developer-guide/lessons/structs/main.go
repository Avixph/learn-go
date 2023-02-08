package main

func main() {
	// angel := person{"Angel", "Fernandez", "New York, NY"}
	// jenn := person{firstName: "Jennifer", lastName: "Smith"}
	// var xavier person
	// xavier.firstName = "Xavier"
	// xavier.lastName = "Fernandez"
	// xavier.address = "New York, NY"
	// // fmt.Println(angel)
	// // fmt.Println(jenn)
	// // fmt.Println(xavier)
	// fmt.Printf("%+v", angel)
	// fmt.Printf("%+v", jenn)
	// fmt.Printf("%+v", xavier)

	lulu := person{firstName: "Lulu", lastName: "Fernandez", contact: contactInfo{email: "lulu@example.com", address: "New York, NY", zipcode: 10472}}
	lulu.print()

	angel := person{firstName: "Angle", lastName: "Fern", contact: contactInfo{email: "angel@example.com", address: "New York, NY", zipcode: 10472}}
	angel.print()

	// THe "&" in "&angel" tells the code to give you the memory address of the value it is pointing at
	angelPointer := &angel
	angelPointer.updateName("Angel", "Fernandez")
	angel.print()

}
