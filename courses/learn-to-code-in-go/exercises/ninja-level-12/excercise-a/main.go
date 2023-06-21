package main

import (
	"fmt"

	"github.com/Avixph/learn-go/courses/learn-to-code-in-go/exercises/ninja-level-12/excercise-a/dog"
)

func main() {
	lulu := dogo{
		name: "Lulu",
		age: ages{
			inHuman: 10,
			inDog:   dog.Years(10),
		},
	}

	fmt.Println(lulu)
}
