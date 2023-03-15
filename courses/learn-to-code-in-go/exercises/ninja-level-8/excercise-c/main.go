package main

import (
	"encoding/json"
	"fmt"
	"os"
)

func main() {
	uA := user{
		FirstName: "James",
		LastName:  "bond",
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
		LastName:  "",
		Age:       54,
		Sayings: []string{
			"Oh, James. You didn't.",
			"Dear God, what has James done now?",
			"Can someone please tell me where James Bond is?",
		},
	}

	users := []user{uA, uB, uC}
	fmt.Println(users)

	err := json.NewEncoder(os.Stdout).Encode(users)
	if err != nil {
		fmt.Println("error:", err)
	}
}
