package main

import "fmt"

func main() {
	var source float64
	_, err := fmt.Scan(&source)
	if err != nil {
		return
	}

	fmt.Println(int(source))
}
