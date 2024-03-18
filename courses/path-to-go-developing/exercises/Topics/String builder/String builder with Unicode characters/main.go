package main

import (
	"fmt"
	"strings"
)

func scanRune() rune {
	var c rune
	_, err := fmt.Scanln(&c)
	if err != nil {
		return 0
	}
	return c
}

func convertToUnicode(b strings.Builder, c rune) string {
	for i := 0; i < 3; i++ {
		b.WriteRune(c)
	}
	return b.String()
}

func main() {
	var b strings.Builder
	c := scanRune()
	fmt.Println(convertToUnicode(b, c))
}
