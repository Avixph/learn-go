package main

import "fmt"

func main() {
	fb := fooBar()
	fmt.Println(fb())
	for _, f := range fb() {
		fmt.Println(f)
	}
}
