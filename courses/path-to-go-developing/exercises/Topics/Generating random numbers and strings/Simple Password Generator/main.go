package main

import (
	"fmt"
	"math/rand"
)

func main() {
	upperCase := "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	lowerCase := "abcdefghijklmnopqrstuvwxyz"
	number := "0123456789"
	specialSymbol := "!?$&%#"

	var seed int64
	_, err := fmt.Scanln(&seed)
	if err != nil {
		return
	}
	rand.Seed(seed)

	// put your code here

}
