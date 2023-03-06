package main

import "fmt"

func main() {
	jb := person{
		firstName:        "James",
		lastName:         "Bond",
		favoriteIceCream: []string{"martini", "chocolate"},
	}
	mm := person{
		firstName:        "Miss",
		lastName:         "Moneypenny",
		favoriteIceCream: []string{"hazelnut", "pistacio"},
	}
	fmt.Printf("%v %v's favorite ice cream flavors are:\n", jb.firstName, jb.lastName)
	for i, favIceCream := range jb.favoriteIceCream {
		fmt.Printf("%v %v \n", i+1, favIceCream)
	}
	fmt.Printf("%v %v's favorite ice cream flavors are:\n", mm.firstName, mm.lastName)
	for i, favIceCream := range mm.favoriteIceCream {
		fmt.Printf("%v %v \n", i+1, favIceCream)
	}
}
