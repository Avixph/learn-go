package main

import (
	"fmt"

	"github.com/Avixph/learn-go/courses/learn-to-code-in-go/exercises/ninja-level-13/excercise-b/quote"
	"github.com/Avixph/learn-go/courses/learn-to-code-in-go/exercises/ninja-level-13/excercise-b/word"
)

func main() {
	fmt.Println(word.Count(quote.SunAlso))

	for k, v := range word.UseCount(quote.SunAlso) {
		fmt.Println(v, k)
	}
}
