package main

import "fmt"

func main() {
	// Create a set named 'mySet' with an empty struct as the value type below.
	// The set can have as many keys as you want; you can get creative!
	mySet := map[string]struct{}{
		"tiger":   struct{}{},
		"snake":   struct{}{},
		"monkey":  struct{}{},
		"horse":   struct{}{},
		"dog":     struct{}{},
		"dragon":  struct{}{},
		"rabbit":  struct{}{},
		"rat":     struct{}{},
		"rooster": struct{}{},
		"ox":      struct{}{},
		"pig":     struct{}{},
		"goat":    struct{}{},
	}
	fmt.Println()
	checkSetTypeValues(mySet) // DO NOT delete this line!
}
