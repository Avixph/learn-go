package main

import (
	"fmt"
	"sort"
)

func main() {
	uA := user{
		FirstName: "James",
		LastName:  "Bond",
		Age:       32,
		Sayings: []string{
			"Shaken, not stirred",
			"Youth is no guarantee of innovation",
			"In his majesty's royal service",
		},
	}
	uB := user{
		FirstName: "Miss",
		LastName:  "Moneypenny",
		Age:       27,
		Sayings: []string{
			"James, it is soo good to see you",
			"Would you like me to take care of that for you, James?",
			"I would really prefer to be a secret agent myself.",
		},
	}
	uC := user{
		FirstName: "M",
		LastName:  "_____",
		Age:       54,
		Sayings: []string{
			"Oh, James. You didn't.",
			"Dear God, what has James done now?",
			"Can someone please tell me where James Bond is?",
		},
	}

	users := []user{uA, uB, uC}
	for _, user := range users {
		fmt.Println("Name:", user.FirstName, user.LastName, "\nAge:", user.Age)
		fmt.Println("Sayings:")
		for _, saying := range user.Sayings {
			fmt.Printf("'%v'\n", saying)
		}
	}
	fmt.Println("------")
	sort.Sort(ByAge(users))
	for _, user := range users {
		fmt.Println("Name:", user.FirstName, user.LastName, "\nAge:", user.Age)
		fmt.Println("Sayings:")
		sort.Strings(user.Sayings)
		for _, saying := range user.Sayings {
			fmt.Printf("'%v'\n", saying)
		}
	}
	fmt.Println("------")
	sort.Sort(ByLastName(users))
	for _, user := range users {
		fmt.Println("Name:", user.FirstName, user.LastName, "\nAge:", user.Age)
		fmt.Println("Sayings:")
		sort.Strings(user.Sayings)
		for _, saying := range user.Sayings {
			fmt.Printf("'%v'\n", saying)
		}
	}
}
