package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	file, err := os.Open("numbers.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	var oddCount int
	scn := bufio.NewScanner(file)

	for scn.Scan() {
		var num int
		num, err = strconv.Atoi(scn.Text())
		if err != nil {
			panic(err)
		}

		if num%2 != 0 {
			oddCount++
		}
	}
	fmt.Println(oddCount)
}
