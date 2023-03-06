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
	peopleMap := map[string]person{
		jb.lastName: jb,
		mm.lastName: mm,
	}
	for _, person := range peopleMap {
		fmt.Printf("%v %v's favorite ice cream flavors are:\n", person.firstName, person.lastName)
		for i, favIceCream := range person.favoriteIceCream {
			fmt.Printf("%v %v \n", i+1, favIceCream)
		}
	}

}
