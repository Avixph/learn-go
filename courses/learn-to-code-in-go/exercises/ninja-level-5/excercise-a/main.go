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
	fmt.Printf("%v\n %v", jb, mm)
}
