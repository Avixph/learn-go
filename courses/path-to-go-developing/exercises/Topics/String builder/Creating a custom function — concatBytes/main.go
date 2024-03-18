package main

import (
	"fmt"
	"strings"
)

func concatBytes(bytes ...byte) string {
	var b strings.Builder
	for _, x := range bytes {
		b.WriteByte(x)
	}
	return b.String()
}

func main() {
	fmt.Println(concatBytes('J', 'B', 'A', 'c', 'a', 'd', 'e', 'm', 'y'))
}
