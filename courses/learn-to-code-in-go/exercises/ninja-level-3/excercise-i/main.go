package main

import "fmt"

func main() {
	favSport := "Fotball"
	switch favSport {
	case "Football":
		fmt.Println("Did you catch the Superbowl?")
	case "Basball":
		fmt.Println("Whos your favorite pitcher at the moment?")
	default:
		fmt.Println("Go Sports!")
	}
}
