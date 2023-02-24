package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	source := rand.NewSource(time.Now().UnixNano())
	random := rand.New(source)
	var numA int = random.Intn(21)
	fmt.Printf("Your number ---> %d\n* In binary ---> %b\n* In hexadecimal ---> %#x\n", numA, numA, numA)
	numB := numA << 2
	fmt.Printf("Your number ---> %d\n* In binary ---> %b\n* In hexadecimal ---> %#x\n", numB, numB, numB)
}
