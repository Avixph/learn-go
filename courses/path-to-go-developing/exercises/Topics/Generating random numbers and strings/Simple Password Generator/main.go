package main

import (
	"fmt"
	"math/rand"
)

func main() {
	upperCase := "ABCDEFGHIJKLMOPQRSTUVWXYZ"
	lowerCase := "abcdefghijklmnopqrstuvwxyz"
	number := "0123456789"
	specialSymbol := "!?$&%#"

	var seed int64
	fmt.Scanln(&seed)
	rand.Seed(seed)

	// put your code here
	l := 4

	password := make([]byte, l)
	password[0] = upperCase[rand.Intn(len(upperCase))]
	password[1] = lowerCase[rand.Intn(len(lowerCase))]
	password[2] = number[rand.Intn(len(number))]
	password[3] = specialSymbol[rand.Intn(len(specialSymbol))]

	fmt.Println(string(password))
}
