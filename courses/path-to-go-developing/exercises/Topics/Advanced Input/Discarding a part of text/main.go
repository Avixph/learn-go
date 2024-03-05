package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

const discardedBytes = 10

func main() {
	rdr := bufio.NewReader(os.Stdin)

	_, err := rdr.Discard(discardedBytes)

	if err != nil {
		fmt.Println(err)
	}

	data, err := rdr.ReadString('\n')
	if err != nil && err != io.EOF {
		fmt.Println(err)
	}

	fmt.Println(data)
}
