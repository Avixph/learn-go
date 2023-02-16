package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func main() {
	urlOne := "https://www.google.com/"
	// urlTwo := "https://www.google.com/search?q=chihuahua"
	resp, err := http.Get(urlOne)
	// resp, err := http.Get(urlTwo)

	if err != nil {
		fmt.Println("Error: ", err)
		os.Exit(1)
	}
	// fmt.Println(resp)
	// byteSlice := make([]byte, 99999)
	// resp.Body.Read(byteSlice)
	// fmt.Println(string(byteSlice))
	io.Copy(os.Stdout, resp.Body)
	lw := logWriter{}
	io.Copy(lw, resp.Body)
}
