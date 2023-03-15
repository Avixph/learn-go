package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	uA := user{
		Name: "James Bond", Age: 32,
	}
	uB := user{
		Name: "Miss Moneypenny", Age: 27,
	}
	uC := user{
		Name: "M", Age: 54,
	}

	users := []user{uA, uB, uC}

	fmt.Println(users)

	bs, err := json.Marshal(users)
	if err != nil {
		fmt.Println("error:", err)
	}
	fmt.Println(string(bs))
}
