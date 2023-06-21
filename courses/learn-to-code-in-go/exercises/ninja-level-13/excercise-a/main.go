package main

import (
	"fmt"

	"github.com/Avixph/learn-go/courses/learn-to-code-in-go/exercises/ninja-level-13/excercise-a/dog"
)

func main() {
	mia := dogo{
		name: "Mia",
		age: ages{
			inHuman: 10,
			inDog:   dog.Years(10),
		},
	}

	fluffy := dogo{
		name: "Fluffy",
		age: ages{
			inHuman: 7,
			inDog:   dog.Years(7),
		},
	}

	fmt.Println(mia, "\n", fluffy)
	fmt.Println(dog.YearsTwo(20))
}
