package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	rdr := bufio.NewReader(os.Stdin)

	var del byte
	_, ScnErr := fmt.Scanf("%c", &del)
	if ScnErr != nil {
		return
	}

	s, err := rdr.ReadString(del)
	if err != nil {
		log.Println(err)
	}

	fmt.Println(s)
}
