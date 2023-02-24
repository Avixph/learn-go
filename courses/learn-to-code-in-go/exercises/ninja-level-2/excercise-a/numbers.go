package main

import (
	"fmt"
	"math/rand"
	"time"
)

type number int

func newNumber() number {
	source := rand.NewSource(time.Now().UnixNano())
	random := rand.New(source)
	num := number(random.Intn(100))
	return num
}

func (num number) printNum() {
	fmt.Printf("Your number is %d\n* In binary it's %b\n* In hex it's %#x\n", num, num, num)
}
