package main

import (
	"fmt"

	"github.com/Avixph/learn-go/courses/learn-to-code-in-go/exercises/ninja-level-13/excercise-c/mymath"
)

func main() {
	xxi := gen()
	for _, v := range xxi {
		fmt.Println(mymath.CenteredAvg(v))
	}
}
