package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

const FloatPrecision = 64

func ScanFloats(data []byte, atEOF bool) (int, []byte, error) {
	advance, token, err := bufio.ScanWords(data, atEOF)
	if err == nil && token != nil {
		_, err = strconv.ParseFloat(string(token), FloatPrecision)
	}
	return advance, token, err
}

func main() {
	scn := bufio.NewScanner(os.Stdin)
	// Set 'ScanFloats' as the split function below
	scn.Split(ScanFloats)

	// Write the for loop to scan and then output the scanned data
	for scn.Scan() {
		fmt.Println(scn.Text())
	}
	if err := scn.Err(); err != nil {
		log.Println(err)
	}
}
