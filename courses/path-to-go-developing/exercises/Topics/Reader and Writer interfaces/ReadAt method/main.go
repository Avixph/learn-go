package main

import (
	"fmt"
	"log"
	"strings"
)

const quote = "Knowledge is better than money"

// nolint: gomnd // DO NOT delete this comment!
func main() {
	reader := strings.NewReader(quote)

	buf := make([]byte, 5)
	_, err := reader.ReadAt(buf, 25)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(buf))
}
