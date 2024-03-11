package main

import (
	"fmt"
	"log"
	"os"
)

// nolint: gomnd // <-- DO NOT delete this line!
func main() {
	data := `"One","Metallica","...And Justice for All",7:27
"Fuel","Metallica","Reload",4:30
"Master of Puppets","Metallica","Master of Puppets",8:36
`

	file, err := os.OpenFile("songs.csv", os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	_, err = fmt.Fprintln(file, data)
	if err != nil {
		log.Fatal(err)
	}

}
