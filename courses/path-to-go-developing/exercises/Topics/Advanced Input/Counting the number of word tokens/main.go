package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	scn := bufio.NewScanner(os.Stdin)

	// Set the correct split function below:
	scn.Split(bufio.ScanWords)

	// Write the for loop below with the correct scanner statement
	count := 0
	for scn.Scan() {
		count++
	}
	fmt.Println(count)
}
