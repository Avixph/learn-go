package main

import (
	"fmt"
	"sort"
)

func main() {
	x := []int{5, 8, 2, 43, 17, 987, 14, 12, 21, 1, 4, 2, 3, 93, 13}
	y := []string{"in", "torpedo", "summers", "random", "rainbow", "delights", "under", "gallantry", "fragmented", "moons", "across", "magenta"}

	fmt.Println(x, "\n------")
	// sort x
	sort.Ints(x)
	fmt.Println(x)

	fmt.Println(y, "\n------")
	// sort y
	sort.Strings(y)
	fmt.Println(y)

}
