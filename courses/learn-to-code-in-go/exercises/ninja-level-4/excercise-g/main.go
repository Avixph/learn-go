package main

import "fmt"

func main() {
	jb := []string{"James", "Bond", "Shaken, not stirred"}
	mm := []string{"Miss", "Moneypenny", "Helloooooo, James."}
	mi6Agents := [][]string{jb, mm}
	// fmt.Println(jb)
	// fmt.Println(mm)
	// fmt.Println(mi6Agents)
	for _, mi6Agent := range mi6Agents {
		for _, agent := range mi6Agent {
			fmt.Printf("%v\n", agent)
		}
	}
}
