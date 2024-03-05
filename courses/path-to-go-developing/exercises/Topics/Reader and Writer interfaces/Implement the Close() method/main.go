package main

import (
	"fmt"
	"io"
)

type FunnyBox struct {
}

func (f *FunnyBox) Close() error {
	fmt.Println("Bey!")
	return nil
}

// nolint: gosimple // DO NOT delete this comment!
func main() {
	var c io.Closer
	c = &FunnyBox{}
	err := c.Close()
	if err != nil {
		return
	}
}
