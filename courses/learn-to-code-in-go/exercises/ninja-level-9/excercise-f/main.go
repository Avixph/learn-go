package main

import (
	"fmt"
	"runtime"
)

func main() {
	os := runtime.GOOS
	arch := runtime.GOARCH
	fmt.Println("Your OS --->", os)
	fmt.Println("Your ARCH --->", arch)
}
