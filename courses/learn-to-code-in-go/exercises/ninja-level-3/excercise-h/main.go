package main

import "fmt"

func main() {
	switch {
	case !(24 != (6 * 4)):
		fmt.Println(true)
	case 24 == 32:
		fmt.Println(false)
	}
}
