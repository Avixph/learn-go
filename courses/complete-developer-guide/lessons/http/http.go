package main

import (
	"fmt"
)

type logWriter struct{}

func (logWriter) Write(byteSlice []byte) (int, error) {
	fmt.Println(string(byteSlice))
	fmt.Println("Just wrote this many bytes:", len(byteSlice))
	return len(byteSlice), nil
}
