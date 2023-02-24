package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	source := rand.NewSource(time.Now().UnixNano())
	random := rand.New(source)
	num := random.Intn(51)
	a := (num == 41)
	b := (num <= 33)
	c := (num >= 25)
	d := (num != 17)
	e := (num < 9)
	f := (num > 11)
	fmt.Println(num)
	fmt.Println(a, b, c, d, e, f)
}
