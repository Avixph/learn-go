package main

import (
	"encoding/json"
	"fmt"
	"log"
)

func main() {
	p1 := person{
		First: "James",
		Last:  "Bond",
		Sayings: []string{
			"Shaken, not stirred",
			"Any last wishes?",
			"Never say never"},
	}

	// bs, _ := json.Marshal(p1)
	bs, err := json.Marshal(p1)

	if err != nil {
		log.Fatal("Error:", err)
	}

	fmt.Println(string(bs))
}
